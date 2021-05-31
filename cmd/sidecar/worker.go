package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/vorteil/direktiv/pkg/flow"
)

type inboundWorker struct {
	id     int
	cancel func()
	lock   sync.Mutex
	srv    *LocalServer
}

func (worker *inboundWorker) Cancel() {

	worker.lock.Lock()

	if worker.cancel != nil {
		log.Debugf("Cancelling worker %d.", worker.id)
		worker.cancel()
	}

	worker.lock.Unlock()

}

func (worker *inboundWorker) run() {

	log.Debugf("Starting worker %d.", worker.id)

	for {
		worker.lock.Lock()

		req, more := <-worker.srv.queue
		if req == nil || !more {
			worker.cancel = nil
			worker.lock.Unlock()
			break
		}

		ctx, cancel := context.WithCancel(req.r.Context())
		worker.cancel = cancel
		req.r = req.r.WithContext(ctx)

		worker.lock.Unlock()

		id := req.r.Header.Get(actionIDHeader)
		log.Debugf("Worker %d picked up request '%s'.", worker.id, id)

		worker.handleIsolateRequest(req)

	}

	log.Debugf("Worker %d shut down.", worker.id)

}

func (worker *inboundWorker) fileReader(ctx context.Context, ir *isolateRequest, f *isolateFiles, pw *io.PipeWriter) error {

	defer pw.Close()

	err := worker.srv.getVar(ctx, ir, pw, nil, f.Scope, f.Key)
	if err != nil {
		return err
	}

	return nil

}

type outcome struct {
	data    []byte
	errCode string
	errMsg  string
}

func (worker *inboundWorker) doIsolateRequest(ctx context.Context, ir *isolateRequest) (*outcome, error) {

	log.Debugf("Forwarding request '%s' to service.", ir.actionId)

	url := "http://localhost:8080"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(ir.input))
	if err != nil {
		return nil, err
	}

	req.Header.Set(actionIDHeader, ir.actionId)
	req.Header.Set("Direktiv-TempDir", worker.isolateDir(ir))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out := new(outcome)

	out.errCode = resp.Header.Get("Direktiv-ErrorCode")
	out.errMsg = resp.Header.Get("Direktiv-ErrorMessage")

	if out.errCode != "" {
		return out, nil
	}

	cap := int64(0x400000) // 4 MiB
	if resp.ContentLength > cap {
		return nil, errors.New("service response is too large")
	}
	r := io.LimitReader(resp.Body, cap)

	out.data, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return out, nil

}

func (worker *inboundWorker) prepOneIsolateFiles(ctx context.Context, ir *isolateRequest, f *isolateFiles) error {

	pr, pw := io.Pipe()

	go func() {
		err := worker.fileReader(ctx, ir, f, pw)
		if err != nil {
			_ = pw.CloseWithError(err)
		} else {
			_ = pw.Close()
		}
	}()

	err := worker.fileWriter(ctx, ir, f, pr)
	if err != nil {
		_ = pr.CloseWithError(err)
		return err
	}

	_ = pr.Close()

	return nil

}

func untarFile(tr *tar.Reader, path string) error {

	pdir, _ := filepath.Split(path)
	err := os.MkdirAll(pdir, 0777)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, tr)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil

}

func untar(dst string, r io.Reader) error {

	err := os.MkdirAll(dst, 0777)
	if err != nil {
		return err
	}

	tr := tar.NewReader(r)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(dst, hdr.Name)

		if hdr.Typeflag == tar.TypeReg {
			err = untarFile(tr, path)
			if err != nil {
				return err
			}
		} else if hdr.Typeflag == tar.TypeDir {
			err = os.MkdirAll(path, 0777)
			if err != nil {
				return err
			}
		} else {
			return errors.New("unsupported tar archive contents")
		}

	}

	return nil

}

func writeFile(dst string, r io.Reader) error {

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil

}

func (worker *inboundWorker) writeFile(ftype, dst string, pr io.Reader) error {

	// TODO: const the types

	var err error

	switch ftype {

	case "":
		fallthrough

	case "plain":
		err = writeFile(dst, pr)
		if err != nil {
			return err
		}

	case "base64":
		r := base64.NewDecoder(base64.StdEncoding, pr)
		err = writeFile(dst, r)
		if err != nil {
			return err
		}

	case "tar":
		err = untar(dst, pr)
		if err != nil {
			return err
		}

	case "tar.gz":
		gr, err := gzip.NewReader(pr)
		if err != nil {
			return err
		}

		err = untar(dst, gr)
		if err != nil {
			return err
		}

		err = gr.Close()
		if err != nil {
			return err
		}

	default:
		panic(ftype)
	}

	return nil

}

func (worker *inboundWorker) fileWriter(ctx context.Context, ir *isolateRequest, f *isolateFiles, pr *io.PipeReader) error {

	// TODO: validate f.Type earlier so that the switch cannot get unexpected data here

	dir := worker.isolateDir(ir)
	dst := f.Key
	if f.As != "" {
		dst = f.As
	}
	dst = filepath.Join(dir, dst)
	dir, _ = filepath.Split(dst)

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	err = worker.writeFile(f.Type, dst, pr)
	if err != nil {
		return err
	}

	return nil

}

func (worker *inboundWorker) isolateDir(ir *isolateRequest) string {
	return filepath.Join(sharedDir, ir.actionId)
}

func (worker *inboundWorker) cleanupIsolateRequest(ir *isolateRequest) {
	dir := worker.isolateDir(ir)
	err := os.RemoveAll(dir)
	if err != nil {
		log.Error(err)
	}
}

func (worker *inboundWorker) prepIsolateRequest(ctx context.Context, ir *isolateRequest) error {

	err := worker.prepIsolateFiles(ctx, ir)
	if err != nil {
		return fmt.Errorf("failed to prepare isolate files: %v", err)
	}

	return nil

}

func (worker *inboundWorker) prepIsolateFiles(ctx context.Context, ir *isolateRequest) error {

	dir := worker.isolateDir(ir)

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	for i, f := range ir.files {
		err = worker.prepOneIsolateFiles(ctx, ir, f)
		if err != nil {
			return fmt.Errorf("failed to prepare isolate files %d: %v", i, err)
		}
	}

	return nil

}

func (worker *inboundWorker) handleIsolateRequest(req *inboundRequest) {

	defer func() {
		close(req.end)
	}()

	ir := worker.validateIsolateRequest(req)
	if ir == nil {
		return
	}
	defer ir.logger.Close()

	ctx := req.r.Context()
	ctx, cancel := context.WithDeadline(ctx, ir.deadline)
	defer cancel()

	defer worker.cleanupIsolateRequest(ir)

	err := worker.prepIsolateRequest(ctx, ir)
	if err != nil {
		worker.reportSidecarError(ir, err)
		return
	}

	// NOTE: rctx exists because we don't want to immediately cancel the isolate request if our context is cancelled
	rctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	worker.srv.registerActiveRequest(ir, rctx, cancel)
	defer worker.srv.deregisterActiveRequest(ir.actionId)
	go func() {
		select {
		case <-rctx.Done():
		case <-ctx.Done():
			worker.srv.cancelActiveRequest(rctx, ir.actionId)
		}
	}()

	out, err := worker.doIsolateRequest(rctx, ir)
	if err != nil {
		worker.reportSidecarError(ir, err)
		return
	}

	worker.respondToFlow(rctx, ir, out)

}

func (worker *inboundWorker) respondToFlow(ctx context.Context, ir *isolateRequest, out *outcome) {

	step := int32(ir.step)

	_, err := worker.srv.flow.ReportActionResults(ctx, &flow.ReportActionResultsRequest{
		InstanceId:   &ir.instanceId,
		Step:         &step,
		ActionId:     &ir.actionId,
		Output:       out.data,
		ErrorCode:    &out.errCode,
		ErrorMessage: &out.errMsg,
	})

	if err != nil {
		log.Errorf("Failed to report results for request '%s': %v.", ir.actionId, err)
		return
	}

	if out.errCode != "" {
		log.Infof("Request '%s' failed with catchable error '%s': %s.", ir.actionId, out.errCode, out.errMsg)
	} else if out.errMsg != "" {
		log.Infof("Request '%s' failed with uncatchable service error: %s.", ir.actionId, out.errMsg)
	} else {
		log.Infof("Request '%s' completed successfully.", ir.actionId)
	}

}

func (worker *inboundWorker) reportSidecarError(ir *isolateRequest, err error) {

	ctx := context.Background() // TODO

	worker.respondToFlow(ctx, ir, &outcome{
		errMsg: err.Error(),
	})

}

func (worker *inboundWorker) reportValidationError(req *inboundRequest, code int, err error) {

	id := req.r.Header.Get(actionIDHeader)

	msg := err.Error()

	http.Error(req.w, msg, code)

	log.Warnf("Request '%s' returned %v due to failed validation: %v.", id, code, err)

}

func (worker *inboundWorker) getRequiredStringHeader(req *inboundRequest, x *string, hdr string) bool {

	s := req.r.Header.Get(hdr)
	*x = s
	if s == "" {
		worker.reportValidationError(req, http.StatusBadRequest, fmt.Errorf("missing %s", hdr))
		return false
	}

	return true

}

func (worker *inboundWorker) validateUintHeader(req *inboundRequest, x *int, hdr, s string) bool {

	var err error

	*x, err = strconv.Atoi(s)
	if err != nil {
		worker.reportValidationError(req, http.StatusBadRequest, fmt.Errorf("invalid %s: %v", hdr, err))
		return false
	}
	if *x < 0 {
		worker.reportValidationError(req, http.StatusBadRequest, fmt.Errorf("invalid %s value: %v", hdr, s))
		return false
	}

	return true

}

func (worker *inboundWorker) validateTimeHeader(req *inboundRequest, x *time.Time, hdr, s string) bool {

	var err error

	*x, err = time.Parse(time.RFC3339, s)
	if err != nil {
		worker.reportValidationError(req, http.StatusBadRequest, fmt.Errorf("invalid %s: %v", hdr, err))
		return false
	}

	return true

}

func (worker *inboundWorker) loadBody(req *inboundRequest, data *[]byte) bool {

	cap := int64(0x400000) // 4 MiB
	if req.r.ContentLength == 0 {
		code := http.StatusLengthRequired
		worker.reportValidationError(req, code, errors.New(http.StatusText(code)))
		return false
	}
	if req.r.ContentLength > cap {
		worker.reportValidationError(req, http.StatusRequestEntityTooLarge, fmt.Errorf("size limit: %d bytes", cap))
		return false
	}
	r := io.LimitReader(req.r.Body, cap)

	var err error
	*data, err = ioutil.ReadAll(r)
	if err != nil {
		worker.reportValidationError(req, http.StatusBadRequest, fmt.Errorf("failed to read request body: %v", err))
		return false
	}
	if int64(len(*data)) != req.r.ContentLength {
		worker.reportValidationError(req, http.StatusBadRequest, fmt.Errorf("request body doesn't match Content-Length"))
		return false
	}

	return true

}

func (worker *inboundWorker) validateFilesHeaders(req *inboundRequest, ifiles *[]*isolateFiles) bool {

	hdr := "Direktiv-Files"
	strs := req.r.Header.Values(hdr)
	for i, s := range strs {

		data, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			worker.reportValidationError(req, http.StatusBadRequest, fmt.Errorf("invalid %s [%d]: %v", hdr, i, err))
			return false
		}

		files := new(isolateFiles)
		dec := json.NewDecoder(bytes.NewReader(data))
		dec.DisallowUnknownFields()
		err = dec.Decode(files)
		if err != nil {
			worker.reportValidationError(req, http.StatusBadRequest, fmt.Errorf("invalid %s [%d]: %v", hdr, i, err))
			return false
		}

		// TODO: extra validation

		*ifiles = append(*ifiles, files)

	}

	return true

}

func (worker *inboundWorker) validateIsolateRequest(req *inboundRequest) *isolateRequest {

	var err error

	ir := new(isolateRequest)

	var step string
	var deadline string

	var headers = []string{actionIDHeader, "Direktiv-InstanceID", "Direktiv-Namespace", "Direktiv-Step", "Direktiv-Deadline"}
	var ptrs = []*string{&ir.actionId, &ir.instanceId, &ir.namespace, &step, &deadline}

	for i := 0; i < len(headers); i++ {
		if !worker.getRequiredStringHeader(req, ptrs[i], headers[i]) {
			return nil
		}
	}

	if !worker.validateUintHeader(req, &ir.step, "Direktiv-Step", step) {
		return nil
	}

	if !worker.validateTimeHeader(req, &ir.deadline, "Direktiv-Deadline", deadline) {
		return nil
	}

	if !worker.loadBody(req, &ir.input) {
		return nil
	}

	if !worker.validateFilesHeaders(req, &ir.files) {
		return nil
	}

	ir.logger, err = worker.srv.logging.LoggerFunc(ir.namespace, ir.instanceId)
	if err != nil {
		code := http.StatusInternalServerError
		msg := http.StatusText(code)
		worker.reportValidationError(req, code, errors.New(msg))
		return nil
	}

	return ir

}
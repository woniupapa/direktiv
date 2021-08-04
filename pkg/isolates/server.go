package isolates

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	igrpc "github.com/vorteil/direktiv/pkg/isolates/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	grpcServer    *grpc.Server
	empty         emptypb.Empty
	isolateConfig config
)

const (
	confFile = "/etc/direktiv/config.yaml"
	port     = 5555
)

// Headers for knative services
const (
	ServiceHeaderName      = "direktiv.io/name"
	ServiceHeaderNamespace = "direktiv.io/namespace"
	ServiceHeaderWorkflow  = "direktiv.io/workflow"
	ServiceHeaderSize      = "direktiv.io/size"
	ServiceHeaderScale     = "direktiv.io/scale"
)

type isolateServer struct {
	igrpc.UnimplementedIsolatesServiceServer
}

// StartServer starts isolate grpc server
func StartServer(echan chan error) {

	cr := newConfigReader()
	go cr.readConfig(confFile, &isolateConfig)

	if len(os.Getenv(envFlow)) == 0 {
		log.Errorf("grpc response is not configured (DIREKTIV_FLOW_ENDPOINT)")
		echan <- fmt.Errorf("grpc response is not configured (DIREKTIV_FLOW_ENDPOINT)")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		echan <- err
	}

	var opts []grpc.ServerOption
	tlsPath := "/etc/certs/direktiv/"
	tlsCert := filepath.Join(tlsPath, "tls.crt")
	tlsKey := filepath.Join(tlsPath, "tls.key")

	if _, err = os.Stat(tlsKey); err == nil {
		log.Infof("enabling tls for %s", "isolates")
		creds, err := credentials.NewServerTLSFromFile(tlsCert, tlsKey)
		if err != nil {
			log.Errorf("failed to configure tls opts: %v", err)
			echan <- fmt.Errorf("could not load TLS keys: %s", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	grpcServer = grpc.NewServer(opts...)
	igrpc.RegisterIsolatesServiceServer(grpcServer, &isolateServer{})

	if err := grpcServer.Serve(lis); err != nil {
		echan <- err
	}

}

// StopServer is stopping server gracefully
func StopServer() {
	if grpcServer != nil {
		grpcServer.GracefulStop()
	}
}

func (is *isolateServer) UpdateIsolate(ctx context.Context,
	in *igrpc.UpdateIsolateRequest) (*emptypb.Empty, error) {

	log.Infof("updating isolate %s", in.GetInfo().GetName())

	if in.GetInfo() == nil {
		return &empty, fmt.Errorf("info can not be nil")
	}

	// create ksvc service
	err := updateKnativeIsolate(in.GetServiceName(), in.GetInfo())
	if err != nil {
		log.Errorf("can not update knative service: %v", err)
		return &empty, err
	}

	return &empty, nil
}

func (is *isolateServer) DeleteIsolates(ctx context.Context,
	in *igrpc.ListIsolatesRequest) (*emptypb.Empty, error) {

	log.Debugf("deleting isolates %v", in.GetAnnotations())

	err := deleteIsolates(in.GetAnnotations())

	return &empty, err
}

func (is *isolateServer) GetIsolate(ctx context.Context,
	in *igrpc.GetIsolateRequest) (*igrpc.GetIsolateResponse, error) {

	var resp igrpc.GetIsolateResponse

	getKnativeIsolate(in.GetServiceName())

	return &resp, nil

}

// ListIsolates returns isoaltes based on label filter
func (is *isolateServer) ListIsolates(ctx context.Context,
	in *igrpc.ListIsolatesRequest) (*igrpc.ListIsolatesResponse, error) {

	var resp igrpc.ListIsolatesResponse

	log.Debugf("list isolates %v", in.GetAnnotations())

	items, err := listKnativeIsolates(in.GetAnnotations())
	if err != nil {
		return &resp, nil
	}

	resp.Isolates = items

	return &resp, nil

}

// StoreIsolate saves or updates isolates which means creating knative services
// baes on the provided configuration
func (is *isolateServer) CreateIsolate(ctx context.Context,
	in *igrpc.CreateIsolateRequest) (*emptypb.Empty, error) {

	log.Infof("storing isolate %s", in.GetInfo().GetName())

	if in.GetInfo() == nil {
		return &empty, fmt.Errorf("info can not be nil")
	}

	// create ksvc service
	err := createKnativeIsolate(in.GetInfo())
	if err != nil {
		log.Errorf("can not create knative service: %v", err)
		return &empty, err
	}

	return &empty, nil

}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package flow

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DirektivFlowClient is the client API for DirektivFlow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirektivFlowClient interface {
	ReportActionResults(ctx context.Context, in *ReportActionResultsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetNamespaceVariable(ctx context.Context, in *GetNamespaceVariableRequest, opts ...grpc.CallOption) (DirektivFlow_GetNamespaceVariableClient, error)
	GetWorkflowVariable(ctx context.Context, in *GetWorkflowVariableRequest, opts ...grpc.CallOption) (DirektivFlow_GetWorkflowVariableClient, error)
	GetInstanceVariable(ctx context.Context, in *GetInstanceVariableRequest, opts ...grpc.CallOption) (DirektivFlow_GetInstanceVariableClient, error)
	SetNamespaceVariable(ctx context.Context, opts ...grpc.CallOption) (DirektivFlow_SetNamespaceVariableClient, error)
	SetWorkflowVariable(ctx context.Context, opts ...grpc.CallOption) (DirektivFlow_SetWorkflowVariableClient, error)
	SetInstanceVariable(ctx context.Context, opts ...grpc.CallOption) (DirektivFlow_SetInstanceVariableClient, error)
	ActionLog(ctx context.Context, in *ActionLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type direktivFlowClient struct {
	cc grpc.ClientConnInterface
}

func NewDirektivFlowClient(cc grpc.ClientConnInterface) DirektivFlowClient {
	return &direktivFlowClient{cc}
}

func (c *direktivFlowClient) ReportActionResults(ctx context.Context, in *ReportActionResultsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/flow.DirektivFlow/ReportActionResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivFlowClient) Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/flow.DirektivFlow/Resume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *direktivFlowClient) GetNamespaceVariable(ctx context.Context, in *GetNamespaceVariableRequest, opts ...grpc.CallOption) (DirektivFlow_GetNamespaceVariableClient, error) {
	stream, err := c.cc.NewStream(ctx, &DirektivFlow_ServiceDesc.Streams[0], "/flow.DirektivFlow/GetNamespaceVariable", opts...)
	if err != nil {
		return nil, err
	}
	x := &direktivFlowGetNamespaceVariableClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DirektivFlow_GetNamespaceVariableClient interface {
	Recv() (*GetNamespaceVariableResponse, error)
	grpc.ClientStream
}

type direktivFlowGetNamespaceVariableClient struct {
	grpc.ClientStream
}

func (x *direktivFlowGetNamespaceVariableClient) Recv() (*GetNamespaceVariableResponse, error) {
	m := new(GetNamespaceVariableResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *direktivFlowClient) GetWorkflowVariable(ctx context.Context, in *GetWorkflowVariableRequest, opts ...grpc.CallOption) (DirektivFlow_GetWorkflowVariableClient, error) {
	stream, err := c.cc.NewStream(ctx, &DirektivFlow_ServiceDesc.Streams[1], "/flow.DirektivFlow/GetWorkflowVariable", opts...)
	if err != nil {
		return nil, err
	}
	x := &direktivFlowGetWorkflowVariableClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DirektivFlow_GetWorkflowVariableClient interface {
	Recv() (*GetWorkflowVariableResponse, error)
	grpc.ClientStream
}

type direktivFlowGetWorkflowVariableClient struct {
	grpc.ClientStream
}

func (x *direktivFlowGetWorkflowVariableClient) Recv() (*GetWorkflowVariableResponse, error) {
	m := new(GetWorkflowVariableResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *direktivFlowClient) GetInstanceVariable(ctx context.Context, in *GetInstanceVariableRequest, opts ...grpc.CallOption) (DirektivFlow_GetInstanceVariableClient, error) {
	stream, err := c.cc.NewStream(ctx, &DirektivFlow_ServiceDesc.Streams[2], "/flow.DirektivFlow/GetInstanceVariable", opts...)
	if err != nil {
		return nil, err
	}
	x := &direktivFlowGetInstanceVariableClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DirektivFlow_GetInstanceVariableClient interface {
	Recv() (*GetInstanceVariableResponse, error)
	grpc.ClientStream
}

type direktivFlowGetInstanceVariableClient struct {
	grpc.ClientStream
}

func (x *direktivFlowGetInstanceVariableClient) Recv() (*GetInstanceVariableResponse, error) {
	m := new(GetInstanceVariableResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *direktivFlowClient) SetNamespaceVariable(ctx context.Context, opts ...grpc.CallOption) (DirektivFlow_SetNamespaceVariableClient, error) {
	stream, err := c.cc.NewStream(ctx, &DirektivFlow_ServiceDesc.Streams[3], "/flow.DirektivFlow/SetNamespaceVariable", opts...)
	if err != nil {
		return nil, err
	}
	x := &direktivFlowSetNamespaceVariableClient{stream}
	return x, nil
}

type DirektivFlow_SetNamespaceVariableClient interface {
	Send(*SetNamespaceVariableRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type direktivFlowSetNamespaceVariableClient struct {
	grpc.ClientStream
}

func (x *direktivFlowSetNamespaceVariableClient) Send(m *SetNamespaceVariableRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *direktivFlowSetNamespaceVariableClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *direktivFlowClient) SetWorkflowVariable(ctx context.Context, opts ...grpc.CallOption) (DirektivFlow_SetWorkflowVariableClient, error) {
	stream, err := c.cc.NewStream(ctx, &DirektivFlow_ServiceDesc.Streams[4], "/flow.DirektivFlow/SetWorkflowVariable", opts...)
	if err != nil {
		return nil, err
	}
	x := &direktivFlowSetWorkflowVariableClient{stream}
	return x, nil
}

type DirektivFlow_SetWorkflowVariableClient interface {
	Send(*SetWorkflowVariableRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type direktivFlowSetWorkflowVariableClient struct {
	grpc.ClientStream
}

func (x *direktivFlowSetWorkflowVariableClient) Send(m *SetWorkflowVariableRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *direktivFlowSetWorkflowVariableClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *direktivFlowClient) SetInstanceVariable(ctx context.Context, opts ...grpc.CallOption) (DirektivFlow_SetInstanceVariableClient, error) {
	stream, err := c.cc.NewStream(ctx, &DirektivFlow_ServiceDesc.Streams[5], "/flow.DirektivFlow/SetInstanceVariable", opts...)
	if err != nil {
		return nil, err
	}
	x := &direktivFlowSetInstanceVariableClient{stream}
	return x, nil
}

type DirektivFlow_SetInstanceVariableClient interface {
	Send(*SetInstanceVariableRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type direktivFlowSetInstanceVariableClient struct {
	grpc.ClientStream
}

func (x *direktivFlowSetInstanceVariableClient) Send(m *SetInstanceVariableRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *direktivFlowSetInstanceVariableClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *direktivFlowClient) ActionLog(ctx context.Context, in *ActionLogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/flow.DirektivFlow/ActionLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirektivFlowServer is the server API for DirektivFlow service.
// All implementations must embed UnimplementedDirektivFlowServer
// for forward compatibility
type DirektivFlowServer interface {
	ReportActionResults(context.Context, *ReportActionResultsRequest) (*emptypb.Empty, error)
	Resume(context.Context, *ResumeRequest) (*emptypb.Empty, error)
	GetNamespaceVariable(*GetNamespaceVariableRequest, DirektivFlow_GetNamespaceVariableServer) error
	GetWorkflowVariable(*GetWorkflowVariableRequest, DirektivFlow_GetWorkflowVariableServer) error
	GetInstanceVariable(*GetInstanceVariableRequest, DirektivFlow_GetInstanceVariableServer) error
	SetNamespaceVariable(DirektivFlow_SetNamespaceVariableServer) error
	SetWorkflowVariable(DirektivFlow_SetWorkflowVariableServer) error
	SetInstanceVariable(DirektivFlow_SetInstanceVariableServer) error
	ActionLog(context.Context, *ActionLogRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDirektivFlowServer()
}

// UnimplementedDirektivFlowServer must be embedded to have forward compatible implementations.
type UnimplementedDirektivFlowServer struct {
}

func (UnimplementedDirektivFlowServer) ReportActionResults(context.Context, *ReportActionResultsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportActionResults not implemented")
}
func (UnimplementedDirektivFlowServer) Resume(context.Context, *ResumeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (UnimplementedDirektivFlowServer) GetNamespaceVariable(*GetNamespaceVariableRequest, DirektivFlow_GetNamespaceVariableServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNamespaceVariable not implemented")
}
func (UnimplementedDirektivFlowServer) GetWorkflowVariable(*GetWorkflowVariableRequest, DirektivFlow_GetWorkflowVariableServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWorkflowVariable not implemented")
}
func (UnimplementedDirektivFlowServer) GetInstanceVariable(*GetInstanceVariableRequest, DirektivFlow_GetInstanceVariableServer) error {
	return status.Errorf(codes.Unimplemented, "method GetInstanceVariable not implemented")
}
func (UnimplementedDirektivFlowServer) SetNamespaceVariable(DirektivFlow_SetNamespaceVariableServer) error {
	return status.Errorf(codes.Unimplemented, "method SetNamespaceVariable not implemented")
}
func (UnimplementedDirektivFlowServer) SetWorkflowVariable(DirektivFlow_SetWorkflowVariableServer) error {
	return status.Errorf(codes.Unimplemented, "method SetWorkflowVariable not implemented")
}
func (UnimplementedDirektivFlowServer) SetInstanceVariable(DirektivFlow_SetInstanceVariableServer) error {
	return status.Errorf(codes.Unimplemented, "method SetInstanceVariable not implemented")
}
func (UnimplementedDirektivFlowServer) ActionLog(context.Context, *ActionLogRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionLog not implemented")
}
func (UnimplementedDirektivFlowServer) mustEmbedUnimplementedDirektivFlowServer() {}

// UnsafeDirektivFlowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirektivFlowServer will
// result in compilation errors.
type UnsafeDirektivFlowServer interface {
	mustEmbedUnimplementedDirektivFlowServer()
}

func RegisterDirektivFlowServer(s grpc.ServiceRegistrar, srv DirektivFlowServer) {
	s.RegisterService(&DirektivFlow_ServiceDesc, srv)
}

func _DirektivFlow_ReportActionResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportActionResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivFlowServer).ReportActionResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.DirektivFlow/ReportActionResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivFlowServer).ReportActionResults(ctx, req.(*ReportActionResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivFlow_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivFlowServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.DirektivFlow/Resume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivFlowServer).Resume(ctx, req.(*ResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirektivFlow_GetNamespaceVariable_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetNamespaceVariableRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DirektivFlowServer).GetNamespaceVariable(m, &direktivFlowGetNamespaceVariableServer{stream})
}

type DirektivFlow_GetNamespaceVariableServer interface {
	Send(*GetNamespaceVariableResponse) error
	grpc.ServerStream
}

type direktivFlowGetNamespaceVariableServer struct {
	grpc.ServerStream
}

func (x *direktivFlowGetNamespaceVariableServer) Send(m *GetNamespaceVariableResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DirektivFlow_GetWorkflowVariable_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetWorkflowVariableRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DirektivFlowServer).GetWorkflowVariable(m, &direktivFlowGetWorkflowVariableServer{stream})
}

type DirektivFlow_GetWorkflowVariableServer interface {
	Send(*GetWorkflowVariableResponse) error
	grpc.ServerStream
}

type direktivFlowGetWorkflowVariableServer struct {
	grpc.ServerStream
}

func (x *direktivFlowGetWorkflowVariableServer) Send(m *GetWorkflowVariableResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DirektivFlow_GetInstanceVariable_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetInstanceVariableRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DirektivFlowServer).GetInstanceVariable(m, &direktivFlowGetInstanceVariableServer{stream})
}

type DirektivFlow_GetInstanceVariableServer interface {
	Send(*GetInstanceVariableResponse) error
	grpc.ServerStream
}

type direktivFlowGetInstanceVariableServer struct {
	grpc.ServerStream
}

func (x *direktivFlowGetInstanceVariableServer) Send(m *GetInstanceVariableResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DirektivFlow_SetNamespaceVariable_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DirektivFlowServer).SetNamespaceVariable(&direktivFlowSetNamespaceVariableServer{stream})
}

type DirektivFlow_SetNamespaceVariableServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*SetNamespaceVariableRequest, error)
	grpc.ServerStream
}

type direktivFlowSetNamespaceVariableServer struct {
	grpc.ServerStream
}

func (x *direktivFlowSetNamespaceVariableServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *direktivFlowSetNamespaceVariableServer) Recv() (*SetNamespaceVariableRequest, error) {
	m := new(SetNamespaceVariableRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DirektivFlow_SetWorkflowVariable_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DirektivFlowServer).SetWorkflowVariable(&direktivFlowSetWorkflowVariableServer{stream})
}

type DirektivFlow_SetWorkflowVariableServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*SetWorkflowVariableRequest, error)
	grpc.ServerStream
}

type direktivFlowSetWorkflowVariableServer struct {
	grpc.ServerStream
}

func (x *direktivFlowSetWorkflowVariableServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *direktivFlowSetWorkflowVariableServer) Recv() (*SetWorkflowVariableRequest, error) {
	m := new(SetWorkflowVariableRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DirektivFlow_SetInstanceVariable_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DirektivFlowServer).SetInstanceVariable(&direktivFlowSetInstanceVariableServer{stream})
}

type DirektivFlow_SetInstanceVariableServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*SetInstanceVariableRequest, error)
	grpc.ServerStream
}

type direktivFlowSetInstanceVariableServer struct {
	grpc.ServerStream
}

func (x *direktivFlowSetInstanceVariableServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *direktivFlowSetInstanceVariableServer) Recv() (*SetInstanceVariableRequest, error) {
	m := new(SetInstanceVariableRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DirektivFlow_ActionLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirektivFlowServer).ActionLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flow.DirektivFlow/ActionLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirektivFlowServer).ActionLog(ctx, req.(*ActionLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DirektivFlow_ServiceDesc is the grpc.ServiceDesc for DirektivFlow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirektivFlow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flow.DirektivFlow",
	HandlerType: (*DirektivFlowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportActionResults",
			Handler:    _DirektivFlow_ReportActionResults_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _DirektivFlow_Resume_Handler,
		},
		{
			MethodName: "ActionLog",
			Handler:    _DirektivFlow_ActionLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNamespaceVariable",
			Handler:       _DirektivFlow_GetNamespaceVariable_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetWorkflowVariable",
			Handler:       _DirektivFlow_GetWorkflowVariable_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetInstanceVariable",
			Handler:       _DirektivFlow_GetInstanceVariable_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SetNamespaceVariable",
			Handler:       _DirektivFlow_SetNamespaceVariable_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SetWorkflowVariable",
			Handler:       _DirektivFlow_SetWorkflowVariable_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SetInstanceVariable",
			Handler:       _DirektivFlow_SetInstanceVariable_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/flow/protocol.proto",
}

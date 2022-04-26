// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: orchestrator.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrchestratorServiceClient is the client API for OrchestratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrchestratorServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type orchestratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorServiceClient(cc grpc.ClientConnInterface) OrchestratorServiceClient {
	return &orchestratorServiceClient{cc}
}

func (c *orchestratorServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/go_playground.OrchestratorService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServiceServer is the server API for OrchestratorService service.
// All implementations must embed UnimplementedOrchestratorServiceServer
// for forward compatibility
type OrchestratorServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedOrchestratorServiceServer()
}

// UnimplementedOrchestratorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrchestratorServiceServer struct {
}

func (UnimplementedOrchestratorServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedOrchestratorServiceServer) mustEmbedUnimplementedOrchestratorServiceServer() {}

// UnsafeOrchestratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrchestratorServiceServer will
// result in compilation errors.
type UnsafeOrchestratorServiceServer interface {
	mustEmbedUnimplementedOrchestratorServiceServer()
}

func RegisterOrchestratorServiceServer(s grpc.ServiceRegistrar, srv OrchestratorServiceServer) {
	s.RegisterService(&OrchestratorService_ServiceDesc, srv)
}

func _OrchestratorService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_playground.OrchestratorService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrchestratorService_ServiceDesc is the grpc.ServiceDesc for OrchestratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrchestratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_playground.OrchestratorService",
	HandlerType: (*OrchestratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _OrchestratorService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestrator.proto",
}

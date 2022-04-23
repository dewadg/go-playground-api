// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: executor.proto

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

// ExecutionServiceClient is the client API for ExecutionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutionServiceClient interface {
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type executionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutionServiceClient(cc grpc.ClientConnInterface) ExecutionServiceClient {
	return &executionServiceClient{cc}
}

func (c *executionServiceClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/go_playground.ExecutionService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutionServiceServer is the server API for ExecutionService service.
// All implementations must embed UnimplementedExecutionServiceServer
// for forward compatibility
type ExecutionServiceServer interface {
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	mustEmbedUnimplementedExecutionServiceServer()
}

// UnimplementedExecutionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExecutionServiceServer struct {
}

func (UnimplementedExecutionServiceServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedExecutionServiceServer) mustEmbedUnimplementedExecutionServiceServer() {}

// UnsafeExecutionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutionServiceServer will
// result in compilation errors.
type UnsafeExecutionServiceServer interface {
	mustEmbedUnimplementedExecutionServiceServer()
}

func RegisterExecutionServiceServer(s grpc.ServiceRegistrar, srv ExecutionServiceServer) {
	s.RegisterService(&ExecutionService_ServiceDesc, srv)
}

func _ExecutionService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_playground.ExecutionService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServiceServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExecutionService_ServiceDesc is the grpc.ServiceDesc for ExecutionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExecutionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_playground.ExecutionService",
	HandlerType: (*ExecutionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _ExecutionService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "executor.proto",
}

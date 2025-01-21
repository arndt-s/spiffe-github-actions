// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: init.proto

package main

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InitAPI_Init_FullMethodName = "/InitAPI/Init"
)

// InitAPIClient is the client API for InitAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InitAPIClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
}

type initAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewInitAPIClient(cc grpc.ClientConnInterface) InitAPIClient {
	return &initAPIClient{cc}
}

func (c *initAPIClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, InitAPI_Init_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InitAPIServer is the server API for InitAPI service.
// All implementations must embed UnimplementedInitAPIServer
// for forward compatibility.
type InitAPIServer interface {
	Init(context.Context, *InitRequest) (*InitResponse, error)
	mustEmbedUnimplementedInitAPIServer()
}

// UnimplementedInitAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInitAPIServer struct{}

func (UnimplementedInitAPIServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedInitAPIServer) mustEmbedUnimplementedInitAPIServer() {}
func (UnimplementedInitAPIServer) testEmbeddedByValue()                 {}

// UnsafeInitAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InitAPIServer will
// result in compilation errors.
type UnsafeInitAPIServer interface {
	mustEmbedUnimplementedInitAPIServer()
}

func RegisterInitAPIServer(s grpc.ServiceRegistrar, srv InitAPIServer) {
	// If the following call pancis, it indicates UnimplementedInitAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InitAPI_ServiceDesc, srv)
}

func _InitAPI_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitAPIServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InitAPI_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitAPIServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InitAPI_ServiceDesc is the grpc.ServiceDesc for InitAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InitAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InitAPI",
	HandlerType: (*InitAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _InitAPI_Init_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "init.proto",
}

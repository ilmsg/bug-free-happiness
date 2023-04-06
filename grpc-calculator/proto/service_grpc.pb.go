// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: service.proto

package proto

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

const (
	CalService_Add_FullMethodName      = "/CalService/Add"
	CalService_Subtract_FullMethodName = "/CalService/Subtract"
)

// CalServiceClient is the client API for CalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalServiceClient interface {
	Add(ctx context.Context, in *CalRequest, opts ...grpc.CallOption) (*CalResponse, error)
	Subtract(ctx context.Context, in *CalRequest, opts ...grpc.CallOption) (*CalResponse, error)
}

type calServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalServiceClient(cc grpc.ClientConnInterface) CalServiceClient {
	return &calServiceClient{cc}
}

func (c *calServiceClient) Add(ctx context.Context, in *CalRequest, opts ...grpc.CallOption) (*CalResponse, error) {
	out := new(CalResponse)
	err := c.cc.Invoke(ctx, CalService_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calServiceClient) Subtract(ctx context.Context, in *CalRequest, opts ...grpc.CallOption) (*CalResponse, error) {
	out := new(CalResponse)
	err := c.cc.Invoke(ctx, CalService_Subtract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalServiceServer is the server API for CalService service.
// All implementations must embed UnimplementedCalServiceServer
// for forward compatibility
type CalServiceServer interface {
	Add(context.Context, *CalRequest) (*CalResponse, error)
	Subtract(context.Context, *CalRequest) (*CalResponse, error)
	mustEmbedUnimplementedCalServiceServer()
}

// UnimplementedCalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalServiceServer struct {
}

func (UnimplementedCalServiceServer) Add(context.Context, *CalRequest) (*CalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalServiceServer) Subtract(context.Context, *CalRequest) (*CalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedCalServiceServer) mustEmbedUnimplementedCalServiceServer() {}

// UnsafeCalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalServiceServer will
// result in compilation errors.
type UnsafeCalServiceServer interface {
	mustEmbedUnimplementedCalServiceServer()
}

func RegisterCalServiceServer(s grpc.ServiceRegistrar, srv CalServiceServer) {
	s.RegisterService(&CalService_ServiceDesc, srv)
}

func _CalService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalServiceServer).Add(ctx, req.(*CalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalService_Subtract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalServiceServer).Subtract(ctx, req.(*CalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalService_ServiceDesc is the grpc.ServiceDesc for CalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CalService",
	HandlerType: (*CalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalService_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _CalService_Subtract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
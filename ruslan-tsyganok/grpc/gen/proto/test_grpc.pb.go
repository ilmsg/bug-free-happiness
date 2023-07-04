// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: test.proto

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
	TestApi_Echo_FullMethodName    = "/main.TestApi/Echo"
	TestApi_GetUser_FullMethodName = "/main.TestApi/GetUser"
)

// TestApiClient is the client API for TestApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestApiClient interface {
	Echo(ctx context.Context, in *ResponseRequest, opts ...grpc.CallOption) (*ResponseRequest, error)
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type testApiClient struct {
	cc grpc.ClientConnInterface
}

func NewTestApiClient(cc grpc.ClientConnInterface) TestApiClient {
	return &testApiClient{cc}
}

func (c *testApiClient) Echo(ctx context.Context, in *ResponseRequest, opts ...grpc.CallOption) (*ResponseRequest, error) {
	out := new(ResponseRequest)
	err := c.cc.Invoke(ctx, TestApi_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, TestApi_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestApiServer is the server API for TestApi service.
// All implementations must embed UnimplementedTestApiServer
// for forward compatibility
type TestApiServer interface {
	Echo(context.Context, *ResponseRequest) (*ResponseRequest, error)
	GetUser(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedTestApiServer()
}

// UnimplementedTestApiServer must be embedded to have forward compatible implementations.
type UnimplementedTestApiServer struct {
}

func (UnimplementedTestApiServer) Echo(context.Context, *ResponseRequest) (*ResponseRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedTestApiServer) GetUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedTestApiServer) mustEmbedUnimplementedTestApiServer() {}

// UnsafeTestApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestApiServer will
// result in compilation errors.
type UnsafeTestApiServer interface {
	mustEmbedUnimplementedTestApiServer()
}

func RegisterTestApiServer(s grpc.ServiceRegistrar, srv TestApiServer) {
	s.RegisterService(&TestApi_ServiceDesc, srv)
}

func _TestApi_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).Echo(ctx, req.(*ResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestApi_ServiceDesc is the grpc.ServiceDesc for TestApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.TestApi",
	HandlerType: (*TestApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _TestApi_Echo_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _TestApi_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

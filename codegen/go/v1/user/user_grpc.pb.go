// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	SignUpRequest(ctx context.Context, in *SignUpUserData, opts ...grpc.CallOption) (*SignUpResponseData, error)
	SignInRequest(ctx context.Context, in *SignInRequestData, opts ...grpc.CallOption) (*SignInResponseData, error)
	SayHello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SayHi, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) SignUpRequest(ctx context.Context, in *SignUpUserData, opts ...grpc.CallOption) (*SignUpResponseData, error) {
	out := new(SignUpResponseData)
	err := c.cc.Invoke(ctx, "/user.User/SignUpRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SignInRequest(ctx context.Context, in *SignInRequestData, opts ...grpc.CallOption) (*SignInResponseData, error) {
	out := new(SignInResponseData)
	err := c.cc.Invoke(ctx, "/user.User/SignInRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SayHello(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SayHi, error) {
	out := new(SayHi)
	err := c.cc.Invoke(ctx, "/user.User/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	SignUpRequest(context.Context, *SignUpUserData) (*SignUpResponseData, error)
	SignInRequest(context.Context, *SignInRequestData) (*SignInResponseData, error)
	SayHello(context.Context, *emptypb.Empty) (*SayHi, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) SignUpRequest(context.Context, *SignUpUserData) (*SignUpResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUpRequest not implemented")
}
func (UnimplementedUserServer) SignInRequest(context.Context, *SignInRequestData) (*SignInResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInRequest not implemented")
}
func (UnimplementedUserServer) SayHello(context.Context, *emptypb.Empty) (*SayHi, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_SignUpRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpUserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SignUpRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SignUpRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SignUpRequest(ctx, req.(*SignUpUserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SignInRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequestData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SignInRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SignInRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SignInRequest(ctx, req.(*SignInRequestData))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SayHello(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUpRequest",
			Handler:    _User_SignUpRequest_Handler,
		},
		{
			MethodName: "SignInRequest",
			Handler:    _User_SignInRequest_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _User_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

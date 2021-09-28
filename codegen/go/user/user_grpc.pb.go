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
	GoogleAuth(ctx context.Context, in *GoogleAuthRequest, opts ...grpc.CallOption) (*SelfUser, error)
	AppleAuth(ctx context.Context, in *AppleAuthRequest, opts ...grpc.CallOption) (*SelfUser, error)
	MicrosoftAuth(ctx context.Context, in *MicrosoftAuthRequest, opts ...grpc.CallOption) (*SelfUser, error)
	YahooAuth(ctx context.Context, in *YahooAuthRequest, opts ...grpc.CallOption) (*SelfUser, error)
	FacebookAuth(ctx context.Context, in *FacebookAuthRequest, opts ...grpc.CallOption) (*SelfUser, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*SelfUser, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*ShortUser, error)
	GetByIds(ctx context.Context, in *GetByIdsRequest, opts ...grpc.CallOption) (*ShortUsers, error)
	ModifyRights(ctx context.Context, in *ModifyRightsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SignUp(ctx context.Context, in *SignUpUser, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GoogleAuth(ctx context.Context, in *GoogleAuthRequest, opts ...grpc.CallOption) (*SelfUser, error) {
	out := new(SelfUser)
	err := c.cc.Invoke(ctx, "/user.User/GoogleAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AppleAuth(ctx context.Context, in *AppleAuthRequest, opts ...grpc.CallOption) (*SelfUser, error) {
	out := new(SelfUser)
	err := c.cc.Invoke(ctx, "/user.User/AppleAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) MicrosoftAuth(ctx context.Context, in *MicrosoftAuthRequest, opts ...grpc.CallOption) (*SelfUser, error) {
	out := new(SelfUser)
	err := c.cc.Invoke(ctx, "/user.User/MicrosoftAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) YahooAuth(ctx context.Context, in *YahooAuthRequest, opts ...grpc.CallOption) (*SelfUser, error) {
	out := new(SelfUser)
	err := c.cc.Invoke(ctx, "/user.User/YahooAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) FacebookAuth(ctx context.Context, in *FacebookAuthRequest, opts ...grpc.CallOption) (*SelfUser, error) {
	out := new(SelfUser)
	err := c.cc.Invoke(ctx, "/user.User/FacebookAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*SelfUser, error) {
	out := new(SelfUser)
	err := c.cc.Invoke(ctx, "/user.User/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*ShortUser, error) {
	out := new(ShortUser)
	err := c.cc.Invoke(ctx, "/user.User/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetByIds(ctx context.Context, in *GetByIdsRequest, opts ...grpc.CallOption) (*ShortUsers, error) {
	out := new(ShortUsers)
	err := c.cc.Invoke(ctx, "/user.User/GetByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ModifyRights(ctx context.Context, in *ModifyRightsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.User/ModifyRights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SignUp(ctx context.Context, in *SignUpUser, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.User/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GoogleAuth(context.Context, *GoogleAuthRequest) (*SelfUser, error)
	AppleAuth(context.Context, *AppleAuthRequest) (*SelfUser, error)
	MicrosoftAuth(context.Context, *MicrosoftAuthRequest) (*SelfUser, error)
	YahooAuth(context.Context, *YahooAuthRequest) (*SelfUser, error)
	FacebookAuth(context.Context, *FacebookAuthRequest) (*SelfUser, error)
	Auth(context.Context, *AuthRequest) (*SelfUser, error)
	GetById(context.Context, *GetByIdRequest) (*ShortUser, error)
	GetByIds(context.Context, *GetByIdsRequest) (*ShortUsers, error)
	ModifyRights(context.Context, *ModifyRightsRequest) (*emptypb.Empty, error)
	SignUp(context.Context, *SignUpUser) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GoogleAuth(context.Context, *GoogleAuthRequest) (*SelfUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoogleAuth not implemented")
}
func (UnimplementedUserServer) AppleAuth(context.Context, *AppleAuthRequest) (*SelfUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppleAuth not implemented")
}
func (UnimplementedUserServer) MicrosoftAuth(context.Context, *MicrosoftAuthRequest) (*SelfUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MicrosoftAuth not implemented")
}
func (UnimplementedUserServer) YahooAuth(context.Context, *YahooAuthRequest) (*SelfUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method YahooAuth not implemented")
}
func (UnimplementedUserServer) FacebookAuth(context.Context, *FacebookAuthRequest) (*SelfUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FacebookAuth not implemented")
}
func (UnimplementedUserServer) Auth(context.Context, *AuthRequest) (*SelfUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedUserServer) GetById(context.Context, *GetByIdRequest) (*ShortUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserServer) GetByIds(context.Context, *GetByIdsRequest) (*ShortUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIds not implemented")
}
func (UnimplementedUserServer) ModifyRights(context.Context, *ModifyRightsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyRights not implemented")
}
func (UnimplementedUserServer) SignUp(context.Context, *SignUpUser) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
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

func _User_GoogleAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoogleAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GoogleAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GoogleAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GoogleAuth(ctx, req.(*GoogleAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AppleAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppleAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AppleAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/AppleAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AppleAuth(ctx, req.(*AppleAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_MicrosoftAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicrosoftAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).MicrosoftAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/MicrosoftAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).MicrosoftAuth(ctx, req.(*MicrosoftAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_YahooAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YahooAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).YahooAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/YahooAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).YahooAuth(ctx, req.(*YahooAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_FacebookAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FacebookAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FacebookAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/FacebookAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FacebookAuth(ctx, req.(*FacebookAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetByIds(ctx, req.(*GetByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ModifyRights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyRightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ModifyRights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/ModifyRights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ModifyRights(ctx, req.(*ModifyRightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SignUp(ctx, req.(*SignUpUser))
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
			MethodName: "GoogleAuth",
			Handler:    _User_GoogleAuth_Handler,
		},
		{
			MethodName: "AppleAuth",
			Handler:    _User_AppleAuth_Handler,
		},
		{
			MethodName: "MicrosoftAuth",
			Handler:    _User_MicrosoftAuth_Handler,
		},
		{
			MethodName: "YahooAuth",
			Handler:    _User_YahooAuth_Handler,
		},
		{
			MethodName: "FacebookAuth",
			Handler:    _User_FacebookAuth_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _User_Auth_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _User_GetById_Handler,
		},
		{
			MethodName: "GetByIds",
			Handler:    _User_GetByIds_Handler,
		},
		{
			MethodName: "ModifyRights",
			Handler:    _User_ModifyRights_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _User_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

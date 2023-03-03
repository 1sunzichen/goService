// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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
	GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListRes, error)
	GetUserMobile(ctx context.Context, in *MobileReq, opts ...grpc.CallOption) (*UserInfoRes, error)
	GetUserId(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoRes, error)
	CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoRes, error)
	UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckPassWord(ctx context.Context, in *PassWordInfo, opts ...grpc.CallOption) (*CheckRes, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListRes, error) {
	out := new(UserListRes)
	err := c.cc.Invoke(ctx, "/User/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserMobile(ctx context.Context, in *MobileReq, opts ...grpc.CallOption) (*UserInfoRes, error) {
	out := new(UserInfoRes)
	err := c.cc.Invoke(ctx, "/User/GetUserMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserId(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoRes, error) {
	out := new(UserInfoRes)
	err := c.cc.Invoke(ctx, "/User/GetUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserInfo, opts ...grpc.CallOption) (*UserInfoRes, error) {
	out := new(UserInfoRes)
	err := c.cc.Invoke(ctx, "/User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckPassWord(ctx context.Context, in *PassWordInfo, opts ...grpc.CallOption) (*CheckRes, error) {
	out := new(CheckRes)
	err := c.cc.Invoke(ctx, "/User/CheckPassWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetUserList(context.Context, *PageInfo) (*UserListRes, error)
	GetUserMobile(context.Context, *MobileReq) (*UserInfoRes, error)
	GetUserId(context.Context, *IdReq) (*UserInfoRes, error)
	CreateUser(context.Context, *CreateUserInfo) (*UserInfoRes, error)
	UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error)
	CheckPassWord(context.Context, *PassWordInfo) (*CheckRes, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetUserList(context.Context, *PageInfo) (*UserListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServer) GetUserMobile(context.Context, *MobileReq) (*UserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMobile not implemented")
}
func (UnimplementedUserServer) GetUserId(context.Context, *IdReq) (*UserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserId not implemented")
}
func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserInfo) (*UserInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) CheckPassWord(context.Context, *PassWordInfo) (*CheckRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassWord not implemented")
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

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserMobile(ctx, req.(*MobileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/GetUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserId(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckPassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassWordInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckPassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/CheckPassWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckPassWord(ctx, req.(*PassWordInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "GetUserMobile",
			Handler:    _User_GetUserMobile_Handler,
		},
		{
			MethodName: "GetUserId",
			Handler:    _User_GetUserId_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "CheckPassWord",
			Handler:    _User_CheckPassWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: api/operate/app.proto

package operate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	App_CreateContent_FullMethodName = "/api.operate.App/CreateContent"
	App_UpdateContent_FullMethodName = "/api.operate.App/UpdateContent"
	App_DeleteContent_FullMethodName = "/api.operate.App/DeleteContent"
	App_FindContent_FullMethodName   = "/api.operate.App/FindContent"
)

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	// 创建内容
	CreateContent(ctx context.Context, in *CreateContentReq, opts ...grpc.CallOption) (*CreateContentRsp, error)
	// 更新内容
	UpdateContent(ctx context.Context, in *UpdateContentReq, opts ...grpc.CallOption) (*UpdateContentRsp, error)
	// 删除内容
	DeleteContent(ctx context.Context, in *DeleteContentReq, opts ...grpc.CallOption) (*DeleteContentRsp, error)
	// 内容查找
	FindContent(ctx context.Context, in *FindContentReq, opts ...grpc.CallOption) (*FindContentRsp, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) CreateContent(ctx context.Context, in *CreateContentReq, opts ...grpc.CallOption) (*CreateContentRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateContentRsp)
	err := c.cc.Invoke(ctx, App_CreateContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) UpdateContent(ctx context.Context, in *UpdateContentReq, opts ...grpc.CallOption) (*UpdateContentRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateContentRsp)
	err := c.cc.Invoke(ctx, App_UpdateContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) DeleteContent(ctx context.Context, in *DeleteContentReq, opts ...grpc.CallOption) (*DeleteContentRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteContentRsp)
	err := c.cc.Invoke(ctx, App_DeleteContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) FindContent(ctx context.Context, in *FindContentReq, opts ...grpc.CallOption) (*FindContentRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindContentRsp)
	err := c.cc.Invoke(ctx, App_FindContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	// 创建内容
	CreateContent(context.Context, *CreateContentReq) (*CreateContentRsp, error)
	// 更新内容
	UpdateContent(context.Context, *UpdateContentReq) (*UpdateContentRsp, error)
	// 删除内容
	DeleteContent(context.Context, *DeleteContentReq) (*DeleteContentRsp, error)
	// 内容查找
	FindContent(context.Context, *FindContentReq) (*FindContentRsp, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) CreateContent(context.Context, *CreateContentReq) (*CreateContentRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContent not implemented")
}
func (UnimplementedAppServer) UpdateContent(context.Context, *UpdateContentReq) (*UpdateContentRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContent not implemented")
}
func (UnimplementedAppServer) DeleteContent(context.Context, *DeleteContentReq) (*DeleteContentRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContent not implemented")
}
func (UnimplementedAppServer) FindContent(context.Context, *FindContentReq) (*FindContentRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindContent not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_CreateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).CreateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_CreateContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).CreateContent(ctx, req.(*CreateContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_UpdateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).UpdateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_UpdateContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).UpdateContent(ctx, req.(*UpdateContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_DeleteContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).DeleteContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_DeleteContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).DeleteContent(ctx, req.(*DeleteContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_FindContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).FindContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_FindContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).FindContent(ctx, req.(*FindContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.operate.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContent",
			Handler:    _App_CreateContent_Handler,
		},
		{
			MethodName: "UpdateContent",
			Handler:    _App_UpdateContent_Handler,
		},
		{
			MethodName: "DeleteContent",
			Handler:    _App_DeleteContent_Handler,
		},
		{
			MethodName: "FindContent",
			Handler:    _App_FindContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/operate/app.proto",
}

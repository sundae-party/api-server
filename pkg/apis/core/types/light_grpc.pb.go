// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LightHandlerClient is the client API for LightHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LightHandlerClient interface {
	Get(ctx context.Context, in *Light, opts ...grpc.CallOption) (*Light, error)
	Create(ctx context.Context, in *Light, opts ...grpc.CallOption) (*Light, error)
	Update(ctx context.Context, in *Light, opts ...grpc.CallOption) (*Light, error)
	Delete(ctx context.Context, in *Light, opts ...grpc.CallOption) (*Light, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (LightHandler_GetAllClient, error)
	WatchAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (LightHandler_WatchAllClient, error)
	SetDesiredState(ctx context.Context, in *SetLightStateRequest, opts ...grpc.CallOption) (*Light, error)
	SetState(ctx context.Context, in *SetLightStateRequest, opts ...grpc.CallOption) (*Light, error)
}

type lightHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewLightHandlerClient(cc grpc.ClientConnInterface) LightHandlerClient {
	return &lightHandlerClient{cc}
}

func (c *lightHandlerClient) Get(ctx context.Context, in *Light, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := c.cc.Invoke(ctx, "/types.LightHandler/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightHandlerClient) Create(ctx context.Context, in *Light, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := c.cc.Invoke(ctx, "/types.LightHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightHandlerClient) Update(ctx context.Context, in *Light, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := c.cc.Invoke(ctx, "/types.LightHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightHandlerClient) Delete(ctx context.Context, in *Light, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := c.cc.Invoke(ctx, "/types.LightHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightHandlerClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (LightHandler_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LightHandler_serviceDesc.Streams[0], "/types.LightHandler/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &lightHandlerGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LightHandler_GetAllClient interface {
	Recv() (*Light, error)
	grpc.ClientStream
}

type lightHandlerGetAllClient struct {
	grpc.ClientStream
}

func (x *lightHandlerGetAllClient) Recv() (*Light, error) {
	m := new(Light)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lightHandlerClient) WatchAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (LightHandler_WatchAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LightHandler_serviceDesc.Streams[1], "/types.LightHandler/WatchAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &lightHandlerWatchAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LightHandler_WatchAllClient interface {
	Recv() (*Light, error)
	grpc.ClientStream
}

type lightHandlerWatchAllClient struct {
	grpc.ClientStream
}

func (x *lightHandlerWatchAllClient) Recv() (*Light, error) {
	m := new(Light)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lightHandlerClient) SetDesiredState(ctx context.Context, in *SetLightStateRequest, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := c.cc.Invoke(ctx, "/types.LightHandler/SetDesiredState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightHandlerClient) SetState(ctx context.Context, in *SetLightStateRequest, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := c.cc.Invoke(ctx, "/types.LightHandler/setState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LightHandlerServer is the server API for LightHandler service.
// All implementations must embed UnimplementedLightHandlerServer
// for forward compatibility
type LightHandlerServer interface {
	Get(context.Context, *Light) (*Light, error)
	Create(context.Context, *Light) (*Light, error)
	Update(context.Context, *Light) (*Light, error)
	Delete(context.Context, *Light) (*Light, error)
	GetAll(*GetAllRequest, LightHandler_GetAllServer) error
	WatchAll(*GetAllRequest, LightHandler_WatchAllServer) error
	SetDesiredState(context.Context, *SetLightStateRequest) (*Light, error)
	SetState(context.Context, *SetLightStateRequest) (*Light, error)
	mustEmbedUnimplementedLightHandlerServer()
}

// UnimplementedLightHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedLightHandlerServer struct {
}

func (UnimplementedLightHandlerServer) Get(context.Context, *Light) (*Light, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLightHandlerServer) Create(context.Context, *Light) (*Light, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLightHandlerServer) Update(context.Context, *Light) (*Light, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLightHandlerServer) Delete(context.Context, *Light) (*Light, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLightHandlerServer) GetAll(*GetAllRequest, LightHandler_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedLightHandlerServer) WatchAll(*GetAllRequest, LightHandler_WatchAllServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchAll not implemented")
}
func (UnimplementedLightHandlerServer) SetDesiredState(context.Context, *SetLightStateRequest) (*Light, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDesiredState not implemented")
}
func (UnimplementedLightHandlerServer) SetState(context.Context, *SetLightStateRequest) (*Light, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetState not implemented")
}
func (UnimplementedLightHandlerServer) mustEmbedUnimplementedLightHandlerServer() {}

// UnsafeLightHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LightHandlerServer will
// result in compilation errors.
type UnsafeLightHandlerServer interface {
	mustEmbedUnimplementedLightHandlerServer()
}

func RegisterLightHandlerServer(s grpc.ServiceRegistrar, srv LightHandlerServer) {
	s.RegisterService(&_LightHandler_serviceDesc, srv)
}

func _LightHandler_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Light)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightHandlerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.LightHandler/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightHandlerServer).Get(ctx, req.(*Light))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Light)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.LightHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightHandlerServer).Create(ctx, req.(*Light))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Light)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.LightHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightHandlerServer).Update(ctx, req.(*Light))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Light)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.LightHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightHandlerServer).Delete(ctx, req.(*Light))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightHandler_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LightHandlerServer).GetAll(m, &lightHandlerGetAllServer{stream})
}

type LightHandler_GetAllServer interface {
	Send(*Light) error
	grpc.ServerStream
}

type lightHandlerGetAllServer struct {
	grpc.ServerStream
}

func (x *lightHandlerGetAllServer) Send(m *Light) error {
	return x.ServerStream.SendMsg(m)
}

func _LightHandler_WatchAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LightHandlerServer).WatchAll(m, &lightHandlerWatchAllServer{stream})
}

type LightHandler_WatchAllServer interface {
	Send(*Light) error
	grpc.ServerStream
}

type lightHandlerWatchAllServer struct {
	grpc.ServerStream
}

func (x *lightHandlerWatchAllServer) Send(m *Light) error {
	return x.ServerStream.SendMsg(m)
}

func _LightHandler_SetDesiredState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLightStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightHandlerServer).SetDesiredState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.LightHandler/SetDesiredState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightHandlerServer).SetDesiredState(ctx, req.(*SetLightStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightHandler_SetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLightStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightHandlerServer).SetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.LightHandler/setState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightHandlerServer).SetState(ctx, req.(*SetLightStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LightHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.LightHandler",
	HandlerType: (*LightHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _LightHandler_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LightHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LightHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LightHandler_Delete_Handler,
		},
		{
			MethodName: "SetDesiredState",
			Handler:    _LightHandler_SetDesiredState_Handler,
		},
		{
			MethodName: "setState",
			Handler:    _LightHandler_SetState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _LightHandler_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchAll",
			Handler:       _LightHandler_WatchAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "light.proto",
}

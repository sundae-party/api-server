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

// SensorHandlerClient is the client API for SensorHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SensorHandlerClient interface {
	Get(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error)
	Create(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error)
	Update(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error)
	Delete(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (SensorHandler_GetAllClient, error)
	WatchAll(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (SensorHandler_WatchAllClient, error)
	SetValue(ctx context.Context, in *SetSensorValueRequest, opts ...grpc.CallOption) (*Sensor, error)
}

type sensorHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewSensorHandlerClient(cc grpc.ClientConnInterface) SensorHandlerClient {
	return &sensorHandlerClient{cc}
}

func (c *sensorHandlerClient) Get(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error) {
	out := new(Sensor)
	err := c.cc.Invoke(ctx, "/types.SensorHandler/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorHandlerClient) Create(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error) {
	out := new(Sensor)
	err := c.cc.Invoke(ctx, "/types.SensorHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorHandlerClient) Update(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error) {
	out := new(Sensor)
	err := c.cc.Invoke(ctx, "/types.SensorHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorHandlerClient) Delete(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error) {
	out := new(Sensor)
	err := c.cc.Invoke(ctx, "/types.SensorHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorHandlerClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (SensorHandler_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SensorHandler_serviceDesc.Streams[0], "/types.SensorHandler/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &sensorHandlerGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SensorHandler_GetAllClient interface {
	Recv() (*Sensor, error)
	grpc.ClientStream
}

type sensorHandlerGetAllClient struct {
	grpc.ClientStream
}

func (x *sensorHandlerGetAllClient) Recv() (*Sensor, error) {
	m := new(Sensor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sensorHandlerClient) WatchAll(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (SensorHandler_WatchAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SensorHandler_serviceDesc.Streams[1], "/types.SensorHandler/WatchAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &sensorHandlerWatchAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SensorHandler_WatchAllClient interface {
	Recv() (*Sensor, error)
	grpc.ClientStream
}

type sensorHandlerWatchAllClient struct {
	grpc.ClientStream
}

func (x *sensorHandlerWatchAllClient) Recv() (*Sensor, error) {
	m := new(Sensor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sensorHandlerClient) SetValue(ctx context.Context, in *SetSensorValueRequest, opts ...grpc.CallOption) (*Sensor, error) {
	out := new(Sensor)
	err := c.cc.Invoke(ctx, "/types.SensorHandler/setValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SensorHandlerServer is the server API for SensorHandler service.
// All implementations must embed UnimplementedSensorHandlerServer
// for forward compatibility
type SensorHandlerServer interface {
	Get(context.Context, *Sensor) (*Sensor, error)
	Create(context.Context, *Sensor) (*Sensor, error)
	Update(context.Context, *Sensor) (*Sensor, error)
	Delete(context.Context, *Sensor) (*Sensor, error)
	GetAll(*GetAllRequest, SensorHandler_GetAllServer) error
	WatchAll(*Sensor, SensorHandler_WatchAllServer) error
	SetValue(context.Context, *SetSensorValueRequest) (*Sensor, error)
	mustEmbedUnimplementedSensorHandlerServer()
}

// UnimplementedSensorHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedSensorHandlerServer struct {
}

func (UnimplementedSensorHandlerServer) Get(context.Context, *Sensor) (*Sensor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSensorHandlerServer) Create(context.Context, *Sensor) (*Sensor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSensorHandlerServer) Update(context.Context, *Sensor) (*Sensor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSensorHandlerServer) Delete(context.Context, *Sensor) (*Sensor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSensorHandlerServer) GetAll(*GetAllRequest, SensorHandler_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedSensorHandlerServer) WatchAll(*Sensor, SensorHandler_WatchAllServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchAll not implemented")
}
func (UnimplementedSensorHandlerServer) SetValue(context.Context, *SetSensorValueRequest) (*Sensor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}
func (UnimplementedSensorHandlerServer) mustEmbedUnimplementedSensorHandlerServer() {}

// UnsafeSensorHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SensorHandlerServer will
// result in compilation errors.
type UnsafeSensorHandlerServer interface {
	mustEmbedUnimplementedSensorHandlerServer()
}

func RegisterSensorHandlerServer(s grpc.ServiceRegistrar, srv SensorHandlerServer) {
	s.RegisterService(&_SensorHandler_serviceDesc, srv)
}

func _SensorHandler_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sensor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorHandlerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.SensorHandler/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorHandlerServer).Get(ctx, req.(*Sensor))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sensor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.SensorHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorHandlerServer).Create(ctx, req.(*Sensor))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sensor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.SensorHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorHandlerServer).Update(ctx, req.(*Sensor))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sensor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.SensorHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorHandlerServer).Delete(ctx, req.(*Sensor))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorHandler_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SensorHandlerServer).GetAll(m, &sensorHandlerGetAllServer{stream})
}

type SensorHandler_GetAllServer interface {
	Send(*Sensor) error
	grpc.ServerStream
}

type sensorHandlerGetAllServer struct {
	grpc.ServerStream
}

func (x *sensorHandlerGetAllServer) Send(m *Sensor) error {
	return x.ServerStream.SendMsg(m)
}

func _SensorHandler_WatchAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Sensor)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SensorHandlerServer).WatchAll(m, &sensorHandlerWatchAllServer{stream})
}

type SensorHandler_WatchAllServer interface {
	Send(*Sensor) error
	grpc.ServerStream
}

type sensorHandlerWatchAllServer struct {
	grpc.ServerStream
}

func (x *sensorHandlerWatchAllServer) Send(m *Sensor) error {
	return x.ServerStream.SendMsg(m)
}

func _SensorHandler_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSensorValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorHandlerServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.SensorHandler/setValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorHandlerServer).SetValue(ctx, req.(*SetSensorValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SensorHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.SensorHandler",
	HandlerType: (*SensorHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SensorHandler_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SensorHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SensorHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SensorHandler_Delete_Handler,
		},
		{
			MethodName: "setValue",
			Handler:    _SensorHandler_SetValue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _SensorHandler_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchAll",
			Handler:       _SensorHandler_WatchAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sensor.proto",
}
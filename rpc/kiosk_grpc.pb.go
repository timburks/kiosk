// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: examples/kiosk/v1/kiosk.proto

package kiosk

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

// DisplayClient is the client API for Display service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DisplayClient interface {
	// Create a kiosk and enroll the kiosk for sign display.
	CreateKiosk(ctx context.Context, in *Kiosk, opts ...grpc.CallOption) (*Kiosk, error)
	// List active kiosks.
	ListKiosks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListKiosksResponse, error)
	// Get a kiosk.
	GetKiosk(ctx context.Context, in *GetKioskRequest, opts ...grpc.CallOption) (*Kiosk, error)
	// Delete a kiosk.
	DeleteKiosk(ctx context.Context, in *DeleteKioskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create a sign and enroll the sign for sign display.
	CreateSign(ctx context.Context, in *Sign, opts ...grpc.CallOption) (*Sign, error)
	// List active signs.
	ListSigns(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListSignsResponse, error)
	// Get a sign.
	GetSign(ctx context.Context, in *GetSignRequest, opts ...grpc.CallOption) (*Sign, error)
	// Delete a sign.
	DeleteSign(ctx context.Context, in *DeleteSignRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Set a sign for display on one or more kiosks.
	SetSignIdForKioskIds(ctx context.Context, in *SetSignIdForKioskIdsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the sign that should be displayed on a kiosk.
	GetSignIdForKioskId(ctx context.Context, in *GetSignIdForKioskIdRequest, opts ...grpc.CallOption) (*GetSignIdResponse, error)
	// Get signs that should be displayed on a kiosk. Streams.
	GetSignIdsForKioskId(ctx context.Context, in *GetSignIdForKioskIdRequest, opts ...grpc.CallOption) (Display_GetSignIdsForKioskIdClient, error)
}

type displayClient struct {
	cc grpc.ClientConnInterface
}

func NewDisplayClient(cc grpc.ClientConnInterface) DisplayClient {
	return &displayClient{cc}
}

func (c *displayClient) CreateKiosk(ctx context.Context, in *Kiosk, opts ...grpc.CallOption) (*Kiosk, error) {
	out := new(Kiosk)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/CreateKiosk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) ListKiosks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListKiosksResponse, error) {
	out := new(ListKiosksResponse)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/ListKiosks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) GetKiosk(ctx context.Context, in *GetKioskRequest, opts ...grpc.CallOption) (*Kiosk, error) {
	out := new(Kiosk)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/GetKiosk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) DeleteKiosk(ctx context.Context, in *DeleteKioskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/DeleteKiosk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) CreateSign(ctx context.Context, in *Sign, opts ...grpc.CallOption) (*Sign, error) {
	out := new(Sign)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/CreateSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) ListSigns(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListSignsResponse, error) {
	out := new(ListSignsResponse)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/ListSigns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) GetSign(ctx context.Context, in *GetSignRequest, opts ...grpc.CallOption) (*Sign, error) {
	out := new(Sign)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/GetSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) DeleteSign(ctx context.Context, in *DeleteSignRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/DeleteSign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) SetSignIdForKioskIds(ctx context.Context, in *SetSignIdForKioskIdsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/SetSignIdForKioskIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) GetSignIdForKioskId(ctx context.Context, in *GetSignIdForKioskIdRequest, opts ...grpc.CallOption) (*GetSignIdResponse, error) {
	out := new(GetSignIdResponse)
	err := c.cc.Invoke(ctx, "/examples.kiosk.v1.Display/GetSignIdForKioskId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayClient) GetSignIdsForKioskId(ctx context.Context, in *GetSignIdForKioskIdRequest, opts ...grpc.CallOption) (Display_GetSignIdsForKioskIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &Display_ServiceDesc.Streams[0], "/examples.kiosk.v1.Display/GetSignIdsForKioskId", opts...)
	if err != nil {
		return nil, err
	}
	x := &displayGetSignIdsForKioskIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Display_GetSignIdsForKioskIdClient interface {
	Recv() (*GetSignIdResponse, error)
	grpc.ClientStream
}

type displayGetSignIdsForKioskIdClient struct {
	grpc.ClientStream
}

func (x *displayGetSignIdsForKioskIdClient) Recv() (*GetSignIdResponse, error) {
	m := new(GetSignIdResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DisplayServer is the server API for Display service.
// All implementations must embed UnimplementedDisplayServer
// for forward compatibility
type DisplayServer interface {
	// Create a kiosk and enroll the kiosk for sign display.
	CreateKiosk(context.Context, *Kiosk) (*Kiosk, error)
	// List active kiosks.
	ListKiosks(context.Context, *emptypb.Empty) (*ListKiosksResponse, error)
	// Get a kiosk.
	GetKiosk(context.Context, *GetKioskRequest) (*Kiosk, error)
	// Delete a kiosk.
	DeleteKiosk(context.Context, *DeleteKioskRequest) (*emptypb.Empty, error)
	// Create a sign and enroll the sign for sign display.
	CreateSign(context.Context, *Sign) (*Sign, error)
	// List active signs.
	ListSigns(context.Context, *emptypb.Empty) (*ListSignsResponse, error)
	// Get a sign.
	GetSign(context.Context, *GetSignRequest) (*Sign, error)
	// Delete a sign.
	DeleteSign(context.Context, *DeleteSignRequest) (*emptypb.Empty, error)
	// Set a sign for display on one or more kiosks.
	SetSignIdForKioskIds(context.Context, *SetSignIdForKioskIdsRequest) (*emptypb.Empty, error)
	// Get the sign that should be displayed on a kiosk.
	GetSignIdForKioskId(context.Context, *GetSignIdForKioskIdRequest) (*GetSignIdResponse, error)
	// Get signs that should be displayed on a kiosk. Streams.
	GetSignIdsForKioskId(*GetSignIdForKioskIdRequest, Display_GetSignIdsForKioskIdServer) error
	mustEmbedUnimplementedDisplayServer()
}

// UnimplementedDisplayServer must be embedded to have forward compatible implementations.
type UnimplementedDisplayServer struct {
}

func (UnimplementedDisplayServer) CreateKiosk(context.Context, *Kiosk) (*Kiosk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKiosk not implemented")
}
func (UnimplementedDisplayServer) ListKiosks(context.Context, *emptypb.Empty) (*ListKiosksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKiosks not implemented")
}
func (UnimplementedDisplayServer) GetKiosk(context.Context, *GetKioskRequest) (*Kiosk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKiosk not implemented")
}
func (UnimplementedDisplayServer) DeleteKiosk(context.Context, *DeleteKioskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKiosk not implemented")
}
func (UnimplementedDisplayServer) CreateSign(context.Context, *Sign) (*Sign, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSign not implemented")
}
func (UnimplementedDisplayServer) ListSigns(context.Context, *emptypb.Empty) (*ListSignsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSigns not implemented")
}
func (UnimplementedDisplayServer) GetSign(context.Context, *GetSignRequest) (*Sign, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSign not implemented")
}
func (UnimplementedDisplayServer) DeleteSign(context.Context, *DeleteSignRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSign not implemented")
}
func (UnimplementedDisplayServer) SetSignIdForKioskIds(context.Context, *SetSignIdForKioskIdsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSignIdForKioskIds not implemented")
}
func (UnimplementedDisplayServer) GetSignIdForKioskId(context.Context, *GetSignIdForKioskIdRequest) (*GetSignIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignIdForKioskId not implemented")
}
func (UnimplementedDisplayServer) GetSignIdsForKioskId(*GetSignIdForKioskIdRequest, Display_GetSignIdsForKioskIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSignIdsForKioskId not implemented")
}
func (UnimplementedDisplayServer) mustEmbedUnimplementedDisplayServer() {}

// UnsafeDisplayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DisplayServer will
// result in compilation errors.
type UnsafeDisplayServer interface {
	mustEmbedUnimplementedDisplayServer()
}

func RegisterDisplayServer(s grpc.ServiceRegistrar, srv DisplayServer) {
	s.RegisterService(&Display_ServiceDesc, srv)
}

func _Display_CreateKiosk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Kiosk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).CreateKiosk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/CreateKiosk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).CreateKiosk(ctx, req.(*Kiosk))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_ListKiosks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).ListKiosks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/ListKiosks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).ListKiosks(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_GetKiosk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKioskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).GetKiosk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/GetKiosk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).GetKiosk(ctx, req.(*GetKioskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_DeleteKiosk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKioskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).DeleteKiosk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/DeleteKiosk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).DeleteKiosk(ctx, req.(*DeleteKioskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_CreateSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sign)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).CreateSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/CreateSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).CreateSign(ctx, req.(*Sign))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_ListSigns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).ListSigns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/ListSigns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).ListSigns(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_GetSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).GetSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/GetSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).GetSign(ctx, req.(*GetSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_DeleteSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).DeleteSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/DeleteSign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).DeleteSign(ctx, req.(*DeleteSignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_SetSignIdForKioskIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSignIdForKioskIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).SetSignIdForKioskIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/SetSignIdForKioskIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).SetSignIdForKioskIds(ctx, req.(*SetSignIdForKioskIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_GetSignIdForKioskId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignIdForKioskIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServer).GetSignIdForKioskId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.kiosk.v1.Display/GetSignIdForKioskId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServer).GetSignIdForKioskId(ctx, req.(*GetSignIdForKioskIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Display_GetSignIdsForKioskId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSignIdForKioskIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DisplayServer).GetSignIdsForKioskId(m, &displayGetSignIdsForKioskIdServer{stream})
}

type Display_GetSignIdsForKioskIdServer interface {
	Send(*GetSignIdResponse) error
	grpc.ServerStream
}

type displayGetSignIdsForKioskIdServer struct {
	grpc.ServerStream
}

func (x *displayGetSignIdsForKioskIdServer) Send(m *GetSignIdResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Display_ServiceDesc is the grpc.ServiceDesc for Display service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Display_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "examples.kiosk.v1.Display",
	HandlerType: (*DisplayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKiosk",
			Handler:    _Display_CreateKiosk_Handler,
		},
		{
			MethodName: "ListKiosks",
			Handler:    _Display_ListKiosks_Handler,
		},
		{
			MethodName: "GetKiosk",
			Handler:    _Display_GetKiosk_Handler,
		},
		{
			MethodName: "DeleteKiosk",
			Handler:    _Display_DeleteKiosk_Handler,
		},
		{
			MethodName: "CreateSign",
			Handler:    _Display_CreateSign_Handler,
		},
		{
			MethodName: "ListSigns",
			Handler:    _Display_ListSigns_Handler,
		},
		{
			MethodName: "GetSign",
			Handler:    _Display_GetSign_Handler,
		},
		{
			MethodName: "DeleteSign",
			Handler:    _Display_DeleteSign_Handler,
		},
		{
			MethodName: "SetSignIdForKioskIds",
			Handler:    _Display_SetSignIdForKioskIds_Handler,
		},
		{
			MethodName: "GetSignIdForKioskId",
			Handler:    _Display_GetSignIdForKioskId_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSignIdsForKioskId",
			Handler:       _Display_GetSignIdsForKioskId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "examples/kiosk/v1/kiosk.proto",
}

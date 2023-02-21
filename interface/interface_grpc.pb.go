// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: interface/interface.proto

package intf

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

// InterfaceClient is the client API for Interface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InterfaceClient interface {
	// SetLoopbackMode is used to set the mode of loopback on a interface.
	SetLoopbackMode(ctx context.Context, in *SetLoopbackModeRequest, opts ...grpc.CallOption) (*SetLoopbackModeResponse, error)
	// GetLoopbackMode is used to get the mode of loopback on a interface.
	GetLoopbackMode(ctx context.Context, in *GetLoopbackModeRequest, opts ...grpc.CallOption) (*GetLoopbackModeResponse, error)
	// ClearInterfaceCounters will reset the counters for the provided interface.
	ClearInterfaceCounters(ctx context.Context, in *ClearInterfaceCountersRequest, opts ...grpc.CallOption) (*ClearInterfaceCountersResponse, error)
}

type interfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewInterfaceClient(cc grpc.ClientConnInterface) InterfaceClient {
	return &interfaceClient{cc}
}

func (c *interfaceClient) SetLoopbackMode(ctx context.Context, in *SetLoopbackModeRequest, opts ...grpc.CallOption) (*SetLoopbackModeResponse, error) {
	out := new(SetLoopbackModeResponse)
	err := c.cc.Invoke(ctx, "/gnoi.interface.Interface/SetLoopbackMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceClient) GetLoopbackMode(ctx context.Context, in *GetLoopbackModeRequest, opts ...grpc.CallOption) (*GetLoopbackModeResponse, error) {
	out := new(GetLoopbackModeResponse)
	err := c.cc.Invoke(ctx, "/gnoi.interface.Interface/GetLoopbackMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interfaceClient) ClearInterfaceCounters(ctx context.Context, in *ClearInterfaceCountersRequest, opts ...grpc.CallOption) (*ClearInterfaceCountersResponse, error) {
	out := new(ClearInterfaceCountersResponse)
	err := c.cc.Invoke(ctx, "/gnoi.interface.Interface/ClearInterfaceCounters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InterfaceServer is the server API for Interface service.
// All implementations must embed UnimplementedInterfaceServer
// for forward compatibility
type InterfaceServer interface {
	// SetLoopbackMode is used to set the mode of loopback on a interface.
	SetLoopbackMode(context.Context, *SetLoopbackModeRequest) (*SetLoopbackModeResponse, error)
	// GetLoopbackMode is used to get the mode of loopback on a interface.
	GetLoopbackMode(context.Context, *GetLoopbackModeRequest) (*GetLoopbackModeResponse, error)
	// ClearInterfaceCounters will reset the counters for the provided interface.
	ClearInterfaceCounters(context.Context, *ClearInterfaceCountersRequest) (*ClearInterfaceCountersResponse, error)
	mustEmbedUnimplementedInterfaceServer()
}

// UnimplementedInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedInterfaceServer struct {
}

func (UnimplementedInterfaceServer) SetLoopbackMode(context.Context, *SetLoopbackModeRequest) (*SetLoopbackModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLoopbackMode not implemented")
}
func (UnimplementedInterfaceServer) GetLoopbackMode(context.Context, *GetLoopbackModeRequest) (*GetLoopbackModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoopbackMode not implemented")
}
func (UnimplementedInterfaceServer) ClearInterfaceCounters(context.Context, *ClearInterfaceCountersRequest) (*ClearInterfaceCountersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearInterfaceCounters not implemented")
}
func (UnimplementedInterfaceServer) mustEmbedUnimplementedInterfaceServer() {}

// UnsafeInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InterfaceServer will
// result in compilation errors.
type UnsafeInterfaceServer interface {
	mustEmbedUnimplementedInterfaceServer()
}

func RegisterInterfaceServer(s grpc.ServiceRegistrar, srv InterfaceServer) {
	s.RegisterService(&Interface_ServiceDesc, srv)
}

func _Interface_SetLoopbackMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLoopbackModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).SetLoopbackMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.interface.Interface/SetLoopbackMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).SetLoopbackMode(ctx, req.(*SetLoopbackModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interface_GetLoopbackMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoopbackModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).GetLoopbackMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.interface.Interface/GetLoopbackMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).GetLoopbackMode(ctx, req.(*GetLoopbackModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interface_ClearInterfaceCounters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearInterfaceCountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterfaceServer).ClearInterfaceCounters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.interface.Interface/ClearInterfaceCounters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterfaceServer).ClearInterfaceCounters(ctx, req.(*ClearInterfaceCountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Interface_ServiceDesc is the grpc.ServiceDesc for Interface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.interface.Interface",
	HandlerType: (*InterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLoopbackMode",
			Handler:    _Interface_SetLoopbackMode_Handler,
		},
		{
			MethodName: "GetLoopbackMode",
			Handler:    _Interface_GetLoopbackMode_Handler,
		},
		{
			MethodName: "ClearInterfaceCounters",
			Handler:    _Interface_ClearInterfaceCounters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interface/interface.proto",
}

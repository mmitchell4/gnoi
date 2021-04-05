// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wavelength_router

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

// WavelengthRouterClient is the client API for WavelengthRouter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WavelengthRouterClient interface {
	// Deprecated: Do not use.
	// AdjustPSD performs a power spectral density (PSD) adjustment on an optical
	// DWDM line system port.
	//
	// This RPC's initial message triggers adjustments. The target should validate
	// the request, and immediately begin the adjustment as long as another
	// adjustment is not already running. As this RPC may take multiple minutes,
	// the state should be regularly streamed to the client. The operation should
	// be completed on the device regardless of the state of the RPC channel to
	// the initiating client. A client that becomes disconnected during an
	// operation can reconnect, requesting the same operation, and receive status
	// updates this way. There is no means by which an adjustment that has been
	// triggered by a client can be cancelled unless the CancelAdjustPSD RPC is
	// used.
	// DEPRECATED, use AdjustSpectrum.
	AdjustPSD(ctx context.Context, in *AdjustPSDRequest, opts ...grpc.CallOption) (WavelengthRouter_AdjustPSDClient, error)
	// Deprecated: Do not use.
	// CancelAdjustPSD cancels an in-progress AdjustPSD request. If an
	// adjustment is not being performed for the provided request, then the RPC
	// should return successfully.
	// DEPRECATED, use CancelAdjustSpectrum.
	CancelAdjustPSD(ctx context.Context, in *AdjustPSDRequest, opts ...grpc.CallOption) (*CancelAdjustPSDResponse, error)
	// AdjustSpectrum performs a spectrum power adjustment on an optical
	// DWDM line system port.
	//
	// This RPC's initial message triggers adjustments. The target should validate
	// the request, and immediately begin the adjustment as long as another
	// adjustment is not already running. As this RPC may take multiple minutes,
	// the state should be regularly streamed to the client. The operation should
	// be completed on the device regardless of the state of the RPC channel to
	// the initiating client. A client that becomes disconnected during an
	// operation can reconnect, requesting the same operation, and receive status
	// updates this way. There is no means by which an adjustment that has been
	// triggered by a client can be cancelled unless the CancelAdjustSpectrum
	// RPC is used.
	AdjustSpectrum(ctx context.Context, in *AdjustSpectrumRequest, opts ...grpc.CallOption) (WavelengthRouter_AdjustSpectrumClient, error)
	// CancelAdjustSpectrum cancels an in-progress AdjustSpectrum request. If an
	// adjustment is not being performed for the provided request, then the RPC
	// should return successfully.
	CancelAdjustSpectrum(ctx context.Context, in *AdjustSpectrumRequest, opts ...grpc.CallOption) (*CancelAdjustSpectrumResponse, error)
}

type wavelengthRouterClient struct {
	cc grpc.ClientConnInterface
}

func NewWavelengthRouterClient(cc grpc.ClientConnInterface) WavelengthRouterClient {
	return &wavelengthRouterClient{cc}
}

// Deprecated: Do not use.
func (c *wavelengthRouterClient) AdjustPSD(ctx context.Context, in *AdjustPSDRequest, opts ...grpc.CallOption) (WavelengthRouter_AdjustPSDClient, error) {
	stream, err := c.cc.NewStream(ctx, &WavelengthRouter_ServiceDesc.Streams[0], "/gnoi.optical.WavelengthRouter/AdjustPSD", opts...)
	if err != nil {
		return nil, err
	}
	x := &wavelengthRouterAdjustPSDClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WavelengthRouter_AdjustPSDClient interface {
	Recv() (*AdjustPSDResponse, error)
	grpc.ClientStream
}

type wavelengthRouterAdjustPSDClient struct {
	grpc.ClientStream
}

func (x *wavelengthRouterAdjustPSDClient) Recv() (*AdjustPSDResponse, error) {
	m := new(AdjustPSDResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Deprecated: Do not use.
func (c *wavelengthRouterClient) CancelAdjustPSD(ctx context.Context, in *AdjustPSDRequest, opts ...grpc.CallOption) (*CancelAdjustPSDResponse, error) {
	out := new(CancelAdjustPSDResponse)
	err := c.cc.Invoke(ctx, "/gnoi.optical.WavelengthRouter/CancelAdjustPSD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wavelengthRouterClient) AdjustSpectrum(ctx context.Context, in *AdjustSpectrumRequest, opts ...grpc.CallOption) (WavelengthRouter_AdjustSpectrumClient, error) {
	stream, err := c.cc.NewStream(ctx, &WavelengthRouter_ServiceDesc.Streams[1], "/gnoi.optical.WavelengthRouter/AdjustSpectrum", opts...)
	if err != nil {
		return nil, err
	}
	x := &wavelengthRouterAdjustSpectrumClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WavelengthRouter_AdjustSpectrumClient interface {
	Recv() (*AdjustSpectrumResponse, error)
	grpc.ClientStream
}

type wavelengthRouterAdjustSpectrumClient struct {
	grpc.ClientStream
}

func (x *wavelengthRouterAdjustSpectrumClient) Recv() (*AdjustSpectrumResponse, error) {
	m := new(AdjustSpectrumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *wavelengthRouterClient) CancelAdjustSpectrum(ctx context.Context, in *AdjustSpectrumRequest, opts ...grpc.CallOption) (*CancelAdjustSpectrumResponse, error) {
	out := new(CancelAdjustSpectrumResponse)
	err := c.cc.Invoke(ctx, "/gnoi.optical.WavelengthRouter/CancelAdjustSpectrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WavelengthRouterServer is the server API for WavelengthRouter service.
// All implementations must embed UnimplementedWavelengthRouterServer
// for forward compatibility
type WavelengthRouterServer interface {
	// Deprecated: Do not use.
	// AdjustPSD performs a power spectral density (PSD) adjustment on an optical
	// DWDM line system port.
	//
	// This RPC's initial message triggers adjustments. The target should validate
	// the request, and immediately begin the adjustment as long as another
	// adjustment is not already running. As this RPC may take multiple minutes,
	// the state should be regularly streamed to the client. The operation should
	// be completed on the device regardless of the state of the RPC channel to
	// the initiating client. A client that becomes disconnected during an
	// operation can reconnect, requesting the same operation, and receive status
	// updates this way. There is no means by which an adjustment that has been
	// triggered by a client can be cancelled unless the CancelAdjustPSD RPC is
	// used.
	// DEPRECATED, use AdjustSpectrum.
	AdjustPSD(*AdjustPSDRequest, WavelengthRouter_AdjustPSDServer) error
	// Deprecated: Do not use.
	// CancelAdjustPSD cancels an in-progress AdjustPSD request. If an
	// adjustment is not being performed for the provided request, then the RPC
	// should return successfully.
	// DEPRECATED, use CancelAdjustSpectrum.
	CancelAdjustPSD(context.Context, *AdjustPSDRequest) (*CancelAdjustPSDResponse, error)
	// AdjustSpectrum performs a spectrum power adjustment on an optical
	// DWDM line system port.
	//
	// This RPC's initial message triggers adjustments. The target should validate
	// the request, and immediately begin the adjustment as long as another
	// adjustment is not already running. As this RPC may take multiple minutes,
	// the state should be regularly streamed to the client. The operation should
	// be completed on the device regardless of the state of the RPC channel to
	// the initiating client. A client that becomes disconnected during an
	// operation can reconnect, requesting the same operation, and receive status
	// updates this way. There is no means by which an adjustment that has been
	// triggered by a client can be cancelled unless the CancelAdjustSpectrum
	// RPC is used.
	AdjustSpectrum(*AdjustSpectrumRequest, WavelengthRouter_AdjustSpectrumServer) error
	// CancelAdjustSpectrum cancels an in-progress AdjustSpectrum request. If an
	// adjustment is not being performed for the provided request, then the RPC
	// should return successfully.
	CancelAdjustSpectrum(context.Context, *AdjustSpectrumRequest) (*CancelAdjustSpectrumResponse, error)
	mustEmbedUnimplementedWavelengthRouterServer()
}

// UnimplementedWavelengthRouterServer must be embedded to have forward compatible implementations.
type UnimplementedWavelengthRouterServer struct {
}

func (UnimplementedWavelengthRouterServer) AdjustPSD(*AdjustPSDRequest, WavelengthRouter_AdjustPSDServer) error {
	return status.Errorf(codes.Unimplemented, "method AdjustPSD not implemented")
}
func (UnimplementedWavelengthRouterServer) CancelAdjustPSD(context.Context, *AdjustPSDRequest) (*CancelAdjustPSDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAdjustPSD not implemented")
}
func (UnimplementedWavelengthRouterServer) AdjustSpectrum(*AdjustSpectrumRequest, WavelengthRouter_AdjustSpectrumServer) error {
	return status.Errorf(codes.Unimplemented, "method AdjustSpectrum not implemented")
}
func (UnimplementedWavelengthRouterServer) CancelAdjustSpectrum(context.Context, *AdjustSpectrumRequest) (*CancelAdjustSpectrumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAdjustSpectrum not implemented")
}
func (UnimplementedWavelengthRouterServer) mustEmbedUnimplementedWavelengthRouterServer() {}

// UnsafeWavelengthRouterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WavelengthRouterServer will
// result in compilation errors.
type UnsafeWavelengthRouterServer interface {
	mustEmbedUnimplementedWavelengthRouterServer()
}

func RegisterWavelengthRouterServer(s grpc.ServiceRegistrar, srv WavelengthRouterServer) {
	s.RegisterService(&WavelengthRouter_ServiceDesc, srv)
}

func _WavelengthRouter_AdjustPSD_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AdjustPSDRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WavelengthRouterServer).AdjustPSD(m, &wavelengthRouterAdjustPSDServer{stream})
}

type WavelengthRouter_AdjustPSDServer interface {
	Send(*AdjustPSDResponse) error
	grpc.ServerStream
}

type wavelengthRouterAdjustPSDServer struct {
	grpc.ServerStream
}

func (x *wavelengthRouterAdjustPSDServer) Send(m *AdjustPSDResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WavelengthRouter_CancelAdjustPSD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustPSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WavelengthRouterServer).CancelAdjustPSD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.optical.WavelengthRouter/CancelAdjustPSD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WavelengthRouterServer).CancelAdjustPSD(ctx, req.(*AdjustPSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WavelengthRouter_AdjustSpectrum_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AdjustSpectrumRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WavelengthRouterServer).AdjustSpectrum(m, &wavelengthRouterAdjustSpectrumServer{stream})
}

type WavelengthRouter_AdjustSpectrumServer interface {
	Send(*AdjustSpectrumResponse) error
	grpc.ServerStream
}

type wavelengthRouterAdjustSpectrumServer struct {
	grpc.ServerStream
}

func (x *wavelengthRouterAdjustSpectrumServer) Send(m *AdjustSpectrumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WavelengthRouter_CancelAdjustSpectrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustSpectrumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WavelengthRouterServer).CancelAdjustSpectrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.optical.WavelengthRouter/CancelAdjustSpectrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WavelengthRouterServer).CancelAdjustSpectrum(ctx, req.(*AdjustSpectrumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WavelengthRouter_ServiceDesc is the grpc.ServiceDesc for WavelengthRouter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WavelengthRouter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.optical.WavelengthRouter",
	HandlerType: (*WavelengthRouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelAdjustPSD",
			Handler:    _WavelengthRouter_CancelAdjustPSD_Handler,
		},
		{
			MethodName: "CancelAdjustSpectrum",
			Handler:    _WavelengthRouter_CancelAdjustSpectrum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AdjustPSD",
			Handler:       _WavelengthRouter_AdjustPSD_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AdjustSpectrum",
			Handler:       _WavelengthRouter_AdjustSpectrum_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "wavelength_router/wavelength_router.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kappa/services/payment_gateway/payment_gateway.proto

package payment_gateway

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

// PaymentGatewayServiceClient is the client API for PaymentGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentGatewayServiceClient interface {
	GetPaymentGatewayByName(ctx context.Context, in *GetPaymentGatewayByNameRequest, opts ...grpc.CallOption) (*GetPaymentGatewayByNameResponse, error)
	GetPaymentGatewayForSource(ctx context.Context, in *GetPaymentGatewayForSourceRequest, opts ...grpc.CallOption) (*GetPaymentGatewayForSourceResponse, error)
	GetPaymentGatewayForDestination(ctx context.Context, in *GetPaymentGatewayForDestinationRequest, opts ...grpc.CallOption) (*GetPaymentGatewayForDestinationResponse, error)
	GetPaymentGatewayFeeSchedule(ctx context.Context, in *GetPaymentGatewayFeeScheduleRequest, opts ...grpc.CallOption) (*GetPaymentGatewayFeeScheduleResponse, error)
	GetBankAccountForPaymentGateway(ctx context.Context, in *GetBankAccountForPaymentGatewayRequest, opts ...grpc.CallOption) (*GetBankAccountForPaymentGatewayResponse, error)
	GetPaymentGatewayByID(ctx context.Context, in *GetPaymentGatewayByIDRequest, opts ...grpc.CallOption) (*GetPaymentGatewayByIDResponse, error)
	// Get the processing schedule for a given gateway, direction, country and currency
	GetPaymentGatewayProcessingSchedule(ctx context.Context, in *GetPaymentGatewayProcessingScheduleRequest, opts ...grpc.CallOption) (*GetPaymentGatewayProcessingScheduleResponse, error)
}

type paymentGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentGatewayServiceClient(cc grpc.ClientConnInterface) PaymentGatewayServiceClient {
	return &paymentGatewayServiceClient{cc}
}

func (c *paymentGatewayServiceClient) GetPaymentGatewayByName(ctx context.Context, in *GetPaymentGatewayByNameRequest, opts ...grpc.CallOption) (*GetPaymentGatewayByNameResponse, error) {
	out := new(GetPaymentGatewayByNameResponse)
	err := c.cc.Invoke(ctx, "/payment_gateway.PaymentGatewayService/GetPaymentGatewayByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayServiceClient) GetPaymentGatewayForSource(ctx context.Context, in *GetPaymentGatewayForSourceRequest, opts ...grpc.CallOption) (*GetPaymentGatewayForSourceResponse, error) {
	out := new(GetPaymentGatewayForSourceResponse)
	err := c.cc.Invoke(ctx, "/payment_gateway.PaymentGatewayService/GetPaymentGatewayForSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayServiceClient) GetPaymentGatewayForDestination(ctx context.Context, in *GetPaymentGatewayForDestinationRequest, opts ...grpc.CallOption) (*GetPaymentGatewayForDestinationResponse, error) {
	out := new(GetPaymentGatewayForDestinationResponse)
	err := c.cc.Invoke(ctx, "/payment_gateway.PaymentGatewayService/GetPaymentGatewayForDestination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayServiceClient) GetPaymentGatewayFeeSchedule(ctx context.Context, in *GetPaymentGatewayFeeScheduleRequest, opts ...grpc.CallOption) (*GetPaymentGatewayFeeScheduleResponse, error) {
	out := new(GetPaymentGatewayFeeScheduleResponse)
	err := c.cc.Invoke(ctx, "/payment_gateway.PaymentGatewayService/GetPaymentGatewayFeeSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayServiceClient) GetBankAccountForPaymentGateway(ctx context.Context, in *GetBankAccountForPaymentGatewayRequest, opts ...grpc.CallOption) (*GetBankAccountForPaymentGatewayResponse, error) {
	out := new(GetBankAccountForPaymentGatewayResponse)
	err := c.cc.Invoke(ctx, "/payment_gateway.PaymentGatewayService/GetBankAccountForPaymentGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayServiceClient) GetPaymentGatewayByID(ctx context.Context, in *GetPaymentGatewayByIDRequest, opts ...grpc.CallOption) (*GetPaymentGatewayByIDResponse, error) {
	out := new(GetPaymentGatewayByIDResponse)
	err := c.cc.Invoke(ctx, "/payment_gateway.PaymentGatewayService/GetPaymentGatewayByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayServiceClient) GetPaymentGatewayProcessingSchedule(ctx context.Context, in *GetPaymentGatewayProcessingScheduleRequest, opts ...grpc.CallOption) (*GetPaymentGatewayProcessingScheduleResponse, error) {
	out := new(GetPaymentGatewayProcessingScheduleResponse)
	err := c.cc.Invoke(ctx, "/payment_gateway.PaymentGatewayService/GetPaymentGatewayProcessingSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentGatewayServiceServer is the server API for PaymentGatewayService service.
// All implementations must embed UnimplementedPaymentGatewayServiceServer
// for forward compatibility
type PaymentGatewayServiceServer interface {
	GetPaymentGatewayByName(context.Context, *GetPaymentGatewayByNameRequest) (*GetPaymentGatewayByNameResponse, error)
	GetPaymentGatewayForSource(context.Context, *GetPaymentGatewayForSourceRequest) (*GetPaymentGatewayForSourceResponse, error)
	GetPaymentGatewayForDestination(context.Context, *GetPaymentGatewayForDestinationRequest) (*GetPaymentGatewayForDestinationResponse, error)
	GetPaymentGatewayFeeSchedule(context.Context, *GetPaymentGatewayFeeScheduleRequest) (*GetPaymentGatewayFeeScheduleResponse, error)
	GetBankAccountForPaymentGateway(context.Context, *GetBankAccountForPaymentGatewayRequest) (*GetBankAccountForPaymentGatewayResponse, error)
	GetPaymentGatewayByID(context.Context, *GetPaymentGatewayByIDRequest) (*GetPaymentGatewayByIDResponse, error)
	// Get the processing schedule for a given gateway, direction, country and currency
	GetPaymentGatewayProcessingSchedule(context.Context, *GetPaymentGatewayProcessingScheduleRequest) (*GetPaymentGatewayProcessingScheduleResponse, error)
	mustEmbedUnimplementedPaymentGatewayServiceServer()
}

// UnimplementedPaymentGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentGatewayServiceServer struct {
}

func (UnimplementedPaymentGatewayServiceServer) GetPaymentGatewayByName(context.Context, *GetPaymentGatewayByNameRequest) (*GetPaymentGatewayByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentGatewayByName not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) GetPaymentGatewayForSource(context.Context, *GetPaymentGatewayForSourceRequest) (*GetPaymentGatewayForSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentGatewayForSource not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) GetPaymentGatewayForDestination(context.Context, *GetPaymentGatewayForDestinationRequest) (*GetPaymentGatewayForDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentGatewayForDestination not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) GetPaymentGatewayFeeSchedule(context.Context, *GetPaymentGatewayFeeScheduleRequest) (*GetPaymentGatewayFeeScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentGatewayFeeSchedule not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) GetBankAccountForPaymentGateway(context.Context, *GetBankAccountForPaymentGatewayRequest) (*GetBankAccountForPaymentGatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankAccountForPaymentGateway not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) GetPaymentGatewayByID(context.Context, *GetPaymentGatewayByIDRequest) (*GetPaymentGatewayByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentGatewayByID not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) GetPaymentGatewayProcessingSchedule(context.Context, *GetPaymentGatewayProcessingScheduleRequest) (*GetPaymentGatewayProcessingScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentGatewayProcessingSchedule not implemented")
}
func (UnimplementedPaymentGatewayServiceServer) mustEmbedUnimplementedPaymentGatewayServiceServer() {}

// UnsafePaymentGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentGatewayServiceServer will
// result in compilation errors.
type UnsafePaymentGatewayServiceServer interface {
	mustEmbedUnimplementedPaymentGatewayServiceServer()
}

func RegisterPaymentGatewayServiceServer(s grpc.ServiceRegistrar, srv PaymentGatewayServiceServer) {
	s.RegisterService(&PaymentGatewayService_ServiceDesc, srv)
}

func _PaymentGatewayService_GetPaymentGatewayByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentGatewayByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment_gateway.PaymentGatewayService/GetPaymentGatewayByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayByName(ctx, req.(*GetPaymentGatewayByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGatewayService_GetPaymentGatewayForSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentGatewayForSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayForSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment_gateway.PaymentGatewayService/GetPaymentGatewayForSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayForSource(ctx, req.(*GetPaymentGatewayForSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGatewayService_GetPaymentGatewayForDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentGatewayForDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayForDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment_gateway.PaymentGatewayService/GetPaymentGatewayForDestination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayForDestination(ctx, req.(*GetPaymentGatewayForDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGatewayService_GetPaymentGatewayFeeSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentGatewayFeeScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayFeeSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment_gateway.PaymentGatewayService/GetPaymentGatewayFeeSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayFeeSchedule(ctx, req.(*GetPaymentGatewayFeeScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGatewayService_GetBankAccountForPaymentGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankAccountForPaymentGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetBankAccountForPaymentGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment_gateway.PaymentGatewayService/GetBankAccountForPaymentGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetBankAccountForPaymentGateway(ctx, req.(*GetBankAccountForPaymentGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGatewayService_GetPaymentGatewayByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentGatewayByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment_gateway.PaymentGatewayService/GetPaymentGatewayByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayByID(ctx, req.(*GetPaymentGatewayByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGatewayService_GetPaymentGatewayProcessingSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentGatewayProcessingScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayProcessingSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment_gateway.PaymentGatewayService/GetPaymentGatewayProcessingSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServiceServer).GetPaymentGatewayProcessingSchedule(ctx, req.(*GetPaymentGatewayProcessingScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentGatewayService_ServiceDesc is the grpc.ServiceDesc for PaymentGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment_gateway.PaymentGatewayService",
	HandlerType: (*PaymentGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPaymentGatewayByName",
			Handler:    _PaymentGatewayService_GetPaymentGatewayByName_Handler,
		},
		{
			MethodName: "GetPaymentGatewayForSource",
			Handler:    _PaymentGatewayService_GetPaymentGatewayForSource_Handler,
		},
		{
			MethodName: "GetPaymentGatewayForDestination",
			Handler:    _PaymentGatewayService_GetPaymentGatewayForDestination_Handler,
		},
		{
			MethodName: "GetPaymentGatewayFeeSchedule",
			Handler:    _PaymentGatewayService_GetPaymentGatewayFeeSchedule_Handler,
		},
		{
			MethodName: "GetBankAccountForPaymentGateway",
			Handler:    _PaymentGatewayService_GetBankAccountForPaymentGateway_Handler,
		},
		{
			MethodName: "GetPaymentGatewayByID",
			Handler:    _PaymentGatewayService_GetPaymentGatewayByID_Handler,
		},
		{
			MethodName: "GetPaymentGatewayProcessingSchedule",
			Handler:    _PaymentGatewayService_GetPaymentGatewayProcessingSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kappa/services/payment_gateway/payment_gateway.proto",
}
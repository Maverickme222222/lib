// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kappa/services/limits/limits.proto

package limits

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

// LimitsServiceClient is the client API for LimitsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LimitsServiceClient interface {
	GetLimitByID(ctx context.Context, in *GetLimitByIDRequest, opts ...grpc.CallOption) (*GetLimitByIDResponse, error)
	CreateLimit(ctx context.Context, in *CreateLimitRequest, opts ...grpc.CallOption) (*CreateLimitResponse, error)
	UpdateLimit(ctx context.Context, in *UpdateLimitRequest, opts ...grpc.CallOption) (*UpdateLimitResponse, error)
	DeleteLimit(ctx context.Context, in *DeleteLimitRequest, opts ...grpc.CallOption) (*DeleteLimitResponse, error)
	GetDefaultLimits(ctx context.Context, in *GetDefaultLimitsRequest, opts ...grpc.CallOption) (*GetDefaultLimitsResponse, error)
	CreateLimitsForEntity(ctx context.Context, in *CreateLimitsForEntityRequest, opts ...grpc.CallOption) (*CreateLimitsForEntityResponse, error)
	GetLimitsByEntityID(ctx context.Context, in *GetLimitsByEntityIDRequest, opts ...grpc.CallOption) (*GetLimitsByEntityIDResponse, error)
	CheckTransaction(ctx context.Context, in *CheckTransactionRequest, opts ...grpc.CallOption) (*CheckTransactionResponse, error)
	CheckTransactionQuote(ctx context.Context, in *CheckTransactionQuoteRequest, opts ...grpc.CallOption) (*CheckTransactionQuoteResponse, error)
}

type limitsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLimitsServiceClient(cc grpc.ClientConnInterface) LimitsServiceClient {
	return &limitsServiceClient{cc}
}

func (c *limitsServiceClient) GetLimitByID(ctx context.Context, in *GetLimitByIDRequest, opts ...grpc.CallOption) (*GetLimitByIDResponse, error) {
	out := new(GetLimitByIDResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/GetLimitByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitsServiceClient) CreateLimit(ctx context.Context, in *CreateLimitRequest, opts ...grpc.CallOption) (*CreateLimitResponse, error) {
	out := new(CreateLimitResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/CreateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitsServiceClient) UpdateLimit(ctx context.Context, in *UpdateLimitRequest, opts ...grpc.CallOption) (*UpdateLimitResponse, error) {
	out := new(UpdateLimitResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/UpdateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitsServiceClient) DeleteLimit(ctx context.Context, in *DeleteLimitRequest, opts ...grpc.CallOption) (*DeleteLimitResponse, error) {
	out := new(DeleteLimitResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/DeleteLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitsServiceClient) GetDefaultLimits(ctx context.Context, in *GetDefaultLimitsRequest, opts ...grpc.CallOption) (*GetDefaultLimitsResponse, error) {
	out := new(GetDefaultLimitsResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/GetDefaultLimits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitsServiceClient) CreateLimitsForEntity(ctx context.Context, in *CreateLimitsForEntityRequest, opts ...grpc.CallOption) (*CreateLimitsForEntityResponse, error) {
	out := new(CreateLimitsForEntityResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/CreateLimitsForEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitsServiceClient) GetLimitsByEntityID(ctx context.Context, in *GetLimitsByEntityIDRequest, opts ...grpc.CallOption) (*GetLimitsByEntityIDResponse, error) {
	out := new(GetLimitsByEntityIDResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/GetLimitsByEntityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitsServiceClient) CheckTransaction(ctx context.Context, in *CheckTransactionRequest, opts ...grpc.CallOption) (*CheckTransactionResponse, error) {
	out := new(CheckTransactionResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/CheckTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *limitsServiceClient) CheckTransactionQuote(ctx context.Context, in *CheckTransactionQuoteRequest, opts ...grpc.CallOption) (*CheckTransactionQuoteResponse, error) {
	out := new(CheckTransactionQuoteResponse)
	err := c.cc.Invoke(ctx, "/limits.LimitsService/CheckTransactionQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LimitsServiceServer is the server API for LimitsService service.
// All implementations must embed UnimplementedLimitsServiceServer
// for forward compatibility
type LimitsServiceServer interface {
	GetLimitByID(context.Context, *GetLimitByIDRequest) (*GetLimitByIDResponse, error)
	CreateLimit(context.Context, *CreateLimitRequest) (*CreateLimitResponse, error)
	UpdateLimit(context.Context, *UpdateLimitRequest) (*UpdateLimitResponse, error)
	DeleteLimit(context.Context, *DeleteLimitRequest) (*DeleteLimitResponse, error)
	GetDefaultLimits(context.Context, *GetDefaultLimitsRequest) (*GetDefaultLimitsResponse, error)
	CreateLimitsForEntity(context.Context, *CreateLimitsForEntityRequest) (*CreateLimitsForEntityResponse, error)
	GetLimitsByEntityID(context.Context, *GetLimitsByEntityIDRequest) (*GetLimitsByEntityIDResponse, error)
	CheckTransaction(context.Context, *CheckTransactionRequest) (*CheckTransactionResponse, error)
	CheckTransactionQuote(context.Context, *CheckTransactionQuoteRequest) (*CheckTransactionQuoteResponse, error)
	mustEmbedUnimplementedLimitsServiceServer()
}

// UnimplementedLimitsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLimitsServiceServer struct {
}

func (UnimplementedLimitsServiceServer) GetLimitByID(context.Context, *GetLimitByIDRequest) (*GetLimitByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimitByID not implemented")
}
func (UnimplementedLimitsServiceServer) CreateLimit(context.Context, *CreateLimitRequest) (*CreateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLimit not implemented")
}
func (UnimplementedLimitsServiceServer) UpdateLimit(context.Context, *UpdateLimitRequest) (*UpdateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLimit not implemented")
}
func (UnimplementedLimitsServiceServer) DeleteLimit(context.Context, *DeleteLimitRequest) (*DeleteLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLimit not implemented")
}
func (UnimplementedLimitsServiceServer) GetDefaultLimits(context.Context, *GetDefaultLimitsRequest) (*GetDefaultLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultLimits not implemented")
}
func (UnimplementedLimitsServiceServer) CreateLimitsForEntity(context.Context, *CreateLimitsForEntityRequest) (*CreateLimitsForEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLimitsForEntity not implemented")
}
func (UnimplementedLimitsServiceServer) GetLimitsByEntityID(context.Context, *GetLimitsByEntityIDRequest) (*GetLimitsByEntityIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimitsByEntityID not implemented")
}
func (UnimplementedLimitsServiceServer) CheckTransaction(context.Context, *CheckTransactionRequest) (*CheckTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTransaction not implemented")
}
func (UnimplementedLimitsServiceServer) CheckTransactionQuote(context.Context, *CheckTransactionQuoteRequest) (*CheckTransactionQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTransactionQuote not implemented")
}
func (UnimplementedLimitsServiceServer) mustEmbedUnimplementedLimitsServiceServer() {}

// UnsafeLimitsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LimitsServiceServer will
// result in compilation errors.
type UnsafeLimitsServiceServer interface {
	mustEmbedUnimplementedLimitsServiceServer()
}

func RegisterLimitsServiceServer(s grpc.ServiceRegistrar, srv LimitsServiceServer) {
	s.RegisterService(&LimitsService_ServiceDesc, srv)
}

func _LimitsService_GetLimitByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLimitByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).GetLimitByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/GetLimitByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).GetLimitByID(ctx, req.(*GetLimitByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitsService_CreateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).CreateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/CreateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).CreateLimit(ctx, req.(*CreateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitsService_UpdateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).UpdateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/UpdateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).UpdateLimit(ctx, req.(*UpdateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitsService_DeleteLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).DeleteLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/DeleteLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).DeleteLimit(ctx, req.(*DeleteLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitsService_GetDefaultLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).GetDefaultLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/GetDefaultLimits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).GetDefaultLimits(ctx, req.(*GetDefaultLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitsService_CreateLimitsForEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLimitsForEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).CreateLimitsForEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/CreateLimitsForEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).CreateLimitsForEntity(ctx, req.(*CreateLimitsForEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitsService_GetLimitsByEntityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLimitsByEntityIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).GetLimitsByEntityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/GetLimitsByEntityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).GetLimitsByEntityID(ctx, req.(*GetLimitsByEntityIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitsService_CheckTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).CheckTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/CheckTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).CheckTransaction(ctx, req.(*CheckTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LimitsService_CheckTransactionQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTransactionQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitsServiceServer).CheckTransactionQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/limits.LimitsService/CheckTransactionQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitsServiceServer).CheckTransactionQuote(ctx, req.(*CheckTransactionQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LimitsService_ServiceDesc is the grpc.ServiceDesc for LimitsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LimitsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "limits.LimitsService",
	HandlerType: (*LimitsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLimitByID",
			Handler:    _LimitsService_GetLimitByID_Handler,
		},
		{
			MethodName: "CreateLimit",
			Handler:    _LimitsService_CreateLimit_Handler,
		},
		{
			MethodName: "UpdateLimit",
			Handler:    _LimitsService_UpdateLimit_Handler,
		},
		{
			MethodName: "DeleteLimit",
			Handler:    _LimitsService_DeleteLimit_Handler,
		},
		{
			MethodName: "GetDefaultLimits",
			Handler:    _LimitsService_GetDefaultLimits_Handler,
		},
		{
			MethodName: "CreateLimitsForEntity",
			Handler:    _LimitsService_CreateLimitsForEntity_Handler,
		},
		{
			MethodName: "GetLimitsByEntityID",
			Handler:    _LimitsService_GetLimitsByEntityID_Handler,
		},
		{
			MethodName: "CheckTransaction",
			Handler:    _LimitsService_CheckTransaction_Handler,
		},
		{
			MethodName: "CheckTransactionQuote",
			Handler:    _LimitsService_CheckTransactionQuote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kappa/services/limits/limits.proto",
}

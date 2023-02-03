// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kappa/services/destinations/destinations.proto

package destinations

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

// DestinationsServiceClient is the client API for DestinationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DestinationsServiceClient interface {
	GetDestinationByID(ctx context.Context, in *GetDestinationByIDRequest, opts ...grpc.CallOption) (*GetDestinationByIDResponse, error)
	// Here "DestinationID" refers to the underlying destination. e.g., the wallet's ID or the linked account ID.
	GetDestinationByEntityIDAndDestinationID(ctx context.Context, in *GetDestinationByEntityIDAndDestinationIDRequest, opts ...grpc.CallOption) (*GetDestinationByEntityIDAndDestinationIDResponse, error)
	CreateDestination(ctx context.Context, in *CreateDestinationRequest, opts ...grpc.CallOption) (*CreateDestinationResponse, error)
	UpdateDestination(ctx context.Context, in *UpdateDestinationRequest, opts ...grpc.CallOption) (*UpdateDestinationResponse, error)
	DeleteDestination(ctx context.Context, in *DeleteDestinationRequest, opts ...grpc.CallOption) (*DeleteDestinationResponse, error)
	GetDepositDestinationsByEntityID(ctx context.Context, in *GetDepositDestinationsByEntityIDRequest, opts ...grpc.CallOption) (*GetDepositDestinationsByEntityIDResponse, error)
	GetTransferDestinationsByEntityID(ctx context.Context, in *GetTransferDestinationsByEntityIDRequest, opts ...grpc.CallOption) (*GetTransferDestinationsByEntityIDResponse, error)
	GetWithdrawalDestinationsByEntityID(ctx context.Context, in *GetWithdrawalDestinationsByEntityIDRequest, opts ...grpc.CallOption) (*GetWithdrawalDestinationsByEntityIDResponse, error)
	GetDestinationsByDomicileAndType(ctx context.Context, in *GetDestinationsByDomicileAndTypeRequest, opts ...grpc.CallOption) (*GetDestinationsByDomicileAndTypeResponse, error)
	SetStatus(ctx context.Context, in *SetStatusRequest, opts ...grpc.CallOption) (*SetStatusResponse, error)
}

type destinationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDestinationsServiceClient(cc grpc.ClientConnInterface) DestinationsServiceClient {
	return &destinationsServiceClient{cc}
}

func (c *destinationsServiceClient) GetDestinationByID(ctx context.Context, in *GetDestinationByIDRequest, opts ...grpc.CallOption) (*GetDestinationByIDResponse, error) {
	out := new(GetDestinationByIDResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/GetDestinationByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) GetDestinationByEntityIDAndDestinationID(ctx context.Context, in *GetDestinationByEntityIDAndDestinationIDRequest, opts ...grpc.CallOption) (*GetDestinationByEntityIDAndDestinationIDResponse, error) {
	out := new(GetDestinationByEntityIDAndDestinationIDResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/GetDestinationByEntityIDAndDestinationID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) CreateDestination(ctx context.Context, in *CreateDestinationRequest, opts ...grpc.CallOption) (*CreateDestinationResponse, error) {
	out := new(CreateDestinationResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/CreateDestination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) UpdateDestination(ctx context.Context, in *UpdateDestinationRequest, opts ...grpc.CallOption) (*UpdateDestinationResponse, error) {
	out := new(UpdateDestinationResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/UpdateDestination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) DeleteDestination(ctx context.Context, in *DeleteDestinationRequest, opts ...grpc.CallOption) (*DeleteDestinationResponse, error) {
	out := new(DeleteDestinationResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/DeleteDestination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) GetDepositDestinationsByEntityID(ctx context.Context, in *GetDepositDestinationsByEntityIDRequest, opts ...grpc.CallOption) (*GetDepositDestinationsByEntityIDResponse, error) {
	out := new(GetDepositDestinationsByEntityIDResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/GetDepositDestinationsByEntityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) GetTransferDestinationsByEntityID(ctx context.Context, in *GetTransferDestinationsByEntityIDRequest, opts ...grpc.CallOption) (*GetTransferDestinationsByEntityIDResponse, error) {
	out := new(GetTransferDestinationsByEntityIDResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/GetTransferDestinationsByEntityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) GetWithdrawalDestinationsByEntityID(ctx context.Context, in *GetWithdrawalDestinationsByEntityIDRequest, opts ...grpc.CallOption) (*GetWithdrawalDestinationsByEntityIDResponse, error) {
	out := new(GetWithdrawalDestinationsByEntityIDResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/GetWithdrawalDestinationsByEntityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) GetDestinationsByDomicileAndType(ctx context.Context, in *GetDestinationsByDomicileAndTypeRequest, opts ...grpc.CallOption) (*GetDestinationsByDomicileAndTypeResponse, error) {
	out := new(GetDestinationsByDomicileAndTypeResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/GetDestinationsByDomicileAndType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationsServiceClient) SetStatus(ctx context.Context, in *SetStatusRequest, opts ...grpc.CallOption) (*SetStatusResponse, error) {
	out := new(SetStatusResponse)
	err := c.cc.Invoke(ctx, "/destinations.DestinationsService/SetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DestinationsServiceServer is the server API for DestinationsService service.
// All implementations must embed UnimplementedDestinationsServiceServer
// for forward compatibility
type DestinationsServiceServer interface {
	GetDestinationByID(context.Context, *GetDestinationByIDRequest) (*GetDestinationByIDResponse, error)
	// Here "DestinationID" refers to the underlying destination. e.g., the wallet's ID or the linked account ID.
	GetDestinationByEntityIDAndDestinationID(context.Context, *GetDestinationByEntityIDAndDestinationIDRequest) (*GetDestinationByEntityIDAndDestinationIDResponse, error)
	CreateDestination(context.Context, *CreateDestinationRequest) (*CreateDestinationResponse, error)
	UpdateDestination(context.Context, *UpdateDestinationRequest) (*UpdateDestinationResponse, error)
	DeleteDestination(context.Context, *DeleteDestinationRequest) (*DeleteDestinationResponse, error)
	GetDepositDestinationsByEntityID(context.Context, *GetDepositDestinationsByEntityIDRequest) (*GetDepositDestinationsByEntityIDResponse, error)
	GetTransferDestinationsByEntityID(context.Context, *GetTransferDestinationsByEntityIDRequest) (*GetTransferDestinationsByEntityIDResponse, error)
	GetWithdrawalDestinationsByEntityID(context.Context, *GetWithdrawalDestinationsByEntityIDRequest) (*GetWithdrawalDestinationsByEntityIDResponse, error)
	GetDestinationsByDomicileAndType(context.Context, *GetDestinationsByDomicileAndTypeRequest) (*GetDestinationsByDomicileAndTypeResponse, error)
	SetStatus(context.Context, *SetStatusRequest) (*SetStatusResponse, error)
	mustEmbedUnimplementedDestinationsServiceServer()
}

// UnimplementedDestinationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDestinationsServiceServer struct {
}

func (UnimplementedDestinationsServiceServer) GetDestinationByID(context.Context, *GetDestinationByIDRequest) (*GetDestinationByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationByID not implemented")
}
func (UnimplementedDestinationsServiceServer) GetDestinationByEntityIDAndDestinationID(context.Context, *GetDestinationByEntityIDAndDestinationIDRequest) (*GetDestinationByEntityIDAndDestinationIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationByEntityIDAndDestinationID not implemented")
}
func (UnimplementedDestinationsServiceServer) CreateDestination(context.Context, *CreateDestinationRequest) (*CreateDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDestination not implemented")
}
func (UnimplementedDestinationsServiceServer) UpdateDestination(context.Context, *UpdateDestinationRequest) (*UpdateDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDestination not implemented")
}
func (UnimplementedDestinationsServiceServer) DeleteDestination(context.Context, *DeleteDestinationRequest) (*DeleteDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDestination not implemented")
}
func (UnimplementedDestinationsServiceServer) GetDepositDestinationsByEntityID(context.Context, *GetDepositDestinationsByEntityIDRequest) (*GetDepositDestinationsByEntityIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepositDestinationsByEntityID not implemented")
}
func (UnimplementedDestinationsServiceServer) GetTransferDestinationsByEntityID(context.Context, *GetTransferDestinationsByEntityIDRequest) (*GetTransferDestinationsByEntityIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransferDestinationsByEntityID not implemented")
}
func (UnimplementedDestinationsServiceServer) GetWithdrawalDestinationsByEntityID(context.Context, *GetWithdrawalDestinationsByEntityIDRequest) (*GetWithdrawalDestinationsByEntityIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawalDestinationsByEntityID not implemented")
}
func (UnimplementedDestinationsServiceServer) GetDestinationsByDomicileAndType(context.Context, *GetDestinationsByDomicileAndTypeRequest) (*GetDestinationsByDomicileAndTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationsByDomicileAndType not implemented")
}
func (UnimplementedDestinationsServiceServer) SetStatus(context.Context, *SetStatusRequest) (*SetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (UnimplementedDestinationsServiceServer) mustEmbedUnimplementedDestinationsServiceServer() {}

// UnsafeDestinationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DestinationsServiceServer will
// result in compilation errors.
type UnsafeDestinationsServiceServer interface {
	mustEmbedUnimplementedDestinationsServiceServer()
}

func RegisterDestinationsServiceServer(s grpc.ServiceRegistrar, srv DestinationsServiceServer) {
	s.RegisterService(&DestinationsService_ServiceDesc, srv)
}

func _DestinationsService_GetDestinationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).GetDestinationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/GetDestinationByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).GetDestinationByID(ctx, req.(*GetDestinationByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_GetDestinationByEntityIDAndDestinationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationByEntityIDAndDestinationIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).GetDestinationByEntityIDAndDestinationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/GetDestinationByEntityIDAndDestinationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).GetDestinationByEntityIDAndDestinationID(ctx, req.(*GetDestinationByEntityIDAndDestinationIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_CreateDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).CreateDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/CreateDestination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).CreateDestination(ctx, req.(*CreateDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_UpdateDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).UpdateDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/UpdateDestination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).UpdateDestination(ctx, req.(*UpdateDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_DeleteDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).DeleteDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/DeleteDestination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).DeleteDestination(ctx, req.(*DeleteDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_GetDepositDestinationsByEntityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepositDestinationsByEntityIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).GetDepositDestinationsByEntityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/GetDepositDestinationsByEntityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).GetDepositDestinationsByEntityID(ctx, req.(*GetDepositDestinationsByEntityIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_GetTransferDestinationsByEntityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransferDestinationsByEntityIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).GetTransferDestinationsByEntityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/GetTransferDestinationsByEntityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).GetTransferDestinationsByEntityID(ctx, req.(*GetTransferDestinationsByEntityIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_GetWithdrawalDestinationsByEntityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawalDestinationsByEntityIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).GetWithdrawalDestinationsByEntityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/GetWithdrawalDestinationsByEntityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).GetWithdrawalDestinationsByEntityID(ctx, req.(*GetWithdrawalDestinationsByEntityIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_GetDestinationsByDomicileAndType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationsByDomicileAndTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).GetDestinationsByDomicileAndType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/GetDestinationsByDomicileAndType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).GetDestinationsByDomicileAndType(ctx, req.(*GetDestinationsByDomicileAndTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DestinationsService_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationsServiceServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/destinations.DestinationsService/SetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationsServiceServer).SetStatus(ctx, req.(*SetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DestinationsService_ServiceDesc is the grpc.ServiceDesc for DestinationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DestinationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "destinations.DestinationsService",
	HandlerType: (*DestinationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDestinationByID",
			Handler:    _DestinationsService_GetDestinationByID_Handler,
		},
		{
			MethodName: "GetDestinationByEntityIDAndDestinationID",
			Handler:    _DestinationsService_GetDestinationByEntityIDAndDestinationID_Handler,
		},
		{
			MethodName: "CreateDestination",
			Handler:    _DestinationsService_CreateDestination_Handler,
		},
		{
			MethodName: "UpdateDestination",
			Handler:    _DestinationsService_UpdateDestination_Handler,
		},
		{
			MethodName: "DeleteDestination",
			Handler:    _DestinationsService_DeleteDestination_Handler,
		},
		{
			MethodName: "GetDepositDestinationsByEntityID",
			Handler:    _DestinationsService_GetDepositDestinationsByEntityID_Handler,
		},
		{
			MethodName: "GetTransferDestinationsByEntityID",
			Handler:    _DestinationsService_GetTransferDestinationsByEntityID_Handler,
		},
		{
			MethodName: "GetWithdrawalDestinationsByEntityID",
			Handler:    _DestinationsService_GetWithdrawalDestinationsByEntityID_Handler,
		},
		{
			MethodName: "GetDestinationsByDomicileAndType",
			Handler:    _DestinationsService_GetDestinationsByDomicileAndType_Handler,
		},
		{
			MethodName: "SetStatus",
			Handler:    _DestinationsService_SetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kappa/services/destinations/destinations.proto",
}

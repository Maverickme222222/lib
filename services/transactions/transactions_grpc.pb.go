// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kappa/services/transactions/transactions.proto

package transactions

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

// TransactionsServiceClient is the client API for TransactionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionsServiceClient interface {
	GetTransactionByID(ctx context.Context, in *GetTransactionByIDRequest, opts ...grpc.CallOption) (*GetTransactionByIDResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
	GetTransactionsBySourceIDAndStatuses(ctx context.Context, in *GetTransactionsBySourceIDAndStatusesRequest, opts ...grpc.CallOption) (*GetTransactionsBySourceIDAndStatusesResponse, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error)
	GetTransactionsBySourceIDs(ctx context.Context, in *GetTransactionsBySourceIDsRequest, opts ...grpc.CallOption) (*GetTransactionsBySourceIDsResponse, error)
	GetTransactionsByDestinationIDsTypeAndStatuses(ctx context.Context, in *GetTransactionsByDestinationIDsTypeAndStatusesRequest, opts ...grpc.CallOption) (*GetTransactionsByDestinationIDsTypeAndStatusesResponse, error)
	GetTransactionsBySourceIDsAndStatuses(ctx context.Context, in *GetTransactionsBySourceIDsAndStatusesRequest, opts ...grpc.CallOption) (*GetTransactionsBySourceIDsAndStatusesResponse, error)
	GetTransactionsByEntityID(ctx context.Context, in *GetTransactionsByEntityIDRequest, opts ...grpc.CallOption) (*GetTransactionsByEntityIDResponse, error)
}

type transactionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionsServiceClient(cc grpc.ClientConnInterface) TransactionsServiceClient {
	return &transactionsServiceClient{cc}
}

func (c *transactionsServiceClient) GetTransactionByID(ctx context.Context, in *GetTransactionByIDRequest, opts ...grpc.CallOption) (*GetTransactionByIDResponse, error) {
	out := new(GetTransactionByIDResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/GetTransactionByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) GetTransactionsBySourceIDAndStatuses(ctx context.Context, in *GetTransactionsBySourceIDAndStatusesRequest, opts ...grpc.CallOption) (*GetTransactionsBySourceIDAndStatusesResponse, error) {
	out := new(GetTransactionsBySourceIDAndStatusesResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/GetTransactionsBySourceIDAndStatuses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequest, opts ...grpc.CallOption) (*UpdateTransactionResponse, error) {
	out := new(UpdateTransactionResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/UpdateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) GetTransactionsBySourceIDs(ctx context.Context, in *GetTransactionsBySourceIDsRequest, opts ...grpc.CallOption) (*GetTransactionsBySourceIDsResponse, error) {
	out := new(GetTransactionsBySourceIDsResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/GetTransactionsBySourceIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) GetTransactionsByDestinationIDsTypeAndStatuses(ctx context.Context, in *GetTransactionsByDestinationIDsTypeAndStatusesRequest, opts ...grpc.CallOption) (*GetTransactionsByDestinationIDsTypeAndStatusesResponse, error) {
	out := new(GetTransactionsByDestinationIDsTypeAndStatusesResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/GetTransactionsByDestinationIDsTypeAndStatuses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) GetTransactionsBySourceIDsAndStatuses(ctx context.Context, in *GetTransactionsBySourceIDsAndStatusesRequest, opts ...grpc.CallOption) (*GetTransactionsBySourceIDsAndStatusesResponse, error) {
	out := new(GetTransactionsBySourceIDsAndStatusesResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/GetTransactionsBySourceIDsAndStatuses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionsServiceClient) GetTransactionsByEntityID(ctx context.Context, in *GetTransactionsByEntityIDRequest, opts ...grpc.CallOption) (*GetTransactionsByEntityIDResponse, error) {
	out := new(GetTransactionsByEntityIDResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionsService/GetTransactionsByEntityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionsServiceServer is the server API for TransactionsService service.
// All implementations must embed UnimplementedTransactionsServiceServer
// for forward compatibility
type TransactionsServiceServer interface {
	GetTransactionByID(context.Context, *GetTransactionByIDRequest) (*GetTransactionByIDResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
	GetTransactionsBySourceIDAndStatuses(context.Context, *GetTransactionsBySourceIDAndStatusesRequest) (*GetTransactionsBySourceIDAndStatusesResponse, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error)
	GetTransactionsBySourceIDs(context.Context, *GetTransactionsBySourceIDsRequest) (*GetTransactionsBySourceIDsResponse, error)
	GetTransactionsByDestinationIDsTypeAndStatuses(context.Context, *GetTransactionsByDestinationIDsTypeAndStatusesRequest) (*GetTransactionsByDestinationIDsTypeAndStatusesResponse, error)
	GetTransactionsBySourceIDsAndStatuses(context.Context, *GetTransactionsBySourceIDsAndStatusesRequest) (*GetTransactionsBySourceIDsAndStatusesResponse, error)
	GetTransactionsByEntityID(context.Context, *GetTransactionsByEntityIDRequest) (*GetTransactionsByEntityIDResponse, error)
	mustEmbedUnimplementedTransactionsServiceServer()
}

// UnimplementedTransactionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionsServiceServer struct {
}

func (UnimplementedTransactionsServiceServer) GetTransactionByID(context.Context, *GetTransactionByIDRequest) (*GetTransactionByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionByID not implemented")
}
func (UnimplementedTransactionsServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedTransactionsServiceServer) GetTransactionsBySourceIDAndStatuses(context.Context, *GetTransactionsBySourceIDAndStatusesRequest) (*GetTransactionsBySourceIDAndStatusesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsBySourceIDAndStatuses not implemented")
}
func (UnimplementedTransactionsServiceServer) UpdateTransaction(context.Context, *UpdateTransactionRequest) (*UpdateTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedTransactionsServiceServer) GetTransactionsBySourceIDs(context.Context, *GetTransactionsBySourceIDsRequest) (*GetTransactionsBySourceIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsBySourceIDs not implemented")
}
func (UnimplementedTransactionsServiceServer) GetTransactionsByDestinationIDsTypeAndStatuses(context.Context, *GetTransactionsByDestinationIDsTypeAndStatusesRequest) (*GetTransactionsByDestinationIDsTypeAndStatusesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsByDestinationIDsTypeAndStatuses not implemented")
}
func (UnimplementedTransactionsServiceServer) GetTransactionsBySourceIDsAndStatuses(context.Context, *GetTransactionsBySourceIDsAndStatusesRequest) (*GetTransactionsBySourceIDsAndStatusesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsBySourceIDsAndStatuses not implemented")
}
func (UnimplementedTransactionsServiceServer) GetTransactionsByEntityID(context.Context, *GetTransactionsByEntityIDRequest) (*GetTransactionsByEntityIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsByEntityID not implemented")
}
func (UnimplementedTransactionsServiceServer) mustEmbedUnimplementedTransactionsServiceServer() {}

// UnsafeTransactionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionsServiceServer will
// result in compilation errors.
type UnsafeTransactionsServiceServer interface {
	mustEmbedUnimplementedTransactionsServiceServer()
}

func RegisterTransactionsServiceServer(s grpc.ServiceRegistrar, srv TransactionsServiceServer) {
	s.RegisterService(&TransactionsService_ServiceDesc, srv)
}

func _TransactionsService_GetTransactionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).GetTransactionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/GetTransactionByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).GetTransactionByID(ctx, req.(*GetTransactionByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_GetTransactionsBySourceIDAndStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsBySourceIDAndStatusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).GetTransactionsBySourceIDAndStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/GetTransactionsBySourceIDAndStatuses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).GetTransactionsBySourceIDAndStatuses(ctx, req.(*GetTransactionsBySourceIDAndStatusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/UpdateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_GetTransactionsBySourceIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsBySourceIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).GetTransactionsBySourceIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/GetTransactionsBySourceIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).GetTransactionsBySourceIDs(ctx, req.(*GetTransactionsBySourceIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_GetTransactionsByDestinationIDsTypeAndStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsByDestinationIDsTypeAndStatusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).GetTransactionsByDestinationIDsTypeAndStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/GetTransactionsByDestinationIDsTypeAndStatuses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).GetTransactionsByDestinationIDsTypeAndStatuses(ctx, req.(*GetTransactionsByDestinationIDsTypeAndStatusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_GetTransactionsBySourceIDsAndStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsBySourceIDsAndStatusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).GetTransactionsBySourceIDsAndStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/GetTransactionsBySourceIDsAndStatuses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).GetTransactionsBySourceIDsAndStatuses(ctx, req.(*GetTransactionsBySourceIDsAndStatusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionsService_GetTransactionsByEntityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsByEntityIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionsServiceServer).GetTransactionsByEntityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionsService/GetTransactionsByEntityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionsServiceServer).GetTransactionsByEntityID(ctx, req.(*GetTransactionsByEntityIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionsService_ServiceDesc is the grpc.ServiceDesc for TransactionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.TransactionsService",
	HandlerType: (*TransactionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransactionByID",
			Handler:    _TransactionsService_GetTransactionByID_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionsService_CreateTransaction_Handler,
		},
		{
			MethodName: "GetTransactionsBySourceIDAndStatuses",
			Handler:    _TransactionsService_GetTransactionsBySourceIDAndStatuses_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _TransactionsService_UpdateTransaction_Handler,
		},
		{
			MethodName: "GetTransactionsBySourceIDs",
			Handler:    _TransactionsService_GetTransactionsBySourceIDs_Handler,
		},
		{
			MethodName: "GetTransactionsByDestinationIDsTypeAndStatuses",
			Handler:    _TransactionsService_GetTransactionsByDestinationIDsTypeAndStatuses_Handler,
		},
		{
			MethodName: "GetTransactionsBySourceIDsAndStatuses",
			Handler:    _TransactionsService_GetTransactionsBySourceIDsAndStatuses_Handler,
		},
		{
			MethodName: "GetTransactionsByEntityID",
			Handler:    _TransactionsService_GetTransactionsByEntityID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kappa/services/transactions/transactions.proto",
}
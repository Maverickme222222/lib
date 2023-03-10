// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kappa/services/bank_accounts/bank_accounts.proto

package bank_accounts

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

// BankAccountsServiceClient is the client API for BankAccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankAccountsServiceClient interface {
	GetBankAccountById(ctx context.Context, in *GetBankAccountByIdRequest, opts ...grpc.CallOption) (*GetBankAccountByIdResponse, error)
	CreateBankAccount(ctx context.Context, in *CreateBankAccountRequest, opts ...grpc.CallOption) (*CreateBankAccountResponse, error)
	UpdateBankAccount(ctx context.Context, in *UpdateBankAccountRequest, opts ...grpc.CallOption) (*UpdateBankAccountResponse, error)
	DeleteBankAccount(ctx context.Context, in *DeleteBankAccountRequest, opts ...grpc.CallOption) (*DeleteBankAccountResponse, error)
	GetBankAccountByEntityIdAndName(ctx context.Context, in *GetBankAccountByEntityIdAndNameRequest, opts ...grpc.CallOption) (*GetBankAccountByEntityIdAndNameResponse, error)
}

type bankAccountsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankAccountsServiceClient(cc grpc.ClientConnInterface) BankAccountsServiceClient {
	return &bankAccountsServiceClient{cc}
}

func (c *bankAccountsServiceClient) GetBankAccountById(ctx context.Context, in *GetBankAccountByIdRequest, opts ...grpc.CallOption) (*GetBankAccountByIdResponse, error) {
	out := new(GetBankAccountByIdResponse)
	err := c.cc.Invoke(ctx, "/bank_accounts.BankAccountsService/GetBankAccountById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountsServiceClient) CreateBankAccount(ctx context.Context, in *CreateBankAccountRequest, opts ...grpc.CallOption) (*CreateBankAccountResponse, error) {
	out := new(CreateBankAccountResponse)
	err := c.cc.Invoke(ctx, "/bank_accounts.BankAccountsService/CreateBankAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountsServiceClient) UpdateBankAccount(ctx context.Context, in *UpdateBankAccountRequest, opts ...grpc.CallOption) (*UpdateBankAccountResponse, error) {
	out := new(UpdateBankAccountResponse)
	err := c.cc.Invoke(ctx, "/bank_accounts.BankAccountsService/UpdateBankAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountsServiceClient) DeleteBankAccount(ctx context.Context, in *DeleteBankAccountRequest, opts ...grpc.CallOption) (*DeleteBankAccountResponse, error) {
	out := new(DeleteBankAccountResponse)
	err := c.cc.Invoke(ctx, "/bank_accounts.BankAccountsService/DeleteBankAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankAccountsServiceClient) GetBankAccountByEntityIdAndName(ctx context.Context, in *GetBankAccountByEntityIdAndNameRequest, opts ...grpc.CallOption) (*GetBankAccountByEntityIdAndNameResponse, error) {
	out := new(GetBankAccountByEntityIdAndNameResponse)
	err := c.cc.Invoke(ctx, "/bank_accounts.BankAccountsService/GetBankAccountByEntityIdAndName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankAccountsServiceServer is the server API for BankAccountsService service.
// All implementations must embed UnimplementedBankAccountsServiceServer
// for forward compatibility
type BankAccountsServiceServer interface {
	GetBankAccountById(context.Context, *GetBankAccountByIdRequest) (*GetBankAccountByIdResponse, error)
	CreateBankAccount(context.Context, *CreateBankAccountRequest) (*CreateBankAccountResponse, error)
	UpdateBankAccount(context.Context, *UpdateBankAccountRequest) (*UpdateBankAccountResponse, error)
	DeleteBankAccount(context.Context, *DeleteBankAccountRequest) (*DeleteBankAccountResponse, error)
	GetBankAccountByEntityIdAndName(context.Context, *GetBankAccountByEntityIdAndNameRequest) (*GetBankAccountByEntityIdAndNameResponse, error)
	mustEmbedUnimplementedBankAccountsServiceServer()
}

// UnimplementedBankAccountsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankAccountsServiceServer struct {
}

func (UnimplementedBankAccountsServiceServer) GetBankAccountById(context.Context, *GetBankAccountByIdRequest) (*GetBankAccountByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankAccountById not implemented")
}
func (UnimplementedBankAccountsServiceServer) CreateBankAccount(context.Context, *CreateBankAccountRequest) (*CreateBankAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBankAccount not implemented")
}
func (UnimplementedBankAccountsServiceServer) UpdateBankAccount(context.Context, *UpdateBankAccountRequest) (*UpdateBankAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBankAccount not implemented")
}
func (UnimplementedBankAccountsServiceServer) DeleteBankAccount(context.Context, *DeleteBankAccountRequest) (*DeleteBankAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBankAccount not implemented")
}
func (UnimplementedBankAccountsServiceServer) GetBankAccountByEntityIdAndName(context.Context, *GetBankAccountByEntityIdAndNameRequest) (*GetBankAccountByEntityIdAndNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankAccountByEntityIdAndName not implemented")
}
func (UnimplementedBankAccountsServiceServer) mustEmbedUnimplementedBankAccountsServiceServer() {}

// UnsafeBankAccountsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankAccountsServiceServer will
// result in compilation errors.
type UnsafeBankAccountsServiceServer interface {
	mustEmbedUnimplementedBankAccountsServiceServer()
}

func RegisterBankAccountsServiceServer(s grpc.ServiceRegistrar, srv BankAccountsServiceServer) {
	s.RegisterService(&BankAccountsService_ServiceDesc, srv)
}

func _BankAccountsService_GetBankAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankAccountByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountsServiceServer).GetBankAccountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_accounts.BankAccountsService/GetBankAccountById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountsServiceServer).GetBankAccountById(ctx, req.(*GetBankAccountByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountsService_CreateBankAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBankAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountsServiceServer).CreateBankAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_accounts.BankAccountsService/CreateBankAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountsServiceServer).CreateBankAccount(ctx, req.(*CreateBankAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountsService_UpdateBankAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBankAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountsServiceServer).UpdateBankAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_accounts.BankAccountsService/UpdateBankAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountsServiceServer).UpdateBankAccount(ctx, req.(*UpdateBankAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountsService_DeleteBankAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBankAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountsServiceServer).DeleteBankAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_accounts.BankAccountsService/DeleteBankAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountsServiceServer).DeleteBankAccount(ctx, req.(*DeleteBankAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankAccountsService_GetBankAccountByEntityIdAndName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBankAccountByEntityIdAndNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankAccountsServiceServer).GetBankAccountByEntityIdAndName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bank_accounts.BankAccountsService/GetBankAccountByEntityIdAndName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankAccountsServiceServer).GetBankAccountByEntityIdAndName(ctx, req.(*GetBankAccountByEntityIdAndNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BankAccountsService_ServiceDesc is the grpc.ServiceDesc for BankAccountsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankAccountsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank_accounts.BankAccountsService",
	HandlerType: (*BankAccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBankAccountById",
			Handler:    _BankAccountsService_GetBankAccountById_Handler,
		},
		{
			MethodName: "CreateBankAccount",
			Handler:    _BankAccountsService_CreateBankAccount_Handler,
		},
		{
			MethodName: "UpdateBankAccount",
			Handler:    _BankAccountsService_UpdateBankAccount_Handler,
		},
		{
			MethodName: "DeleteBankAccount",
			Handler:    _BankAccountsService_DeleteBankAccount_Handler,
		},
		{
			MethodName: "GetBankAccountByEntityIdAndName",
			Handler:    _BankAccountsService_GetBankAccountByEntityIdAndName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kappa/services/bank_accounts/bank_accounts.proto",
}

// ------------------------------------------------------------------------------
// <copyright company="Kappa Pay Inc.">
// Copyright (C) Kappa Pay Inc.  All rights reserved.
// Unauthorised copying of this file, via any medium, is strictly prohibited.
// Proprietary and confidential
// </copyright>
// <author>Pankaj Pant</author>
// ------------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: kappa/services/mobile_money_accounts/mobile_money_accounts.proto

package mobile_money_accounts

import (
	models "github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMobileMoneyAccountByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMobileMoneyAccountByIdRequest) Reset() {
	*x = GetMobileMoneyAccountByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMobileMoneyAccountByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMobileMoneyAccountByIdRequest) ProtoMessage() {}

func (x *GetMobileMoneyAccountByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMobileMoneyAccountByIdRequest.ProtoReflect.Descriptor instead.
func (*GetMobileMoneyAccountByIdRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *GetMobileMoneyAccountByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMobileMoneyAccountByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MobileMoneyAccount *models.MobileMoneyAccount `protobuf:"bytes,1,opt,name=mobile_money_account,json=mobileMoneyAccount,proto3" json:"mobile_money_account,omitempty"`
}

func (x *GetMobileMoneyAccountByIdResponse) Reset() {
	*x = GetMobileMoneyAccountByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMobileMoneyAccountByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMobileMoneyAccountByIdResponse) ProtoMessage() {}

func (x *GetMobileMoneyAccountByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMobileMoneyAccountByIdResponse.ProtoReflect.Descriptor instead.
func (*GetMobileMoneyAccountByIdResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *GetMobileMoneyAccountByIdResponse) GetMobileMoneyAccount() *models.MobileMoneyAccount {
	if x != nil {
		return x.MobileMoneyAccount
	}
	return nil
}

type CreateMobileMoneyAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MobileMoneyAccount *models.MobileMoneyAccount `protobuf:"bytes,1,opt,name=mobile_money_account,json=mobileMoneyAccount,proto3" json:"mobile_money_account,omitempty"`
}

func (x *CreateMobileMoneyAccountRequest) Reset() {
	*x = CreateMobileMoneyAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMobileMoneyAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMobileMoneyAccountRequest) ProtoMessage() {}

func (x *CreateMobileMoneyAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMobileMoneyAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateMobileMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMobileMoneyAccountRequest) GetMobileMoneyAccount() *models.MobileMoneyAccount {
	if x != nil {
		return x.MobileMoneyAccount
	}
	return nil
}

type CreateMobileMoneyAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MobileMoneyAccount *models.MobileMoneyAccount `protobuf:"bytes,1,opt,name=mobile_money_account,json=mobileMoneyAccount,proto3" json:"mobile_money_account,omitempty"`
}

func (x *CreateMobileMoneyAccountResponse) Reset() {
	*x = CreateMobileMoneyAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMobileMoneyAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMobileMoneyAccountResponse) ProtoMessage() {}

func (x *CreateMobileMoneyAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMobileMoneyAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateMobileMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMobileMoneyAccountResponse) GetMobileMoneyAccount() *models.MobileMoneyAccount {
	if x != nil {
		return x.MobileMoneyAccount
	}
	return nil
}

type UpdateMobileMoneyAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MobileMoneyAccount *models.MobileMoneyAccount `protobuf:"bytes,1,opt,name=mobile_money_account,json=mobileMoneyAccount,proto3" json:"mobile_money_account,omitempty"`
}

func (x *UpdateMobileMoneyAccountRequest) Reset() {
	*x = UpdateMobileMoneyAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMobileMoneyAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMobileMoneyAccountRequest) ProtoMessage() {}

func (x *UpdateMobileMoneyAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMobileMoneyAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateMobileMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMobileMoneyAccountRequest) GetMobileMoneyAccount() *models.MobileMoneyAccount {
	if x != nil {
		return x.MobileMoneyAccount
	}
	return nil
}

type UpdateMobileMoneyAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MobileMoneyAccount *models.MobileMoneyAccount `protobuf:"bytes,1,opt,name=mobile_money_account,json=mobileMoneyAccount,proto3" json:"mobile_money_account,omitempty"`
}

func (x *UpdateMobileMoneyAccountResponse) Reset() {
	*x = UpdateMobileMoneyAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMobileMoneyAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMobileMoneyAccountResponse) ProtoMessage() {}

func (x *UpdateMobileMoneyAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMobileMoneyAccountResponse.ProtoReflect.Descriptor instead.
func (*UpdateMobileMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateMobileMoneyAccountResponse) GetMobileMoneyAccount() *models.MobileMoneyAccount {
	if x != nil {
		return x.MobileMoneyAccount
	}
	return nil
}

type DeleteMobileMoneyAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMobileMoneyAccountRequest) Reset() {
	*x = DeleteMobileMoneyAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMobileMoneyAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMobileMoneyAccountRequest) ProtoMessage() {}

func (x *DeleteMobileMoneyAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMobileMoneyAccountRequest.ProtoReflect.Descriptor instead.
func (*DeleteMobileMoneyAccountRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteMobileMoneyAccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteMobileMoneyAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMobileMoneyAccountResponse) Reset() {
	*x = DeleteMobileMoneyAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMobileMoneyAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMobileMoneyAccountResponse) ProtoMessage() {}

func (x *DeleteMobileMoneyAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMobileMoneyAccountResponse.ProtoReflect.Descriptor instead.
func (*DeleteMobileMoneyAccountResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteMobileMoneyAccountResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_kappa_services_mobile_money_accounts_mobile_money_accounts_proto protoreflect.FileDescriptor

var file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDesc = []byte{
	0x0a, 0x40, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x20, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x70, 0x0a,
	0x21, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x12, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x6e, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4b, 0x0a, 0x14, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x12, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x6f, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x12, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x6e, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x14, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x12, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x6f, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x12, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x31, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xdf, 0x04, 0x0a, 0x1a, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x90, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x37, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38,
	0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8d, 0x01, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x37, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8d, 0x01, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x37, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8d, 0x01, 0x0a, 0x18, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x37, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x70, 0x61,
	0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescOnce sync.Once
	file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescData = file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDesc
)

func file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescGZIP() []byte {
	file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescOnce.Do(func() {
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescData)
	})
	return file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDescData
}

var file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_goTypes = []interface{}{
	(*GetMobileMoneyAccountByIdRequest)(nil),  // 0: mobile_money_accounts.GetMobileMoneyAccountByIdRequest
	(*GetMobileMoneyAccountByIdResponse)(nil), // 1: mobile_money_accounts.GetMobileMoneyAccountByIdResponse
	(*CreateMobileMoneyAccountRequest)(nil),   // 2: mobile_money_accounts.CreateMobileMoneyAccountRequest
	(*CreateMobileMoneyAccountResponse)(nil),  // 3: mobile_money_accounts.CreateMobileMoneyAccountResponse
	(*UpdateMobileMoneyAccountRequest)(nil),   // 4: mobile_money_accounts.UpdateMobileMoneyAccountRequest
	(*UpdateMobileMoneyAccountResponse)(nil),  // 5: mobile_money_accounts.UpdateMobileMoneyAccountResponse
	(*DeleteMobileMoneyAccountRequest)(nil),   // 6: mobile_money_accounts.DeleteMobileMoneyAccountRequest
	(*DeleteMobileMoneyAccountResponse)(nil),  // 7: mobile_money_accounts.DeleteMobileMoneyAccountResponse
	(*models.MobileMoneyAccount)(nil),         // 8: kappa.MobileMoneyAccount
}
var file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_depIdxs = []int32{
	8, // 0: mobile_money_accounts.GetMobileMoneyAccountByIdResponse.mobile_money_account:type_name -> kappa.MobileMoneyAccount
	8, // 1: mobile_money_accounts.CreateMobileMoneyAccountRequest.mobile_money_account:type_name -> kappa.MobileMoneyAccount
	8, // 2: mobile_money_accounts.CreateMobileMoneyAccountResponse.mobile_money_account:type_name -> kappa.MobileMoneyAccount
	8, // 3: mobile_money_accounts.UpdateMobileMoneyAccountRequest.mobile_money_account:type_name -> kappa.MobileMoneyAccount
	8, // 4: mobile_money_accounts.UpdateMobileMoneyAccountResponse.mobile_money_account:type_name -> kappa.MobileMoneyAccount
	0, // 5: mobile_money_accounts.MobileMoneyAccountsService.GetMobileMoneyAccountById:input_type -> mobile_money_accounts.GetMobileMoneyAccountByIdRequest
	2, // 6: mobile_money_accounts.MobileMoneyAccountsService.CreateMobileMoneyAccount:input_type -> mobile_money_accounts.CreateMobileMoneyAccountRequest
	4, // 7: mobile_money_accounts.MobileMoneyAccountsService.UpdateMobileMoneyAccount:input_type -> mobile_money_accounts.UpdateMobileMoneyAccountRequest
	6, // 8: mobile_money_accounts.MobileMoneyAccountsService.DeleteMobileMoneyAccount:input_type -> mobile_money_accounts.DeleteMobileMoneyAccountRequest
	1, // 9: mobile_money_accounts.MobileMoneyAccountsService.GetMobileMoneyAccountById:output_type -> mobile_money_accounts.GetMobileMoneyAccountByIdResponse
	3, // 10: mobile_money_accounts.MobileMoneyAccountsService.CreateMobileMoneyAccount:output_type -> mobile_money_accounts.CreateMobileMoneyAccountResponse
	5, // 11: mobile_money_accounts.MobileMoneyAccountsService.UpdateMobileMoneyAccount:output_type -> mobile_money_accounts.UpdateMobileMoneyAccountResponse
	7, // 12: mobile_money_accounts.MobileMoneyAccountsService.DeleteMobileMoneyAccount:output_type -> mobile_money_accounts.DeleteMobileMoneyAccountResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_init() }
func file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_init() {
	if File_kappa_services_mobile_money_accounts_mobile_money_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMobileMoneyAccountByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMobileMoneyAccountByIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMobileMoneyAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMobileMoneyAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMobileMoneyAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMobileMoneyAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMobileMoneyAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMobileMoneyAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_goTypes,
		DependencyIndexes: file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_depIdxs,
		MessageInfos:      file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_msgTypes,
	}.Build()
	File_kappa_services_mobile_money_accounts_mobile_money_accounts_proto = out.File
	file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_rawDesc = nil
	file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_goTypes = nil
	file_kappa_services_mobile_money_accounts_mobile_money_accounts_proto_depIdxs = nil
}

// ------------------------------------------------------------------------------
// <copyright company="Kappa Pay Inc.">
// Copyright (C) Kappa Pay Inc.  All rights reserved.
// Unauthorised copying of this file, via any medium, is strictly prohibited.
// Proprietary and confidential
// </copyright>
// <author>Andreas Fragner</author>
// ------------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: kappa/common/models/linked_account.proto

package models

import (
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

type LinkedAccountType int32

const (
	LinkedAccountType_LINKED_ACCOUNT_TYPE_UNSPECIFIED          LinkedAccountType = 0
	LinkedAccountType_LINKED_ACCOUNT_TYPE_BANK_ACCOUNT         LinkedAccountType = 1
	LinkedAccountType_LINKED_ACCOUNT_TYPE_MOBILE_MONEY_ACCOUNT LinkedAccountType = 2
)

// Enum value maps for LinkedAccountType.
var (
	LinkedAccountType_name = map[int32]string{
		0: "LINKED_ACCOUNT_TYPE_UNSPECIFIED",
		1: "LINKED_ACCOUNT_TYPE_BANK_ACCOUNT",
		2: "LINKED_ACCOUNT_TYPE_MOBILE_MONEY_ACCOUNT",
	}
	LinkedAccountType_value = map[string]int32{
		"LINKED_ACCOUNT_TYPE_UNSPECIFIED":          0,
		"LINKED_ACCOUNT_TYPE_BANK_ACCOUNT":         1,
		"LINKED_ACCOUNT_TYPE_MOBILE_MONEY_ACCOUNT": 2,
	}
)

func (x LinkedAccountType) Enum() *LinkedAccountType {
	p := new(LinkedAccountType)
	*p = x
	return p
}

func (x LinkedAccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkedAccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_linked_account_proto_enumTypes[0].Descriptor()
}

func (LinkedAccountType) Type() protoreflect.EnumType {
	return &file_kappa_common_models_linked_account_proto_enumTypes[0]
}

func (x LinkedAccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkedAccountType.Descriptor instead.
func (LinkedAccountType) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_linked_account_proto_rawDescGZIP(), []int{0}
}

type LinkedAccountLinkProvider int32

const (
	LinkedAccountLinkProvider_LINKED_ACCOUNT_LINK_PROVIDER_UNSPECIFIED LinkedAccountLinkProvider = 0
	LinkedAccountLinkProvider_LINKED_ACCOUNT_LINK_PROVIDER_KAPPA       LinkedAccountLinkProvider = 1
	LinkedAccountLinkProvider_LINKED_ACCOUNT_LINK_PROVIDER_PLAID       LinkedAccountLinkProvider = 2
)

// Enum value maps for LinkedAccountLinkProvider.
var (
	LinkedAccountLinkProvider_name = map[int32]string{
		0: "LINKED_ACCOUNT_LINK_PROVIDER_UNSPECIFIED",
		1: "LINKED_ACCOUNT_LINK_PROVIDER_KAPPA",
		2: "LINKED_ACCOUNT_LINK_PROVIDER_PLAID",
	}
	LinkedAccountLinkProvider_value = map[string]int32{
		"LINKED_ACCOUNT_LINK_PROVIDER_UNSPECIFIED": 0,
		"LINKED_ACCOUNT_LINK_PROVIDER_KAPPA":       1,
		"LINKED_ACCOUNT_LINK_PROVIDER_PLAID":       2,
	}
)

func (x LinkedAccountLinkProvider) Enum() *LinkedAccountLinkProvider {
	p := new(LinkedAccountLinkProvider)
	*p = x
	return p
}

func (x LinkedAccountLinkProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkedAccountLinkProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_linked_account_proto_enumTypes[1].Descriptor()
}

func (LinkedAccountLinkProvider) Type() protoreflect.EnumType {
	return &file_kappa_common_models_linked_account_proto_enumTypes[1]
}

func (x LinkedAccountLinkProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkedAccountLinkProvider.Descriptor instead.
func (LinkedAccountLinkProvider) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_linked_account_proto_rawDescGZIP(), []int{1}
}

type LinkedAccountStatus int32

const (
	LinkedAccountStatus_LINKED_ACCOUNT_STATUS_UNSPECIFIED               LinkedAccountStatus = 0
	LinkedAccountStatus_LINKED_ACCOUNT_STATUS_PENDING_AUTHORIZATION     LinkedAccountStatus = 1
	LinkedAccountStatus_LINKED_ACCOUNT_STATUS_AUTHORIZED                LinkedAccountStatus = 2
	LinkedAccountStatus_LINKED_ACCOUNT_STATUS_FAILED_TO_AUTHORIZE       LinkedAccountStatus = 3
	LinkedAccountStatus_LINKED_ACCOUNT_STATUS_ACTIVE_REQUIRES_RELINKING LinkedAccountStatus = 4
)

// Enum value maps for LinkedAccountStatus.
var (
	LinkedAccountStatus_name = map[int32]string{
		0: "LINKED_ACCOUNT_STATUS_UNSPECIFIED",
		1: "LINKED_ACCOUNT_STATUS_PENDING_AUTHORIZATION",
		2: "LINKED_ACCOUNT_STATUS_AUTHORIZED",
		3: "LINKED_ACCOUNT_STATUS_FAILED_TO_AUTHORIZE",
		4: "LINKED_ACCOUNT_STATUS_ACTIVE_REQUIRES_RELINKING",
	}
	LinkedAccountStatus_value = map[string]int32{
		"LINKED_ACCOUNT_STATUS_UNSPECIFIED":               0,
		"LINKED_ACCOUNT_STATUS_PENDING_AUTHORIZATION":     1,
		"LINKED_ACCOUNT_STATUS_AUTHORIZED":                2,
		"LINKED_ACCOUNT_STATUS_FAILED_TO_AUTHORIZE":       3,
		"LINKED_ACCOUNT_STATUS_ACTIVE_REQUIRES_RELINKING": 4,
	}
)

func (x LinkedAccountStatus) Enum() *LinkedAccountStatus {
	p := new(LinkedAccountStatus)
	*p = x
	return p
}

func (x LinkedAccountStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkedAccountStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_linked_account_proto_enumTypes[2].Descriptor()
}

func (LinkedAccountStatus) Type() protoreflect.EnumType {
	return &file_kappa_common_models_linked_account_proto_enumTypes[2]
}

func (x LinkedAccountStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkedAccountStatus.Descriptor instead.
func (LinkedAccountStatus) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_linked_account_proto_rawDescGZIP(), []int{2}
}

type LinkedAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type         LinkedAccountType         `protobuf:"varint,2,opt,name=type,proto3,enum=kappa.LinkedAccountType" json:"type,omitempty"`
	Entity       *Entity                   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	LinkProvider LinkedAccountLinkProvider `protobuf:"varint,4,opt,name=link_provider,json=linkProvider,proto3,enum=kappa.LinkedAccountLinkProvider" json:"link_provider,omitempty"`
	Status       LinkedAccountStatus       `protobuf:"varint,5,opt,name=status,proto3,enum=kappa.LinkedAccountStatus" json:"status,omitempty"`
	// Types that are assignable to LinkedAccount:
	//	*LinkedAccount_BankAccount
	//	*LinkedAccount_MobileMoneyAccount
	LinkedAccount isLinkedAccount_LinkedAccount `protobuf_oneof:"linked_account"`
}

func (x *LinkedAccount) Reset() {
	*x = LinkedAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_linked_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkedAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedAccount) ProtoMessage() {}

func (x *LinkedAccount) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_linked_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedAccount.ProtoReflect.Descriptor instead.
func (*LinkedAccount) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_linked_account_proto_rawDescGZIP(), []int{0}
}

func (x *LinkedAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LinkedAccount) GetType() LinkedAccountType {
	if x != nil {
		return x.Type
	}
	return LinkedAccountType_LINKED_ACCOUNT_TYPE_UNSPECIFIED
}

func (x *LinkedAccount) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *LinkedAccount) GetLinkProvider() LinkedAccountLinkProvider {
	if x != nil {
		return x.LinkProvider
	}
	return LinkedAccountLinkProvider_LINKED_ACCOUNT_LINK_PROVIDER_UNSPECIFIED
}

func (x *LinkedAccount) GetStatus() LinkedAccountStatus {
	if x != nil {
		return x.Status
	}
	return LinkedAccountStatus_LINKED_ACCOUNT_STATUS_UNSPECIFIED
}

func (m *LinkedAccount) GetLinkedAccount() isLinkedAccount_LinkedAccount {
	if m != nil {
		return m.LinkedAccount
	}
	return nil
}

func (x *LinkedAccount) GetBankAccount() *BankAccount {
	if x, ok := x.GetLinkedAccount().(*LinkedAccount_BankAccount); ok {
		return x.BankAccount
	}
	return nil
}

func (x *LinkedAccount) GetMobileMoneyAccount() *MobileMoneyAccount {
	if x, ok := x.GetLinkedAccount().(*LinkedAccount_MobileMoneyAccount); ok {
		return x.MobileMoneyAccount
	}
	return nil
}

type isLinkedAccount_LinkedAccount interface {
	isLinkedAccount_LinkedAccount()
}

type LinkedAccount_BankAccount struct {
	BankAccount *BankAccount `protobuf:"bytes,6,opt,name=bank_account,json=bankAccount,proto3,oneof"`
}

type LinkedAccount_MobileMoneyAccount struct {
	MobileMoneyAccount *MobileMoneyAccount `protobuf:"bytes,7,opt,name=mobile_money_account,json=mobileMoneyAccount,proto3,oneof"`
}

func (*LinkedAccount_BankAccount) isLinkedAccount_LinkedAccount() {}

func (*LinkedAccount_MobileMoneyAccount) isLinkedAccount_LinkedAccount() {}

var File_kappa_common_models_linked_account_proto protoreflect.FileDescriptor

var file_kappa_common_models_linked_account_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6b, 0x61, 0x70, 0x70,
	0x61, 0x1a, 0x26, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x03,
	0x0a, 0x0d, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0d, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x6c,
	0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x37, 0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x42, 0x61,
	0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x14, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x48, 0x00, 0x52, 0x12, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x8c, 0x01, 0x0a, 0x11, 0x4c, 0x69,
	0x6e, 0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x23, 0x0a, 0x1f, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x4e, 0x4b,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x2c, 0x0a, 0x28, 0x4c, 0x49,
	0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x2a, 0x99, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x28, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x52,
	0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x52, 0x4f, 0x56,
	0x49, 0x44, 0x45, 0x52, 0x5f, 0x4b, 0x41, 0x50, 0x50, 0x41, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22,
	0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c,
	0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41,
	0x49, 0x44, 0x10, 0x02, 0x2a, 0xf7, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x21,
	0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x2f, 0x0a, 0x2b, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x55,
	0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x02, 0x12, 0x2d, 0x0a, 0x29, 0x4c, 0x49,
	0x4e, 0x4b, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x41, 0x55,
	0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x10, 0x03, 0x12, 0x33, 0x0a, 0x2f, 0x4c, 0x49, 0x4e,
	0x4b, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52,
	0x45, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x4e, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x42, 0x40,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x70,
	0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69,
	0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61, 0x70,
	0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_common_models_linked_account_proto_rawDescOnce sync.Once
	file_kappa_common_models_linked_account_proto_rawDescData = file_kappa_common_models_linked_account_proto_rawDesc
)

func file_kappa_common_models_linked_account_proto_rawDescGZIP() []byte {
	file_kappa_common_models_linked_account_proto_rawDescOnce.Do(func() {
		file_kappa_common_models_linked_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_common_models_linked_account_proto_rawDescData)
	})
	return file_kappa_common_models_linked_account_proto_rawDescData
}

var file_kappa_common_models_linked_account_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_kappa_common_models_linked_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kappa_common_models_linked_account_proto_goTypes = []interface{}{
	(LinkedAccountType)(0),         // 0: kappa.LinkedAccountType
	(LinkedAccountLinkProvider)(0), // 1: kappa.LinkedAccountLinkProvider
	(LinkedAccountStatus)(0),       // 2: kappa.LinkedAccountStatus
	(*LinkedAccount)(nil),          // 3: kappa.LinkedAccount
	(*Entity)(nil),                 // 4: kappa.Entity
	(*BankAccount)(nil),            // 5: kappa.BankAccount
	(*MobileMoneyAccount)(nil),     // 6: kappa.MobileMoneyAccount
}
var file_kappa_common_models_linked_account_proto_depIdxs = []int32{
	0, // 0: kappa.LinkedAccount.type:type_name -> kappa.LinkedAccountType
	4, // 1: kappa.LinkedAccount.entity:type_name -> kappa.Entity
	1, // 2: kappa.LinkedAccount.link_provider:type_name -> kappa.LinkedAccountLinkProvider
	2, // 3: kappa.LinkedAccount.status:type_name -> kappa.LinkedAccountStatus
	5, // 4: kappa.LinkedAccount.bank_account:type_name -> kappa.BankAccount
	6, // 5: kappa.LinkedAccount.mobile_money_account:type_name -> kappa.MobileMoneyAccount
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_kappa_common_models_linked_account_proto_init() }
func file_kappa_common_models_linked_account_proto_init() {
	if File_kappa_common_models_linked_account_proto != nil {
		return
	}
	file_kappa_common_models_bank_account_proto_init()
	file_kappa_common_models_entities_proto_init()
	file_kappa_common_models_mobile_money_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kappa_common_models_linked_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkedAccount); i {
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
	file_kappa_common_models_linked_account_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LinkedAccount_BankAccount)(nil),
		(*LinkedAccount_MobileMoneyAccount)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kappa_common_models_linked_account_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kappa_common_models_linked_account_proto_goTypes,
		DependencyIndexes: file_kappa_common_models_linked_account_proto_depIdxs,
		EnumInfos:         file_kappa_common_models_linked_account_proto_enumTypes,
		MessageInfos:      file_kappa_common_models_linked_account_proto_msgTypes,
	}.Build()
	File_kappa_common_models_linked_account_proto = out.File
	file_kappa_common_models_linked_account_proto_rawDesc = nil
	file_kappa_common_models_linked_account_proto_goTypes = nil
	file_kappa_common_models_linked_account_proto_depIdxs = nil
}

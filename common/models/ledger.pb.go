// ------------------------------------------------------------------------------
// <copyright company="Kappa Pay Inc.">
// Copyright (C) Kappa Pay Inc.  All rights reserved.
// Unauthorised copying of this file, via any medium, is strictly prohibited.
// Proprietary and confidential
// </copyright>
// <author>Chris Nyaga</author>
// ------------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: kappa/common/models/ledger.proto

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

type LedgerAccountType int32

const (
	LedgerAccountType_LEDGER_ACCOUNT_TYPE_UNSPECIFIED   LedgerAccountType = 0
	LedgerAccountType_LEDGER_ACCOUNT_TYPE_DEBIT_NORMAL  LedgerAccountType = 1
	LedgerAccountType_LEDGER_ACCOUNT_TYPE_CREDIT_NORMAL LedgerAccountType = 2
)

// Enum value maps for LedgerAccountType.
var (
	LedgerAccountType_name = map[int32]string{
		0: "LEDGER_ACCOUNT_TYPE_UNSPECIFIED",
		1: "LEDGER_ACCOUNT_TYPE_DEBIT_NORMAL",
		2: "LEDGER_ACCOUNT_TYPE_CREDIT_NORMAL",
	}
	LedgerAccountType_value = map[string]int32{
		"LEDGER_ACCOUNT_TYPE_UNSPECIFIED":   0,
		"LEDGER_ACCOUNT_TYPE_DEBIT_NORMAL":  1,
		"LEDGER_ACCOUNT_TYPE_CREDIT_NORMAL": 2,
	}
)

func (x LedgerAccountType) Enum() *LedgerAccountType {
	p := new(LedgerAccountType)
	*p = x
	return p
}

func (x LedgerAccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LedgerAccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_ledger_proto_enumTypes[0].Descriptor()
}

func (LedgerAccountType) Type() protoreflect.EnumType {
	return &file_kappa_common_models_ledger_proto_enumTypes[0]
}

func (x LedgerAccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LedgerAccountType.Descriptor instead.
func (LedgerAccountType) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_ledger_proto_rawDescGZIP(), []int{0}
}

type LedgerEntryDirection int32

const (
	LedgerEntryDirection_LEDGER_ENTRY_DIRECTION_UNSPECIFIED LedgerEntryDirection = 0
	LedgerEntryDirection_LEDGER_ENTRY_DIRECTION_DEBIT       LedgerEntryDirection = 1
	LedgerEntryDirection_LEDGER_ENTRY_DIRECTION_CREDIT      LedgerEntryDirection = 2
)

// Enum value maps for LedgerEntryDirection.
var (
	LedgerEntryDirection_name = map[int32]string{
		0: "LEDGER_ENTRY_DIRECTION_UNSPECIFIED",
		1: "LEDGER_ENTRY_DIRECTION_DEBIT",
		2: "LEDGER_ENTRY_DIRECTION_CREDIT",
	}
	LedgerEntryDirection_value = map[string]int32{
		"LEDGER_ENTRY_DIRECTION_UNSPECIFIED": 0,
		"LEDGER_ENTRY_DIRECTION_DEBIT":       1,
		"LEDGER_ENTRY_DIRECTION_CREDIT":      2,
	}
)

func (x LedgerEntryDirection) Enum() *LedgerEntryDirection {
	p := new(LedgerEntryDirection)
	*p = x
	return p
}

func (x LedgerEntryDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LedgerEntryDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_ledger_proto_enumTypes[1].Descriptor()
}

func (LedgerEntryDirection) Type() protoreflect.EnumType {
	return &file_kappa_common_models_ledger_proto_enumTypes[1]
}

func (x LedgerEntryDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LedgerEntryDirection.Descriptor instead.
func (LedgerEntryDirection) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_ledger_proto_rawDescGZIP(), []int{1}
}

type Ledger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Ledger) Reset() {
	*x = Ledger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_ledger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ledger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ledger) ProtoMessage() {}

func (x *Ledger) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_ledger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ledger.ProtoReflect.Descriptor instead.
func (*Ledger) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_ledger_proto_rawDescGZIP(), []int{0}
}

func (x *Ledger) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ledger) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LedgerAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type     LedgerAccountType `protobuf:"varint,2,opt,name=type,proto3,enum=kappa.LedgerAccountType" json:"type,omitempty"`
	EntityId string            `protobuf:"bytes,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"` // This is used by the path field as well as to identify all the accounts owned by a particular entity
	Path     []string          `protobuf:"bytes,4,rep,name=path,proto3" json:"path,omitempty"`
	Name     string            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Balance  *Amount           `protobuf:"bytes,6,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *LedgerAccount) Reset() {
	*x = LedgerAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_ledger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedgerAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedgerAccount) ProtoMessage() {}

func (x *LedgerAccount) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_ledger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedgerAccount.ProtoReflect.Descriptor instead.
func (*LedgerAccount) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_ledger_proto_rawDescGZIP(), []int{1}
}

func (x *LedgerAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LedgerAccount) GetType() LedgerAccountType {
	if x != nil {
		return x.Type
	}
	return LedgerAccountType_LEDGER_ACCOUNT_TYPE_UNSPECIFIED
}

func (x *LedgerAccount) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *LedgerAccount) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *LedgerAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LedgerAccount) GetBalance() *Amount {
	if x != nil {
		return x.Balance
	}
	return nil
}

type LedgerEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LedgerAccountId string               `protobuf:"bytes,1,opt,name=ledger_account_id,json=ledgerAccountId,proto3" json:"ledger_account_id,omitempty"`
	Direction       LedgerEntryDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=kappa.LedgerEntryDirection" json:"direction,omitempty"`
	Amount          *Amount              `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *LedgerEntry) Reset() {
	*x = LedgerEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_ledger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedgerEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedgerEntry) ProtoMessage() {}

func (x *LedgerEntry) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_ledger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedgerEntry.ProtoReflect.Descriptor instead.
func (*LedgerEntry) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_ledger_proto_rawDescGZIP(), []int{2}
}

func (x *LedgerEntry) GetLedgerAccountId() string {
	if x != nil {
		return x.LedgerAccountId
	}
	return ""
}

func (x *LedgerEntry) GetDirection() LedgerEntryDirection {
	if x != nil {
		return x.Direction
	}
	return LedgerEntryDirection_LEDGER_ENTRY_DIRECTION_UNSPECIFIED
}

func (x *LedgerEntry) GetAmount() *Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

type LedgerTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is a required uuid field that must be defined during posting of a Ledger Transaction. It is used as an idempotency key by the ledger.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// When posting transactions to the journal, you can (optionally) supply zero or more relatives. Relatives are tokens for other transactions which are "related" to the transaction being added.
	// Relations are modelled as a graph and resolved dynamically in the following ways:
	//
	// With:
	//
	//    transaction1/token1 has no relatives
	//    transaction2/token2 specifies token1 as a relative
	//  Then, getting transactions related to:
	//
	//    token1 will yield transaction1, transaction2.
	//    token2 will yield transaction1, transaction2.
	//
	// With:
	//
	//    transaction1/token1 specifies token2 as a relative
	//    transaction2/token2 specifies token1 as a relative
	//    transaction3/token3 specifies token1 as a relative.
	//  Then, getting transactions related to:
	//
	//    token1 will yield transaction1, transaction2, transaction3.
	//    token2 will yield transaction1, transaction2, transaction3.
	//    token3 will yield transaction1, transaction2, transaction3.
	//
	// With:
	//
	//    transaction1/token1 specifies token10 as a relative
	//    transaction2/token2 specifies token10 as a relative.
	//  Then, getting transactions related to:
	//
	//    token1 will yield transaction1, transaction2.
	//    token2 will yield transaction1, transaction2.
	//    token10 will yield transaction1, transaction2.
	Relations []string `protobuf:"bytes,2,rep,name=relations,proto3" json:"relations,omitempty"`
	// Balanced set of ledger entries
	LedgerEntries []*LedgerEntry `protobuf:"bytes,3,rep,name=ledger_entries,json=ledgerEntries,proto3" json:"ledger_entries,omitempty"`
	// Optional description note for the transaction
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *LedgerTransaction) Reset() {
	*x = LedgerTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_ledger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedgerTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedgerTransaction) ProtoMessage() {}

func (x *LedgerTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_ledger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedgerTransaction.ProtoReflect.Descriptor instead.
func (*LedgerTransaction) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_ledger_proto_rawDescGZIP(), []int{3}
}

func (x *LedgerTransaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LedgerTransaction) GetRelations() []string {
	if x != nil {
		return x.Relations
	}
	return nil
}

func (x *LedgerTransaction) GetLedgerEntries() []*LedgerEntry {
	if x != nil {
		return x.LedgerEntries
	}
	return nil
}

func (x *LedgerTransaction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_kappa_common_models_ledger_proto protoreflect.FileDescriptor

var file_kappa_common_models_ledger_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x1a, 0x20, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x0d, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6b, 0x61, 0x70, 0x70,
	0x61, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x4c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x11, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x85, 0x01, 0x0a, 0x11, 0x4c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f,
	0x4c, 0x45, 0x44, 0x47, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x24, 0x0a, 0x20, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x42, 0x49, 0x54, 0x5f, 0x4e,
	0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x4c, 0x45, 0x44, 0x47, 0x45,
	0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x52, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x2a, 0x83,
	0x01, 0x0a, 0x14, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x22, 0x4c, 0x45, 0x44, 0x47, 0x45,
	0x52, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x20, 0x0a, 0x1c, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f,
	0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x42, 0x49, 0x54, 0x10,
	0x01, 0x12, 0x21, 0x0a, 0x1d, 0x4c, 0x45, 0x44, 0x47, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x52,
	0x59, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x44,
	0x49, 0x54, 0x10, 0x02, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_common_models_ledger_proto_rawDescOnce sync.Once
	file_kappa_common_models_ledger_proto_rawDescData = file_kappa_common_models_ledger_proto_rawDesc
)

func file_kappa_common_models_ledger_proto_rawDescGZIP() []byte {
	file_kappa_common_models_ledger_proto_rawDescOnce.Do(func() {
		file_kappa_common_models_ledger_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_common_models_ledger_proto_rawDescData)
	})
	return file_kappa_common_models_ledger_proto_rawDescData
}

var file_kappa_common_models_ledger_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_kappa_common_models_ledger_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kappa_common_models_ledger_proto_goTypes = []interface{}{
	(LedgerAccountType)(0),    // 0: kappa.LedgerAccountType
	(LedgerEntryDirection)(0), // 1: kappa.LedgerEntryDirection
	(*Ledger)(nil),            // 2: kappa.Ledger
	(*LedgerAccount)(nil),     // 3: kappa.LedgerAccount
	(*LedgerEntry)(nil),       // 4: kappa.LedgerEntry
	(*LedgerTransaction)(nil), // 5: kappa.LedgerTransaction
	(*Amount)(nil),            // 6: kappa.Amount
}
var file_kappa_common_models_ledger_proto_depIdxs = []int32{
	0, // 0: kappa.LedgerAccount.type:type_name -> kappa.LedgerAccountType
	6, // 1: kappa.LedgerAccount.balance:type_name -> kappa.Amount
	1, // 2: kappa.LedgerEntry.direction:type_name -> kappa.LedgerEntryDirection
	6, // 3: kappa.LedgerEntry.amount:type_name -> kappa.Amount
	4, // 4: kappa.LedgerTransaction.ledger_entries:type_name -> kappa.LedgerEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_kappa_common_models_ledger_proto_init() }
func file_kappa_common_models_ledger_proto_init() {
	if File_kappa_common_models_ledger_proto != nil {
		return
	}
	file_kappa_common_models_amount_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kappa_common_models_ledger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ledger); i {
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
		file_kappa_common_models_ledger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedgerAccount); i {
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
		file_kappa_common_models_ledger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedgerEntry); i {
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
		file_kappa_common_models_ledger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedgerTransaction); i {
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
			RawDescriptor: file_kappa_common_models_ledger_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kappa_common_models_ledger_proto_goTypes,
		DependencyIndexes: file_kappa_common_models_ledger_proto_depIdxs,
		EnumInfos:         file_kappa_common_models_ledger_proto_enumTypes,
		MessageInfos:      file_kappa_common_models_ledger_proto_msgTypes,
	}.Build()
	File_kappa_common_models_ledger_proto = out.File
	file_kappa_common_models_ledger_proto_rawDesc = nil
	file_kappa_common_models_ledger_proto_goTypes = nil
	file_kappa_common_models_ledger_proto_depIdxs = nil
}

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
// source: kappa/common/models/quote.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionQuoteStatus int32

const (
	TransactionQuoteStatus_TRANSACTION_QUOTE_STATUS_UNSPECIFIED TransactionQuoteStatus = 0
	TransactionQuoteStatus_TRANSACTION_QUOTE_STATUS_REQUESTED   TransactionQuoteStatus = 1
	TransactionQuoteStatus_TRANSACTION_QUOTE_STATUS_ACCEPTED    TransactionQuoteStatus = 2
	TransactionQuoteStatus_TRANSACTION_QUOTE_STATUS_REJECTED    TransactionQuoteStatus = 3
	TransactionQuoteStatus_TRANSACTION_QUOTE_STATUS_EXPIRED     TransactionQuoteStatus = 4
)

// Enum value maps for TransactionQuoteStatus.
var (
	TransactionQuoteStatus_name = map[int32]string{
		0: "TRANSACTION_QUOTE_STATUS_UNSPECIFIED",
		1: "TRANSACTION_QUOTE_STATUS_REQUESTED",
		2: "TRANSACTION_QUOTE_STATUS_ACCEPTED",
		3: "TRANSACTION_QUOTE_STATUS_REJECTED",
		4: "TRANSACTION_QUOTE_STATUS_EXPIRED",
	}
	TransactionQuoteStatus_value = map[string]int32{
		"TRANSACTION_QUOTE_STATUS_UNSPECIFIED": 0,
		"TRANSACTION_QUOTE_STATUS_REQUESTED":   1,
		"TRANSACTION_QUOTE_STATUS_ACCEPTED":    2,
		"TRANSACTION_QUOTE_STATUS_REJECTED":    3,
		"TRANSACTION_QUOTE_STATUS_EXPIRED":     4,
	}
)

func (x TransactionQuoteStatus) Enum() *TransactionQuoteStatus {
	p := new(TransactionQuoteStatus)
	*p = x
	return p
}

func (x TransactionQuoteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionQuoteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_quote_proto_enumTypes[0].Descriptor()
}

func (TransactionQuoteStatus) Type() protoreflect.EnumType {
	return &file_kappa_common_models_quote_proto_enumTypes[0]
}

func (x TransactionQuoteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionQuoteStatus.Descriptor instead.
func (TransactionQuoteStatus) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_quote_proto_rawDescGZIP(), []int{0}
}

type ExecutionQuoteStatus int32

const (
	ExecutionQuoteStatus_EXECUTION_QUOTE_STATUS_UNSPECIFIED ExecutionQuoteStatus = 0
	ExecutionQuoteStatus_EXECUTION_QUOTE_STATUS_REQUESTED   ExecutionQuoteStatus = 1
	ExecutionQuoteStatus_EXECUTION_QUOTE_STATUS_ACCEPTED    ExecutionQuoteStatus = 2
	ExecutionQuoteStatus_EXECUTION_QUOTE_STATUS_REJECTED    ExecutionQuoteStatus = 3
	ExecutionQuoteStatus_EXECUTION_QUOTE_STATUS_EXPIRED     ExecutionQuoteStatus = 4
)

// Enum value maps for ExecutionQuoteStatus.
var (
	ExecutionQuoteStatus_name = map[int32]string{
		0: "EXECUTION_QUOTE_STATUS_UNSPECIFIED",
		1: "EXECUTION_QUOTE_STATUS_REQUESTED",
		2: "EXECUTION_QUOTE_STATUS_ACCEPTED",
		3: "EXECUTION_QUOTE_STATUS_REJECTED",
		4: "EXECUTION_QUOTE_STATUS_EXPIRED",
	}
	ExecutionQuoteStatus_value = map[string]int32{
		"EXECUTION_QUOTE_STATUS_UNSPECIFIED": 0,
		"EXECUTION_QUOTE_STATUS_REQUESTED":   1,
		"EXECUTION_QUOTE_STATUS_ACCEPTED":    2,
		"EXECUTION_QUOTE_STATUS_REJECTED":    3,
		"EXECUTION_QUOTE_STATUS_EXPIRED":     4,
	}
)

func (x ExecutionQuoteStatus) Enum() *ExecutionQuoteStatus {
	p := new(ExecutionQuoteStatus)
	*p = x
	return p
}

func (x ExecutionQuoteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecutionQuoteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_quote_proto_enumTypes[1].Descriptor()
}

func (ExecutionQuoteStatus) Type() protoreflect.EnumType {
	return &file_kappa_common_models_quote_proto_enumTypes[1]
}

func (x ExecutionQuoteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecutionQuoteStatus.Descriptor instead.
func (ExecutionQuoteStatus) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_quote_proto_rawDescGZIP(), []int{1}
}

type TransactionQuote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               TransactionQuoteStatus  `protobuf:"varint,2,opt,name=status,proto3,enum=kappa.TransactionQuoteStatus" json:"status,omitempty"`
	TransactionType      TransactionType         `protobuf:"varint,3,opt,name=transaction_type,json=transactionType,proto3,enum=kappa.TransactionType" json:"transaction_type,omitempty"`
	Source               *Source                 `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Destination          *Destination            `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
	SourceAmount         *Amount                 `protobuf:"bytes,6,opt,name=source_amount,json=sourceAmount,proto3" json:"source_amount,omitempty"`
	DestinationAmount    *Amount                 `protobuf:"bytes,7,opt,name=destination_amount,json=destinationAmount,proto3" json:"destination_amount,omitempty"`
	Fees                 *TransactionFees        `protobuf:"bytes,8,opt,name=fees,proto3" json:"fees,omitempty"`
	ExchangeRate         *ExchangeRate           `protobuf:"bytes,9,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"`
	EstimatedArrivalTime *timestamppb.Timestamp  `protobuf:"bytes,10,opt,name=estimated_arrival_time,json=estimatedArrivalTime,proto3" json:"estimated_arrival_time,omitempty"`
	ExpiresAt            *timestamppb.Timestamp  `protobuf:"bytes,11,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CreatedAt            *timestamppb.Timestamp  `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp  `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamppb.Timestamp  `protobuf:"bytes,14,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	ConvertedAmount      *Amount                 `protobuf:"bytes,15,opt,name=converted_amount,json=convertedAmount,proto3" json:"converted_amount,omitempty"`
	Taxes                []*TransactionTaxAmount `protobuf:"bytes,16,rep,name=taxes,proto3" json:"taxes,omitempty"`
}

func (x *TransactionQuote) Reset() {
	*x = TransactionQuote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionQuote) ProtoMessage() {}

func (x *TransactionQuote) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionQuote.ProtoReflect.Descriptor instead.
func (*TransactionQuote) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_quote_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionQuote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TransactionQuote) GetStatus() TransactionQuoteStatus {
	if x != nil {
		return x.Status
	}
	return TransactionQuoteStatus_TRANSACTION_QUOTE_STATUS_UNSPECIFIED
}

func (x *TransactionQuote) GetTransactionType() TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return TransactionType_TRANSACTION_TYPE_UNSPECIFIED
}

func (x *TransactionQuote) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *TransactionQuote) GetDestination() *Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *TransactionQuote) GetSourceAmount() *Amount {
	if x != nil {
		return x.SourceAmount
	}
	return nil
}

func (x *TransactionQuote) GetDestinationAmount() *Amount {
	if x != nil {
		return x.DestinationAmount
	}
	return nil
}

func (x *TransactionQuote) GetFees() *TransactionFees {
	if x != nil {
		return x.Fees
	}
	return nil
}

func (x *TransactionQuote) GetExchangeRate() *ExchangeRate {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

func (x *TransactionQuote) GetEstimatedArrivalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EstimatedArrivalTime
	}
	return nil
}

func (x *TransactionQuote) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *TransactionQuote) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TransactionQuote) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *TransactionQuote) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *TransactionQuote) GetConvertedAmount() *Amount {
	if x != nil {
		return x.ConvertedAmount
	}
	return nil
}

func (x *TransactionQuote) GetTaxes() []*TransactionTaxAmount {
	if x != nil {
		return x.Taxes
	}
	return nil
}

type ExecutionQuote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RequesterId          string                 `protobuf:"bytes,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`       // ID of the entity requesting the quote
	ProviderId           string                 `protobuf:"bytes,3,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`          // ID of the entity providing the quote
	TransactionId        string                 `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"` // ID of the transaction to be serviced by the partner
	Status               ExecutionQuoteStatus   `protobuf:"varint,5,opt,name=status,proto3,enum=kappa.ExecutionQuoteStatus" json:"status,omitempty"`
	Source               *Source                `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	Destination          *Destination           `protobuf:"bytes,7,opt,name=destination,proto3" json:"destination,omitempty"`
	SourceAmount         *Amount                `protobuf:"bytes,8,opt,name=source_amount,json=sourceAmount,proto3" json:"source_amount,omitempty"`
	DestinationAmount    *Amount                `protobuf:"bytes,9,opt,name=destination_amount,json=destinationAmount,proto3" json:"destination_amount,omitempty"`
	ConvertedAmount      *Amount                `protobuf:"bytes,10,opt,name=converted_amount,json=convertedAmount,proto3" json:"converted_amount,omitempty"`
	Fees                 *ExecutionFees         `protobuf:"bytes,11,opt,name=fees,proto3" json:"fees,omitempty"`
	ExchangeRate         *ExchangeRate          `protobuf:"bytes,12,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"`
	EstimatedArrivalTime *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=estimated_arrival_time,json=estimatedArrivalTime,proto3" json:"estimated_arrival_time,omitempty"`
	ExpiresAt            *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamppb.Timestamp `protobuf:"bytes,17,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *ExecutionQuote) Reset() {
	*x = ExecutionQuote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionQuote) ProtoMessage() {}

func (x *ExecutionQuote) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionQuote.ProtoReflect.Descriptor instead.
func (*ExecutionQuote) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_quote_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionQuote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExecutionQuote) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *ExecutionQuote) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *ExecutionQuote) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *ExecutionQuote) GetStatus() ExecutionQuoteStatus {
	if x != nil {
		return x.Status
	}
	return ExecutionQuoteStatus_EXECUTION_QUOTE_STATUS_UNSPECIFIED
}

func (x *ExecutionQuote) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ExecutionQuote) GetDestination() *Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *ExecutionQuote) GetSourceAmount() *Amount {
	if x != nil {
		return x.SourceAmount
	}
	return nil
}

func (x *ExecutionQuote) GetDestinationAmount() *Amount {
	if x != nil {
		return x.DestinationAmount
	}
	return nil
}

func (x *ExecutionQuote) GetConvertedAmount() *Amount {
	if x != nil {
		return x.ConvertedAmount
	}
	return nil
}

func (x *ExecutionQuote) GetFees() *ExecutionFees {
	if x != nil {
		return x.Fees
	}
	return nil
}

func (x *ExecutionQuote) GetExchangeRate() *ExchangeRate {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

func (x *ExecutionQuote) GetEstimatedArrivalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EstimatedArrivalTime
	}
	return nil
}

func (x *ExecutionQuote) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *ExecutionQuote) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ExecutionQuote) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ExecutionQuote) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_kappa_common_models_quote_proto protoreflect.FileDescriptor

var file_kappa_common_models_quote_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x1a, 0x20, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6b, 0x61, 0x70, 0x70,
	0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x78, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x66, 0x65, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x06, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x41, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70,
	0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x73, 0x52, 0x04, 0x66, 0x65, 0x65, 0x73,
	0x12, 0x38, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x16, 0x65, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x74, 0x61, 0x78, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05,
	0x74, 0x61, 0x78, 0x65, 0x73, 0x22, 0xeb, 0x06, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0c, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x12, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x65, 0x65, 0x73, 0x52, 0x04, 0x66, 0x65, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0d, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x16, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x14, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x41, 0x72, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x2a, 0xde, 0x01, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28,
	0x0a, 0x24, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x55,
	0x4f, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x25, 0x0a, 0x21, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x51, 0x55, 0x4f, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x43,
	0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x24,
	0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x55,
	0x4f, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52,
	0x45, 0x44, 0x10, 0x04, 0x2a, 0xd2, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a,
	0x22, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x23, 0x0a, 0x1f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x55,
	0x4f, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43,
	0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x04, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_kappa_common_models_quote_proto_rawDescOnce sync.Once
	file_kappa_common_models_quote_proto_rawDescData = file_kappa_common_models_quote_proto_rawDesc
)

func file_kappa_common_models_quote_proto_rawDescGZIP() []byte {
	file_kappa_common_models_quote_proto_rawDescOnce.Do(func() {
		file_kappa_common_models_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_common_models_quote_proto_rawDescData)
	})
	return file_kappa_common_models_quote_proto_rawDescData
}

var file_kappa_common_models_quote_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_kappa_common_models_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kappa_common_models_quote_proto_goTypes = []interface{}{
	(TransactionQuoteStatus)(0),   // 0: kappa.TransactionQuoteStatus
	(ExecutionQuoteStatus)(0),     // 1: kappa.ExecutionQuoteStatus
	(*TransactionQuote)(nil),      // 2: kappa.TransactionQuote
	(*ExecutionQuote)(nil),        // 3: kappa.ExecutionQuote
	(TransactionType)(0),          // 4: kappa.TransactionType
	(*Source)(nil),                // 5: kappa.Source
	(*Destination)(nil),           // 6: kappa.Destination
	(*Amount)(nil),                // 7: kappa.Amount
	(*TransactionFees)(nil),       // 8: kappa.TransactionFees
	(*ExchangeRate)(nil),          // 9: kappa.ExchangeRate
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*TransactionTaxAmount)(nil),  // 11: kappa.TransactionTaxAmount
	(*ExecutionFees)(nil),         // 12: kappa.ExecutionFees
}
var file_kappa_common_models_quote_proto_depIdxs = []int32{
	0,  // 0: kappa.TransactionQuote.status:type_name -> kappa.TransactionQuoteStatus
	4,  // 1: kappa.TransactionQuote.transaction_type:type_name -> kappa.TransactionType
	5,  // 2: kappa.TransactionQuote.source:type_name -> kappa.Source
	6,  // 3: kappa.TransactionQuote.destination:type_name -> kappa.Destination
	7,  // 4: kappa.TransactionQuote.source_amount:type_name -> kappa.Amount
	7,  // 5: kappa.TransactionQuote.destination_amount:type_name -> kappa.Amount
	8,  // 6: kappa.TransactionQuote.fees:type_name -> kappa.TransactionFees
	9,  // 7: kappa.TransactionQuote.exchange_rate:type_name -> kappa.ExchangeRate
	10, // 8: kappa.TransactionQuote.estimated_arrival_time:type_name -> google.protobuf.Timestamp
	10, // 9: kappa.TransactionQuote.expires_at:type_name -> google.protobuf.Timestamp
	10, // 10: kappa.TransactionQuote.created_at:type_name -> google.protobuf.Timestamp
	10, // 11: kappa.TransactionQuote.updated_at:type_name -> google.protobuf.Timestamp
	10, // 12: kappa.TransactionQuote.deleted_at:type_name -> google.protobuf.Timestamp
	7,  // 13: kappa.TransactionQuote.converted_amount:type_name -> kappa.Amount
	11, // 14: kappa.TransactionQuote.taxes:type_name -> kappa.TransactionTaxAmount
	1,  // 15: kappa.ExecutionQuote.status:type_name -> kappa.ExecutionQuoteStatus
	5,  // 16: kappa.ExecutionQuote.source:type_name -> kappa.Source
	6,  // 17: kappa.ExecutionQuote.destination:type_name -> kappa.Destination
	7,  // 18: kappa.ExecutionQuote.source_amount:type_name -> kappa.Amount
	7,  // 19: kappa.ExecutionQuote.destination_amount:type_name -> kappa.Amount
	7,  // 20: kappa.ExecutionQuote.converted_amount:type_name -> kappa.Amount
	12, // 21: kappa.ExecutionQuote.fees:type_name -> kappa.ExecutionFees
	9,  // 22: kappa.ExecutionQuote.exchange_rate:type_name -> kappa.ExchangeRate
	10, // 23: kappa.ExecutionQuote.estimated_arrival_time:type_name -> google.protobuf.Timestamp
	10, // 24: kappa.ExecutionQuote.expires_at:type_name -> google.protobuf.Timestamp
	10, // 25: kappa.ExecutionQuote.created_at:type_name -> google.protobuf.Timestamp
	10, // 26: kappa.ExecutionQuote.updated_at:type_name -> google.protobuf.Timestamp
	10, // 27: kappa.ExecutionQuote.deleted_at:type_name -> google.protobuf.Timestamp
	28, // [28:28] is the sub-list for method output_type
	28, // [28:28] is the sub-list for method input_type
	28, // [28:28] is the sub-list for extension type_name
	28, // [28:28] is the sub-list for extension extendee
	0,  // [0:28] is the sub-list for field type_name
}

func init() { file_kappa_common_models_quote_proto_init() }
func file_kappa_common_models_quote_proto_init() {
	if File_kappa_common_models_quote_proto != nil {
		return
	}
	file_kappa_common_models_amount_proto_init()
	file_kappa_common_models_exchange_rate_proto_init()
	file_kappa_common_models_source_proto_init()
	file_kappa_common_models_destination_proto_init()
	file_kappa_common_models_transaction_type_proto_init()
	file_kappa_common_models_transaction_taxes_proto_init()
	file_kappa_common_models_fees_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kappa_common_models_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionQuote); i {
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
		file_kappa_common_models_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionQuote); i {
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
			RawDescriptor: file_kappa_common_models_quote_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kappa_common_models_quote_proto_goTypes,
		DependencyIndexes: file_kappa_common_models_quote_proto_depIdxs,
		EnumInfos:         file_kappa_common_models_quote_proto_enumTypes,
		MessageInfos:      file_kappa_common_models_quote_proto_msgTypes,
	}.Build()
	File_kappa_common_models_quote_proto = out.File
	file_kappa_common_models_quote_proto_rawDesc = nil
	file_kappa_common_models_quote_proto_goTypes = nil
	file_kappa_common_models_quote_proto_depIdxs = nil
}
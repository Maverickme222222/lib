// ------------------------------------------------------------------------------
// <copyright company="Kappa Pay Inc.">
// Copyright (C) Kappa Pay Inc.  All rights reserved.
// Unauthorised copying of this file, via any medium, is strictly prohibited.
// Proprietary and confidential
// </copyright>
// <author>Shashank Agarwal</author>
// ------------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: kappa/services/exchange_rates/exchange_rates.proto

package exchange_rates

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

type GetExchangeRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasetSource models.DataSource `protobuf:"varint,1,opt,name=dataset_source,json=datasetSource,proto3,enum=kappa.DataSource" json:"dataset_source,omitempty"`
	DatasetName   string            `protobuf:"bytes,2,opt,name=dataset_name,json=datasetName,proto3" json:"dataset_name,omitempty"`
	Field         models.DataField  `protobuf:"varint,3,opt,name=field,proto3,enum=kappa.DataField" json:"field,omitempty"`
	CurrencyFrom  string            `protobuf:"bytes,4,opt,name=currency_from,json=currencyFrom,proto3" json:"currency_from,omitempty"`
	CurrencyTo    string            `protobuf:"bytes,5,opt,name=currency_to,json=currencyTo,proto3" json:"currency_to,omitempty"`
}

func (x *GetExchangeRateRequest) Reset() {
	*x = GetExchangeRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeRateRequest) ProtoMessage() {}

func (x *GetExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*GetExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_exchange_rates_exchange_rates_proto_rawDescGZIP(), []int{0}
}

func (x *GetExchangeRateRequest) GetDatasetSource() models.DataSource {
	if x != nil {
		return x.DatasetSource
	}
	return models.DataSource(0)
}

func (x *GetExchangeRateRequest) GetDatasetName() string {
	if x != nil {
		return x.DatasetName
	}
	return ""
}

func (x *GetExchangeRateRequest) GetField() models.DataField {
	if x != nil {
		return x.Field
	}
	return models.DataField(0)
}

func (x *GetExchangeRateRequest) GetCurrencyFrom() string {
	if x != nil {
		return x.CurrencyFrom
	}
	return ""
}

func (x *GetExchangeRateRequest) GetCurrencyTo() string {
	if x != nil {
		return x.CurrencyTo
	}
	return ""
}

type GetExchangeRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeRate *models.ExchangeRateDataPoint `protobuf:"bytes,1,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"`
}

func (x *GetExchangeRateResponse) Reset() {
	*x = GetExchangeRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeRateResponse) ProtoMessage() {}

func (x *GetExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*GetExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_exchange_rates_exchange_rates_proto_rawDescGZIP(), []int{1}
}

func (x *GetExchangeRateResponse) GetExchangeRate() *models.ExchangeRateDataPoint {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

type GetDataSetBySourceAndNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source models.DataSource `protobuf:"varint,1,opt,name=source,proto3,enum=kappa.DataSource" json:"source,omitempty"`
	Name   string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDataSetBySourceAndNameRequest) Reset() {
	*x = GetDataSetBySourceAndNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataSetBySourceAndNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataSetBySourceAndNameRequest) ProtoMessage() {}

func (x *GetDataSetBySourceAndNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataSetBySourceAndNameRequest.ProtoReflect.Descriptor instead.
func (*GetDataSetBySourceAndNameRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_exchange_rates_exchange_rates_proto_rawDescGZIP(), []int{2}
}

func (x *GetDataSetBySourceAndNameRequest) GetSource() models.DataSource {
	if x != nil {
		return x.Source
	}
	return models.DataSource(0)
}

func (x *GetDataSetBySourceAndNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDataSetBySourceAndNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataset *models.DataSet `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
}

func (x *GetDataSetBySourceAndNameResponse) Reset() {
	*x = GetDataSetBySourceAndNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataSetBySourceAndNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataSetBySourceAndNameResponse) ProtoMessage() {}

func (x *GetDataSetBySourceAndNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataSetBySourceAndNameResponse.ProtoReflect.Descriptor instead.
func (*GetDataSetBySourceAndNameResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_exchange_rates_exchange_rates_proto_rawDescGZIP(), []int{3}
}

func (x *GetDataSetBySourceAndNameResponse) GetDataset() *models.DataSet {
	if x != nil {
		return x.Dataset
	}
	return nil
}

type GetBlendedExchangeRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyFrom string           `protobuf:"bytes,1,opt,name=currency_from,json=currencyFrom,proto3" json:"currency_from,omitempty"`
	CurrencyTo   string           `protobuf:"bytes,2,opt,name=currency_to,json=currencyTo,proto3" json:"currency_to,omitempty"`
	NettingRatio float64          `protobuf:"fixed64,3,opt,name=netting_ratio,json=nettingRatio,proto3" json:"netting_ratio,omitempty"`
	Field        models.DataField `protobuf:"varint,4,opt,name=field,proto3,enum=kappa.DataField" json:"field,omitempty"`
}

func (x *GetBlendedExchangeRateRequest) Reset() {
	*x = GetBlendedExchangeRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlendedExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlendedExchangeRateRequest) ProtoMessage() {}

func (x *GetBlendedExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlendedExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*GetBlendedExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_exchange_rates_exchange_rates_proto_rawDescGZIP(), []int{4}
}

func (x *GetBlendedExchangeRateRequest) GetCurrencyFrom() string {
	if x != nil {
		return x.CurrencyFrom
	}
	return ""
}

func (x *GetBlendedExchangeRateRequest) GetCurrencyTo() string {
	if x != nil {
		return x.CurrencyTo
	}
	return ""
}

func (x *GetBlendedExchangeRateRequest) GetNettingRatio() float64 {
	if x != nil {
		return x.NettingRatio
	}
	return 0
}

func (x *GetBlendedExchangeRateRequest) GetField() models.DataField {
	if x != nil {
		return x.Field
	}
	return models.DataField(0)
}

type GetBlendedExchangeRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeRate *models.ExchangeRate `protobuf:"bytes,1,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"`
}

func (x *GetBlendedExchangeRateResponse) Reset() {
	*x = GetBlendedExchangeRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlendedExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlendedExchangeRateResponse) ProtoMessage() {}

func (x *GetBlendedExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlendedExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*GetBlendedExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_exchange_rates_exchange_rates_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlendedExchangeRateResponse) GetExchangeRate() *models.ExchangeRate {
	if x != nil {
		return x.ExchangeRate
	}
	return nil
}

var File_kappa_services_exchange_rates_exchange_rates_proto protoreflect.FileDescriptor

var file_kappa_services_exchange_rates_exchange_rates_proto_rawDesc = []byte{
	0x0a, 0x32, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x1a, 0x1e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x54, 0x6f, 0x22, 0x5c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x0d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x22, 0x61, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x42,
	0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x61, 0x70,
	0x70, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x6e,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x12, 0x26, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x5a, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x32, 0xfc, 0x02, 0x0a, 0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x26, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x30, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x42, 0x79, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x42,
	0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x2d, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_services_exchange_rates_exchange_rates_proto_rawDescOnce sync.Once
	file_kappa_services_exchange_rates_exchange_rates_proto_rawDescData = file_kappa_services_exchange_rates_exchange_rates_proto_rawDesc
)

func file_kappa_services_exchange_rates_exchange_rates_proto_rawDescGZIP() []byte {
	file_kappa_services_exchange_rates_exchange_rates_proto_rawDescOnce.Do(func() {
		file_kappa_services_exchange_rates_exchange_rates_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_services_exchange_rates_exchange_rates_proto_rawDescData)
	})
	return file_kappa_services_exchange_rates_exchange_rates_proto_rawDescData
}

var file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kappa_services_exchange_rates_exchange_rates_proto_goTypes = []interface{}{
	(*GetExchangeRateRequest)(nil),            // 0: exchange_rates.GetExchangeRateRequest
	(*GetExchangeRateResponse)(nil),           // 1: exchange_rates.GetExchangeRateResponse
	(*GetDataSetBySourceAndNameRequest)(nil),  // 2: exchange_rates.GetDataSetBySourceAndNameRequest
	(*GetDataSetBySourceAndNameResponse)(nil), // 3: exchange_rates.GetDataSetBySourceAndNameResponse
	(*GetBlendedExchangeRateRequest)(nil),     // 4: exchange_rates.GetBlendedExchangeRateRequest
	(*GetBlendedExchangeRateResponse)(nil),    // 5: exchange_rates.GetBlendedExchangeRateResponse
	(models.DataSource)(0),                    // 6: kappa.DataSource
	(models.DataField)(0),                     // 7: kappa.DataField
	(*models.ExchangeRateDataPoint)(nil),      // 8: kappa.ExchangeRateDataPoint
	(*models.DataSet)(nil),                    // 9: kappa.DataSet
	(*models.ExchangeRate)(nil),               // 10: kappa.ExchangeRate
}
var file_kappa_services_exchange_rates_exchange_rates_proto_depIdxs = []int32{
	6,  // 0: exchange_rates.GetExchangeRateRequest.dataset_source:type_name -> kappa.DataSource
	7,  // 1: exchange_rates.GetExchangeRateRequest.field:type_name -> kappa.DataField
	8,  // 2: exchange_rates.GetExchangeRateResponse.exchange_rate:type_name -> kappa.ExchangeRateDataPoint
	6,  // 3: exchange_rates.GetDataSetBySourceAndNameRequest.source:type_name -> kappa.DataSource
	9,  // 4: exchange_rates.GetDataSetBySourceAndNameResponse.dataset:type_name -> kappa.DataSet
	7,  // 5: exchange_rates.GetBlendedExchangeRateRequest.field:type_name -> kappa.DataField
	10, // 6: exchange_rates.GetBlendedExchangeRateResponse.exchange_rate:type_name -> kappa.ExchangeRate
	0,  // 7: exchange_rates.ExchangeRatesService.GetExchangeRate:input_type -> exchange_rates.GetExchangeRateRequest
	2,  // 8: exchange_rates.ExchangeRatesService.GetDataSetBySourceAndName:input_type -> exchange_rates.GetDataSetBySourceAndNameRequest
	4,  // 9: exchange_rates.ExchangeRatesService.GetBlendedExchangeRate:input_type -> exchange_rates.GetBlendedExchangeRateRequest
	1,  // 10: exchange_rates.ExchangeRatesService.GetExchangeRate:output_type -> exchange_rates.GetExchangeRateResponse
	3,  // 11: exchange_rates.ExchangeRatesService.GetDataSetBySourceAndName:output_type -> exchange_rates.GetDataSetBySourceAndNameResponse
	5,  // 12: exchange_rates.ExchangeRatesService.GetBlendedExchangeRate:output_type -> exchange_rates.GetBlendedExchangeRateResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_kappa_services_exchange_rates_exchange_rates_proto_init() }
func file_kappa_services_exchange_rates_exchange_rates_proto_init() {
	if File_kappa_services_exchange_rates_exchange_rates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExchangeRateRequest); i {
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
		file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExchangeRateResponse); i {
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
		file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataSetBySourceAndNameRequest); i {
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
		file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataSetBySourceAndNameResponse); i {
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
		file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlendedExchangeRateRequest); i {
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
		file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlendedExchangeRateResponse); i {
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
			RawDescriptor: file_kappa_services_exchange_rates_exchange_rates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kappa_services_exchange_rates_exchange_rates_proto_goTypes,
		DependencyIndexes: file_kappa_services_exchange_rates_exchange_rates_proto_depIdxs,
		MessageInfos:      file_kappa_services_exchange_rates_exchange_rates_proto_msgTypes,
	}.Build()
	File_kappa_services_exchange_rates_exchange_rates_proto = out.File
	file_kappa_services_exchange_rates_exchange_rates_proto_rawDesc = nil
	file_kappa_services_exchange_rates_exchange_rates_proto_goTypes = nil
	file_kappa_services_exchange_rates_exchange_rates_proto_depIdxs = nil
}

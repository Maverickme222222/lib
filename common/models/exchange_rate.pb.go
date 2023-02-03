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
// source: kappa/common/models/exchange_rate.proto

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

type ExchangeRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        float64   `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	CurrencyFrom *Currency `protobuf:"bytes,2,opt,name=currency_from,json=currencyFrom,proto3" json:"currency_from,omitempty"`
	CurrencyTo   *Currency `protobuf:"bytes,3,opt,name=currency_to,json=currencyTo,proto3" json:"currency_to,omitempty"`
}

func (x *ExchangeRate) Reset() {
	*x = ExchangeRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_exchange_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRate) ProtoMessage() {}

func (x *ExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_exchange_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRate.ProtoReflect.Descriptor instead.
func (*ExchangeRate) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_exchange_rate_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeRate) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ExchangeRate) GetCurrencyFrom() *Currency {
	if x != nil {
		return x.CurrencyFrom
	}
	return nil
}

func (x *ExchangeRate) GetCurrencyTo() *Currency {
	if x != nil {
		return x.CurrencyTo
	}
	return nil
}

type ExchangeRateDataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Dataset         *DataSet               `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`                   // the dataset this exchange rate comes from
	Field           DataField              `protobuf:"varint,3,opt,name=field,proto3,enum=kappa.DataField" json:"field,omitempty"` // the field this point represents, e.g. BID, ASK
	Value           float64                `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`                     // the value of the exchange rate
	CurrencyFrom    *Currency              `protobuf:"bytes,5,opt,name=currency_from,json=currencyFrom,proto3" json:"currency_from,omitempty"`
	CurrencyTo      *Currency              `protobuf:"bytes,6,opt,name=currency_to,json=currencyTo,proto3" json:"currency_to,omitempty"`
	ReferenceTime   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=reference_time,json=referenceTime,proto3" json:"reference_time,omitempty"`       // the reference time for this value
	ObservationTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=observation_time,json=observationTime,proto3" json:"observation_time,omitempty"` // the time this value was observed at
}

func (x *ExchangeRateDataPoint) Reset() {
	*x = ExchangeRateDataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_exchange_rate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRateDataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateDataPoint) ProtoMessage() {}

func (x *ExchangeRateDataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_exchange_rate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateDataPoint.ProtoReflect.Descriptor instead.
func (*ExchangeRateDataPoint) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_exchange_rate_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeRateDataPoint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExchangeRateDataPoint) GetDataset() *DataSet {
	if x != nil {
		return x.Dataset
	}
	return nil
}

func (x *ExchangeRateDataPoint) GetField() DataField {
	if x != nil {
		return x.Field
	}
	return DataField_DATA_FIELD_BID
}

func (x *ExchangeRateDataPoint) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ExchangeRateDataPoint) GetCurrencyFrom() *Currency {
	if x != nil {
		return x.CurrencyFrom
	}
	return nil
}

func (x *ExchangeRateDataPoint) GetCurrencyTo() *Currency {
	if x != nil {
		return x.CurrencyTo
	}
	return nil
}

func (x *ExchangeRateDataPoint) GetReferenceTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReferenceTime
	}
	return nil
}

func (x *ExchangeRateDataPoint) GetObservationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ObservationTime
	}
	return nil
}

var File_kappa_common_models_exchange_rate_proto protoreflect.FileDescriptor

var file_kappa_common_models_exchange_rate_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x1a, 0x1e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x34, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x30, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0a, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x15, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x65, 0x74, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x26, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x30, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x54, 0x6f, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x70, 0x70,
	0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69, 0x62,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61, 0x70, 0x70,
	0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_common_models_exchange_rate_proto_rawDescOnce sync.Once
	file_kappa_common_models_exchange_rate_proto_rawDescData = file_kappa_common_models_exchange_rate_proto_rawDesc
)

func file_kappa_common_models_exchange_rate_proto_rawDescGZIP() []byte {
	file_kappa_common_models_exchange_rate_proto_rawDescOnce.Do(func() {
		file_kappa_common_models_exchange_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_common_models_exchange_rate_proto_rawDescData)
	})
	return file_kappa_common_models_exchange_rate_proto_rawDescData
}

var file_kappa_common_models_exchange_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kappa_common_models_exchange_rate_proto_goTypes = []interface{}{
	(*ExchangeRate)(nil),          // 0: kappa.ExchangeRate
	(*ExchangeRateDataPoint)(nil), // 1: kappa.ExchangeRateDataPoint
	(*Currency)(nil),              // 2: kappa.Currency
	(*DataSet)(nil),               // 3: kappa.DataSet
	(DataField)(0),                // 4: kappa.DataField
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_kappa_common_models_exchange_rate_proto_depIdxs = []int32{
	2, // 0: kappa.ExchangeRate.currency_from:type_name -> kappa.Currency
	2, // 1: kappa.ExchangeRate.currency_to:type_name -> kappa.Currency
	3, // 2: kappa.ExchangeRateDataPoint.dataset:type_name -> kappa.DataSet
	4, // 3: kappa.ExchangeRateDataPoint.field:type_name -> kappa.DataField
	2, // 4: kappa.ExchangeRateDataPoint.currency_from:type_name -> kappa.Currency
	2, // 5: kappa.ExchangeRateDataPoint.currency_to:type_name -> kappa.Currency
	5, // 6: kappa.ExchangeRateDataPoint.reference_time:type_name -> google.protobuf.Timestamp
	5, // 7: kappa.ExchangeRateDataPoint.observation_time:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_kappa_common_models_exchange_rate_proto_init() }
func file_kappa_common_models_exchange_rate_proto_init() {
	if File_kappa_common_models_exchange_rate_proto != nil {
		return
	}
	file_kappa_common_models_data_proto_init()
	file_kappa_common_models_configuration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kappa_common_models_exchange_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRate); i {
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
		file_kappa_common_models_exchange_rate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRateDataPoint); i {
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
			RawDescriptor: file_kappa_common_models_exchange_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kappa_common_models_exchange_rate_proto_goTypes,
		DependencyIndexes: file_kappa_common_models_exchange_rate_proto_depIdxs,
		MessageInfos:      file_kappa_common_models_exchange_rate_proto_msgTypes,
	}.Build()
	File_kappa_common_models_exchange_rate_proto = out.File
	file_kappa_common_models_exchange_rate_proto_rawDesc = nil
	file_kappa_common_models_exchange_rate_proto_goTypes = nil
	file_kappa_common_models_exchange_rate_proto_depIdxs = nil
}

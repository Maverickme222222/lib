// ------------------------------------------------------------------------------
// <copyright company="Kappa Pay Inc.">
// Copyright (C) Kappa Pay Inc.  All rights reserved.
// Unauthorised copying of this file, via any medium, is strictly prohibited.
// Proprietary and confidential
// </copyright>
// <author>Flavio Rajta</author>
// ------------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: kappa/common/models/sila.proto

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

type SilaEventMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType string `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	EventUuid string `protobuf:"bytes,2,opt,name=event_uuid,json=eventUuid,proto3" json:"event_uuid,omitempty"`
}

func (x *SilaEventMetadata) Reset() {
	*x = SilaEventMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_sila_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SilaEventMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SilaEventMetadata) ProtoMessage() {}

func (x *SilaEventMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_sila_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SilaEventMetadata.ProtoReflect.Descriptor instead.
func (*SilaEventMetadata) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_sila_proto_rawDescGZIP(), []int{0}
}

func (x *SilaEventMetadata) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *SilaEventMetadata) GetEventUuid() string {
	if x != nil {
		return x.EventUuid
	}
	return ""
}

// SilaTransactionUpdateEvent is the request data from sila transaction_update webhook
// docs: https://docs.silamoney.com/docs/webhook-event-reference#transaction-status-update-transaction_update
type SilaTransactionUpdateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata        *SilaEventMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Transaction     string             `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`                                // transaction id
	TransactionType string             `protobuf:"bytes,3,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"` // issue,transfer,redeem
	SilaAmount      int32              `protobuf:"varint,4,opt,name=sila_amount,json=silaAmount,proto3" json:"sila_amount,omitempty"`
	Outcome         string             `protobuf:"bytes,5,opt,name=outcome,proto3" json:"outcome,omitempty"` // success,failed,review,queued
	Entity          string             `protobuf:"bytes,6,opt,name=entity,proto3" json:"entity,omitempty"`
	ProcessingType  string             `protobuf:"bytes,7,opt,name=processing_type,json=processingType,proto3" json:"processing_type,omitempty"`
	ProviderStatus  string             `protobuf:"bytes,8,opt,name=provider_status,json=providerStatus,proto3" json:"provider_status,omitempty"`
}

func (x *SilaTransactionUpdateEvent) Reset() {
	*x = SilaTransactionUpdateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_sila_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SilaTransactionUpdateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SilaTransactionUpdateEvent) ProtoMessage() {}

func (x *SilaTransactionUpdateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_sila_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SilaTransactionUpdateEvent.ProtoReflect.Descriptor instead.
func (*SilaTransactionUpdateEvent) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_sila_proto_rawDescGZIP(), []int{1}
}

func (x *SilaTransactionUpdateEvent) GetMetadata() *SilaEventMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SilaTransactionUpdateEvent) GetTransaction() string {
	if x != nil {
		return x.Transaction
	}
	return ""
}

func (x *SilaTransactionUpdateEvent) GetTransactionType() string {
	if x != nil {
		return x.TransactionType
	}
	return ""
}

func (x *SilaTransactionUpdateEvent) GetSilaAmount() int32 {
	if x != nil {
		return x.SilaAmount
	}
	return 0
}

func (x *SilaTransactionUpdateEvent) GetOutcome() string {
	if x != nil {
		return x.Outcome
	}
	return ""
}

func (x *SilaTransactionUpdateEvent) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *SilaTransactionUpdateEvent) GetProcessingType() string {
	if x != nil {
		return x.ProcessingType
	}
	return ""
}

func (x *SilaTransactionUpdateEvent) GetProviderStatus() string {
	if x != nil {
		return x.ProviderStatus
	}
	return ""
}

var File_kappa_common_models_sila_proto protoreflect.FileDescriptor

var file_kappa_common_models_sila_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x69, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x22, 0x51, 0x0a, 0x11, 0x53, 0x69, 0x6c, 0x61, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0xc4, 0x02, 0x0a, 0x1a, 0x53,
	0x69, 0x6c, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2e, 0x53, 0x69, 0x6c, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x69, 0x6c, 0x61, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x69, 0x6c, 0x61, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_common_models_sila_proto_rawDescOnce sync.Once
	file_kappa_common_models_sila_proto_rawDescData = file_kappa_common_models_sila_proto_rawDesc
)

func file_kappa_common_models_sila_proto_rawDescGZIP() []byte {
	file_kappa_common_models_sila_proto_rawDescOnce.Do(func() {
		file_kappa_common_models_sila_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_common_models_sila_proto_rawDescData)
	})
	return file_kappa_common_models_sila_proto_rawDescData
}

var file_kappa_common_models_sila_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kappa_common_models_sila_proto_goTypes = []interface{}{
	(*SilaEventMetadata)(nil),          // 0: kappa.SilaEventMetadata
	(*SilaTransactionUpdateEvent)(nil), // 1: kappa.SilaTransactionUpdateEvent
}
var file_kappa_common_models_sila_proto_depIdxs = []int32{
	0, // 0: kappa.SilaTransactionUpdateEvent.metadata:type_name -> kappa.SilaEventMetadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kappa_common_models_sila_proto_init() }
func file_kappa_common_models_sila_proto_init() {
	if File_kappa_common_models_sila_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kappa_common_models_sila_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SilaEventMetadata); i {
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
		file_kappa_common_models_sila_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SilaTransactionUpdateEvent); i {
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
			RawDescriptor: file_kappa_common_models_sila_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kappa_common_models_sila_proto_goTypes,
		DependencyIndexes: file_kappa_common_models_sila_proto_depIdxs,
		MessageInfos:      file_kappa_common_models_sila_proto_msgTypes,
	}.Build()
	File_kappa_common_models_sila_proto = out.File
	file_kappa_common_models_sila_proto_rawDesc = nil
	file_kappa_common_models_sila_proto_goTypes = nil
	file_kappa_common_models_sila_proto_depIdxs = nil
}

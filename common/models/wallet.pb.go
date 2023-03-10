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
// source: kappa/common/models/wallet.proto

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

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                       // human-friendly name for frontend display
	Entity       *Entity   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`                                   // the entity controlling the wallet
	HostEntityId string    `protobuf:"bytes,4,opt,name=host_entity_id,json=hostEntityId,proto3" json:"host_entity_id,omitempty"` // the hosting entity which maintains the wallet, e.g. Kappa Pay Inc.
	Currency     *Currency `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`                               // the currency of the wallet
	LedgerName   *string   `protobuf:"bytes,6,opt,name=ledger_name,json=ledgerName,proto3,oneof" json:"ledger_name,omitempty"`   // the name of the ledger that this entity's wallet is created in
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *Wallet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Wallet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Wallet) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *Wallet) GetHostEntityId() string {
	if x != nil {
		return x.HostEntityId
	}
	return ""
}

func (x *Wallet) GetCurrency() *Currency {
	if x != nil {
		return x.Currency
	}
	return nil
}

func (x *Wallet) GetLedgerName() string {
	if x != nil && x.LedgerName != nil {
		return *x.LedgerName
	}
	return ""
}

var File_kappa_common_models_wallet_proto protoreflect.FileDescriptor

var file_kappa_common_models_wallet_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x1a, 0x22, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x0e,
	0x68, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x24, 0x0a, 0x0b, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_common_models_wallet_proto_rawDescOnce sync.Once
	file_kappa_common_models_wallet_proto_rawDescData = file_kappa_common_models_wallet_proto_rawDesc
)

func file_kappa_common_models_wallet_proto_rawDescGZIP() []byte {
	file_kappa_common_models_wallet_proto_rawDescOnce.Do(func() {
		file_kappa_common_models_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_common_models_wallet_proto_rawDescData)
	})
	return file_kappa_common_models_wallet_proto_rawDescData
}

var file_kappa_common_models_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kappa_common_models_wallet_proto_goTypes = []interface{}{
	(*Wallet)(nil),   // 0: kappa.Wallet
	(*Entity)(nil),   // 1: kappa.Entity
	(*Currency)(nil), // 2: kappa.Currency
}
var file_kappa_common_models_wallet_proto_depIdxs = []int32{
	1, // 0: kappa.Wallet.entity:type_name -> kappa.Entity
	2, // 1: kappa.Wallet.currency:type_name -> kappa.Currency
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kappa_common_models_wallet_proto_init() }
func file_kappa_common_models_wallet_proto_init() {
	if File_kappa_common_models_wallet_proto != nil {
		return
	}
	file_kappa_common_models_entities_proto_init()
	file_kappa_common_models_configuration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kappa_common_models_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
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
	file_kappa_common_models_wallet_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kappa_common_models_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kappa_common_models_wallet_proto_goTypes,
		DependencyIndexes: file_kappa_common_models_wallet_proto_depIdxs,
		MessageInfos:      file_kappa_common_models_wallet_proto_msgTypes,
	}.Build()
	File_kappa_common_models_wallet_proto = out.File
	file_kappa_common_models_wallet_proto_rawDesc = nil
	file_kappa_common_models_wallet_proto_goTypes = nil
	file_kappa_common_models_wallet_proto_depIdxs = nil
}

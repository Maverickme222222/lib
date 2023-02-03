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
// source: kappa/common/models/identity.proto

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

type IdentityVerificationStatus int32

const (
	IdentityVerificationStatus_IDENTITY_VERIFICATION_STATUS_UNSPECIFIED IdentityVerificationStatus = 0
	IdentityVerificationStatus_IDENTITY_VERIFICATION_STATUS_CREATED     IdentityVerificationStatus = 1
	IdentityVerificationStatus_IDENTITY_VERIFICATION_STATUS_PENDING     IdentityVerificationStatus = 2
	IdentityVerificationStatus_IDENTITY_VERIFICATION_STATUS_PROCESSING  IdentityVerificationStatus = 3
	IdentityVerificationStatus_IDENTITY_VERIFICATION_STATUS_COMPLETE    IdentityVerificationStatus = 4
	IdentityVerificationStatus_IDENTITY_VERIFICATION_STATUS_ERROR       IdentityVerificationStatus = 5
)

// Enum value maps for IdentityVerificationStatus.
var (
	IdentityVerificationStatus_name = map[int32]string{
		0: "IDENTITY_VERIFICATION_STATUS_UNSPECIFIED",
		1: "IDENTITY_VERIFICATION_STATUS_CREATED",
		2: "IDENTITY_VERIFICATION_STATUS_PENDING",
		3: "IDENTITY_VERIFICATION_STATUS_PROCESSING",
		4: "IDENTITY_VERIFICATION_STATUS_COMPLETE",
		5: "IDENTITY_VERIFICATION_STATUS_ERROR",
	}
	IdentityVerificationStatus_value = map[string]int32{
		"IDENTITY_VERIFICATION_STATUS_UNSPECIFIED": 0,
		"IDENTITY_VERIFICATION_STATUS_CREATED":     1,
		"IDENTITY_VERIFICATION_STATUS_PENDING":     2,
		"IDENTITY_VERIFICATION_STATUS_PROCESSING":  3,
		"IDENTITY_VERIFICATION_STATUS_COMPLETE":    4,
		"IDENTITY_VERIFICATION_STATUS_ERROR":       5,
	}
)

func (x IdentityVerificationStatus) Enum() *IdentityVerificationStatus {
	p := new(IdentityVerificationStatus)
	*p = x
	return p
}

func (x IdentityVerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityVerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_identity_proto_enumTypes[0].Descriptor()
}

func (IdentityVerificationStatus) Type() protoreflect.EnumType {
	return &file_kappa_common_models_identity_proto_enumTypes[0]
}

func (x IdentityVerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityVerificationStatus.Descriptor instead.
func (IdentityVerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_identity_proto_rawDescGZIP(), []int{0}
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value              string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Type               *IdentityType          `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	EntityId           string                 `protobuf:"bytes,4,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Verified           bool                   `protobuf:"varint,5,opt,name=verified,proto3" json:"verified,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt          *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	LatestVerification *IdentityVerification  `protobuf:"bytes,10,opt,name=latest_verification,json=latestVerification,proto3" json:"latest_verification,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_identity_proto_rawDescGZIP(), []int{0}
}

func (x *Identity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Identity) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Identity) GetType() *IdentityType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Identity) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *Identity) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *Identity) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Identity) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Identity) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Identity) GetLatestVerification() *IdentityVerification {
	if x != nil {
		return x.LatestVerification
	}
	return nil
}

type IdentityDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IdentityId   string                 `protobuf:"bytes,2,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`
	Name         string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type         *IdentityDocumentType  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	DocumentLink string                 `protobuf:"bytes,5,opt,name=document_link,json=documentLink,proto3" json:"document_link,omitempty"`
	Verified     bool                   `protobuf:"varint,6,opt,name=verified,proto3" json:"verified,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *IdentityDocument) Reset() {
	*x = IdentityDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityDocument) ProtoMessage() {}

func (x *IdentityDocument) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityDocument.ProtoReflect.Descriptor instead.
func (*IdentityDocument) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_identity_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityDocument) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IdentityDocument) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

func (x *IdentityDocument) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IdentityDocument) GetType() *IdentityDocumentType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *IdentityDocument) GetDocumentLink() string {
	if x != nil {
		return x.DocumentLink
	}
	return ""
}

func (x *IdentityDocument) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *IdentityDocument) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IdentityDocument) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *IdentityDocument) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type IdentityVerification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IdentityId   string                     `protobuf:"bytes,2,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`
	VendorHandle string                     `protobuf:"bytes,3,opt,name=vendor_handle,json=vendorHandle,proto3" json:"vendor_handle,omitempty"` // identifier used by the vendor who performed the verification
	EntityId     string                     `protobuf:"bytes,4,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Status       IdentityVerificationStatus `protobuf:"varint,5,opt,name=status,proto3,enum=kappa.IdentityVerificationStatus" json:"status,omitempty"`
	CreatedAt    *timestamppb.Timestamp     `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	RequestedAt  *timestamppb.Timestamp     `protobuf:"bytes,7,opt,name=requested_at,json=requestedAt,proto3" json:"requested_at,omitempty"`
	SubmittedAt  *timestamppb.Timestamp     `protobuf:"bytes,8,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	CompletedAt  *timestamppb.Timestamp     `protobuf:"bytes,9,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	Attempt      int64                      `protobuf:"varint,10,opt,name=attempt,proto3" json:"attempt,omitempty"`
}

func (x *IdentityVerification) Reset() {
	*x = IdentityVerification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityVerification) ProtoMessage() {}

func (x *IdentityVerification) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityVerification.ProtoReflect.Descriptor instead.
func (*IdentityVerification) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_identity_proto_rawDescGZIP(), []int{2}
}

func (x *IdentityVerification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IdentityVerification) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

func (x *IdentityVerification) GetVendorHandle() string {
	if x != nil {
		return x.VendorHandle
	}
	return ""
}

func (x *IdentityVerification) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *IdentityVerification) GetStatus() IdentityVerificationStatus {
	if x != nil {
		return x.Status
	}
	return IdentityVerificationStatus_IDENTITY_VERIFICATION_STATUS_UNSPECIFIED
}

func (x *IdentityVerification) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IdentityVerification) GetRequestedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedAt
	}
	return nil
}

func (x *IdentityVerification) GetSubmittedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmittedAt
	}
	return nil
}

func (x *IdentityVerification) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *IdentityVerification) GetAttempt() int64 {
	if x != nil {
		return x.Attempt
	}
	return 0
}

var File_kappa_common_models_identity_proto protoreflect.FileDescriptor

var file_kappa_common_models_identity_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4c, 0x0a, 0x13, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfa, 0x02, 0x0a, 0x10, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd6, 0x03, 0x0a, 0x14, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x2a,
	0x9e, 0x02, 0x0a, 0x1a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c,
	0x0a, 0x28, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24,
	0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x28, 0x0a, 0x24, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x2b, 0x0a, 0x27, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52,
	0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x29, 0x0a,
	0x25, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x26, 0x0a, 0x22, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05,
	0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_common_models_identity_proto_rawDescOnce sync.Once
	file_kappa_common_models_identity_proto_rawDescData = file_kappa_common_models_identity_proto_rawDesc
)

func file_kappa_common_models_identity_proto_rawDescGZIP() []byte {
	file_kappa_common_models_identity_proto_rawDescOnce.Do(func() {
		file_kappa_common_models_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_common_models_identity_proto_rawDescData)
	})
	return file_kappa_common_models_identity_proto_rawDescData
}

var file_kappa_common_models_identity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kappa_common_models_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kappa_common_models_identity_proto_goTypes = []interface{}{
	(IdentityVerificationStatus)(0), // 0: kappa.IdentityVerificationStatus
	(*Identity)(nil),                // 1: kappa.Identity
	(*IdentityDocument)(nil),        // 2: kappa.IdentityDocument
	(*IdentityVerification)(nil),    // 3: kappa.IdentityVerification
	(*IdentityType)(nil),            // 4: kappa.IdentityType
	(*timestamppb.Timestamp)(nil),   // 5: google.protobuf.Timestamp
	(*IdentityDocumentType)(nil),    // 6: kappa.IdentityDocumentType
}
var file_kappa_common_models_identity_proto_depIdxs = []int32{
	4,  // 0: kappa.Identity.type:type_name -> kappa.IdentityType
	5,  // 1: kappa.Identity.created_at:type_name -> google.protobuf.Timestamp
	5,  // 2: kappa.Identity.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 3: kappa.Identity.deleted_at:type_name -> google.protobuf.Timestamp
	3,  // 4: kappa.Identity.latest_verification:type_name -> kappa.IdentityVerification
	6,  // 5: kappa.IdentityDocument.type:type_name -> kappa.IdentityDocumentType
	5,  // 6: kappa.IdentityDocument.created_at:type_name -> google.protobuf.Timestamp
	5,  // 7: kappa.IdentityDocument.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 8: kappa.IdentityDocument.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 9: kappa.IdentityVerification.status:type_name -> kappa.IdentityVerificationStatus
	5,  // 10: kappa.IdentityVerification.created_at:type_name -> google.protobuf.Timestamp
	5,  // 11: kappa.IdentityVerification.requested_at:type_name -> google.protobuf.Timestamp
	5,  // 12: kappa.IdentityVerification.submitted_at:type_name -> google.protobuf.Timestamp
	5,  // 13: kappa.IdentityVerification.completed_at:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_kappa_common_models_identity_proto_init() }
func file_kappa_common_models_identity_proto_init() {
	if File_kappa_common_models_identity_proto != nil {
		return
	}
	file_kappa_common_models_configuration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kappa_common_models_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_kappa_common_models_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityDocument); i {
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
		file_kappa_common_models_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityVerification); i {
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
			RawDescriptor: file_kappa_common_models_identity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kappa_common_models_identity_proto_goTypes,
		DependencyIndexes: file_kappa_common_models_identity_proto_depIdxs,
		EnumInfos:         file_kappa_common_models_identity_proto_enumTypes,
		MessageInfos:      file_kappa_common_models_identity_proto_msgTypes,
	}.Build()
	File_kappa_common_models_identity_proto = out.File
	file_kappa_common_models_identity_proto_rawDesc = nil
	file_kappa_common_models_identity_proto_goTypes = nil
	file_kappa_common_models_identity_proto_depIdxs = nil
}

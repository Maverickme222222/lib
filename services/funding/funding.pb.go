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
// source: kappa/services/funding/funding.proto

package funding

import (
	models "github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAvailableBalanceForSourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
}

func (x *GetAvailableBalanceForSourceRequest) Reset() {
	*x = GetAvailableBalanceForSourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableBalanceForSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableBalanceForSourceRequest) ProtoMessage() {}

func (x *GetAvailableBalanceForSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableBalanceForSourceRequest.ProtoReflect.Descriptor instead.
func (*GetAvailableBalanceForSourceRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{0}
}

func (x *GetAvailableBalanceForSourceRequest) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

type GetAvailableBalanceForSourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableBalance *models.Amount `protobuf:"bytes,1,opt,name=available_balance,json=availableBalance,proto3" json:"available_balance,omitempty"`
}

func (x *GetAvailableBalanceForSourceResponse) Reset() {
	*x = GetAvailableBalanceForSourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableBalanceForSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableBalanceForSourceResponse) ProtoMessage() {}

func (x *GetAvailableBalanceForSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableBalanceForSourceResponse.ProtoReflect.Descriptor instead.
func (*GetAvailableBalanceForSourceResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{1}
}

func (x *GetAvailableBalanceForSourceResponse) GetAvailableBalance() *models.Amount {
	if x != nil {
		return x.AvailableBalance
	}
	return nil
}

type ProcessSilaTransactionUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *models.SilaTransactionUpdateEvent `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *ProcessSilaTransactionUpdateRequest) Reset() {
	*x = ProcessSilaTransactionUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessSilaTransactionUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessSilaTransactionUpdateRequest) ProtoMessage() {}

func (x *ProcessSilaTransactionUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessSilaTransactionUpdateRequest.ProtoReflect.Descriptor instead.
func (*ProcessSilaTransactionUpdateRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessSilaTransactionUpdateRequest) GetEvent() *models.SilaTransactionUpdateEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type GetFundingRequestByTransactionReferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionReference string `protobuf:"bytes,1,opt,name=transaction_reference,json=transactionReference,proto3" json:"transaction_reference,omitempty"`
	PaymentGatewayId     string `protobuf:"bytes,2,opt,name=payment_gateway_id,json=paymentGatewayId,proto3" json:"payment_gateway_id,omitempty"`
}

func (x *GetFundingRequestByTransactionReferenceRequest) Reset() {
	*x = GetFundingRequestByTransactionReferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundingRequestByTransactionReferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundingRequestByTransactionReferenceRequest) ProtoMessage() {}

func (x *GetFundingRequestByTransactionReferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundingRequestByTransactionReferenceRequest.ProtoReflect.Descriptor instead.
func (*GetFundingRequestByTransactionReferenceRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{3}
}

func (x *GetFundingRequestByTransactionReferenceRequest) GetTransactionReference() string {
	if x != nil {
		return x.TransactionReference
	}
	return ""
}

func (x *GetFundingRequestByTransactionReferenceRequest) GetPaymentGatewayId() string {
	if x != nil {
		return x.PaymentGatewayId
	}
	return ""
}

type GetFundingRequestByTransactionReferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FundingRequest *models.FundingRequest `protobuf:"bytes,1,opt,name=funding_request,json=fundingRequest,proto3" json:"funding_request,omitempty"`
}

func (x *GetFundingRequestByTransactionReferenceResponse) Reset() {
	*x = GetFundingRequestByTransactionReferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFundingRequestByTransactionReferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFundingRequestByTransactionReferenceResponse) ProtoMessage() {}

func (x *GetFundingRequestByTransactionReferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFundingRequestByTransactionReferenceResponse.ProtoReflect.Descriptor instead.
func (*GetFundingRequestByTransactionReferenceResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{4}
}

func (x *GetFundingRequestByTransactionReferenceResponse) GetFundingRequest() *models.FundingRequest {
	if x != nil {
		return x.FundingRequest
	}
	return nil
}

type UpdateFundingRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FundingRequest *models.FundingRequest `protobuf:"bytes,1,opt,name=funding_request,json=fundingRequest,proto3" json:"funding_request,omitempty"`
}

func (x *UpdateFundingRequestRequest) Reset() {
	*x = UpdateFundingRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFundingRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFundingRequestRequest) ProtoMessage() {}

func (x *UpdateFundingRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFundingRequestRequest.ProtoReflect.Descriptor instead.
func (*UpdateFundingRequestRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFundingRequestRequest) GetFundingRequest() *models.FundingRequest {
	if x != nil {
		return x.FundingRequest
	}
	return nil
}

type UpdateFundingRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FundingRequest *models.FundingRequest `protobuf:"bytes,1,opt,name=funding_request,json=fundingRequest,proto3" json:"funding_request,omitempty"`
}

func (x *UpdateFundingRequestResponse) Reset() {
	*x = UpdateFundingRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFundingRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFundingRequestResponse) ProtoMessage() {}

func (x *UpdateFundingRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFundingRequestResponse.ProtoReflect.Descriptor instead.
func (*UpdateFundingRequestResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFundingRequestResponse) GetFundingRequest() *models.FundingRequest {
	if x != nil {
		return x.FundingRequest
	}
	return nil
}

type InitiateFundingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string  `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	SessionId     *string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3,oneof" json:"session_id,omitempty"`
}

func (x *InitiateFundingRequest) Reset() {
	*x = InitiateFundingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateFundingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateFundingRequest) ProtoMessage() {}

func (x *InitiateFundingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateFundingRequest.ProtoReflect.Descriptor instead.
func (*InitiateFundingRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{7}
}

func (x *InitiateFundingRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *InitiateFundingRequest) GetSessionId() string {
	if x != nil && x.SessionId != nil {
		return *x.SessionId
	}
	return ""
}

type InitiateFundingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FundingRequest *models.FundingRequest `protobuf:"bytes,1,opt,name=funding_request,json=fundingRequest,proto3" json:"funding_request,omitempty"`
	Limits         []*models.Limit        `protobuf:"bytes,2,rep,name=limits,proto3" json:"limits,omitempty"` // Limits are returned when the funding request is not created because the transaction breaches one or more limits
}

func (x *InitiateFundingResponse) Reset() {
	*x = InitiateFundingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_funding_funding_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateFundingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateFundingResponse) ProtoMessage() {}

func (x *InitiateFundingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_funding_funding_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateFundingResponse.ProtoReflect.Descriptor instead.
func (*InitiateFundingResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_funding_funding_proto_rawDescGZIP(), []int{8}
}

func (x *InitiateFundingResponse) GetFundingRequest() *models.FundingRequest {
	if x != nil {
		return x.FundingRequest
	}
	return nil
}

func (x *InitiateFundingResponse) GetLimits() []*models.Limit {
	if x != nil {
		return x.Limits
	}
	return nil
}

var File_kappa_services_funding_funding_proto protoreflect.FileDescriptor

var file_kappa_services_funding_funding_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x1a,
	0x20, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x69, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x42, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6b, 0x61, 0x70, 0x70,
	0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x5e, 0x0a, 0x23, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x6c, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x53, 0x69, 0x6c, 0x61, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64,
	0x22, 0x71, 0x0a, 0x2f, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x5e, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x72, 0x0a, 0x16, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x46, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x7f, 0x0a, 0x17, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x61, 0x70,
	0x70, 0x61, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52,
	0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x32, 0xcd, 0x04, 0x0a, 0x0e, 0x46, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x46, 0x6f, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2c, 0x2e, 0x66, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x69, 0x6c, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x6c, 0x61, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x9c, 0x01,
	0x0a, 0x27, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x37, 0x2e, 0x66, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x38, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x54, 0x0a, 0x0f, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_services_funding_funding_proto_rawDescOnce sync.Once
	file_kappa_services_funding_funding_proto_rawDescData = file_kappa_services_funding_funding_proto_rawDesc
)

func file_kappa_services_funding_funding_proto_rawDescGZIP() []byte {
	file_kappa_services_funding_funding_proto_rawDescOnce.Do(func() {
		file_kappa_services_funding_funding_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_services_funding_funding_proto_rawDescData)
	})
	return file_kappa_services_funding_funding_proto_rawDescData
}

var file_kappa_services_funding_funding_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_kappa_services_funding_funding_proto_goTypes = []interface{}{
	(*GetAvailableBalanceForSourceRequest)(nil),             // 0: funding.GetAvailableBalanceForSourceRequest
	(*GetAvailableBalanceForSourceResponse)(nil),            // 1: funding.GetAvailableBalanceForSourceResponse
	(*ProcessSilaTransactionUpdateRequest)(nil),             // 2: funding.ProcessSilaTransactionUpdateRequest
	(*GetFundingRequestByTransactionReferenceRequest)(nil),  // 3: funding.GetFundingRequestByTransactionReferenceRequest
	(*GetFundingRequestByTransactionReferenceResponse)(nil), // 4: funding.GetFundingRequestByTransactionReferenceResponse
	(*UpdateFundingRequestRequest)(nil),                     // 5: funding.UpdateFundingRequestRequest
	(*UpdateFundingRequestResponse)(nil),                    // 6: funding.UpdateFundingRequestResponse
	(*InitiateFundingRequest)(nil),                          // 7: funding.InitiateFundingRequest
	(*InitiateFundingResponse)(nil),                         // 8: funding.InitiateFundingResponse
	(*models.Amount)(nil),                                   // 9: kappa.Amount
	(*models.SilaTransactionUpdateEvent)(nil),               // 10: kappa.SilaTransactionUpdateEvent
	(*models.FundingRequest)(nil),                           // 11: kappa.FundingRequest
	(*models.Limit)(nil),                                    // 12: kappa.Limit
	(*emptypb.Empty)(nil),                                   // 13: google.protobuf.Empty
}
var file_kappa_services_funding_funding_proto_depIdxs = []int32{
	9,  // 0: funding.GetAvailableBalanceForSourceResponse.available_balance:type_name -> kappa.Amount
	10, // 1: funding.ProcessSilaTransactionUpdateRequest.event:type_name -> kappa.SilaTransactionUpdateEvent
	11, // 2: funding.GetFundingRequestByTransactionReferenceResponse.funding_request:type_name -> kappa.FundingRequest
	11, // 3: funding.UpdateFundingRequestRequest.funding_request:type_name -> kappa.FundingRequest
	11, // 4: funding.UpdateFundingRequestResponse.funding_request:type_name -> kappa.FundingRequest
	11, // 5: funding.InitiateFundingResponse.funding_request:type_name -> kappa.FundingRequest
	12, // 6: funding.InitiateFundingResponse.limits:type_name -> kappa.Limit
	0,  // 7: funding.FundingService.GetAvailableBalanceForSource:input_type -> funding.GetAvailableBalanceForSourceRequest
	2,  // 8: funding.FundingService.ProcessSilaTransactionUpdate:input_type -> funding.ProcessSilaTransactionUpdateRequest
	3,  // 9: funding.FundingService.GetFundingRequestByTransactionReference:input_type -> funding.GetFundingRequestByTransactionReferenceRequest
	5,  // 10: funding.FundingService.UpdateFundingRequest:input_type -> funding.UpdateFundingRequestRequest
	7,  // 11: funding.FundingService.InitiateFunding:input_type -> funding.InitiateFundingRequest
	1,  // 12: funding.FundingService.GetAvailableBalanceForSource:output_type -> funding.GetAvailableBalanceForSourceResponse
	13, // 13: funding.FundingService.ProcessSilaTransactionUpdate:output_type -> google.protobuf.Empty
	4,  // 14: funding.FundingService.GetFundingRequestByTransactionReference:output_type -> funding.GetFundingRequestByTransactionReferenceResponse
	6,  // 15: funding.FundingService.UpdateFundingRequest:output_type -> funding.UpdateFundingRequestResponse
	8,  // 16: funding.FundingService.InitiateFunding:output_type -> funding.InitiateFundingResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_kappa_services_funding_funding_proto_init() }
func file_kappa_services_funding_funding_proto_init() {
	if File_kappa_services_funding_funding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kappa_services_funding_funding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailableBalanceForSourceRequest); i {
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
		file_kappa_services_funding_funding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailableBalanceForSourceResponse); i {
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
		file_kappa_services_funding_funding_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessSilaTransactionUpdateRequest); i {
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
		file_kappa_services_funding_funding_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundingRequestByTransactionReferenceRequest); i {
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
		file_kappa_services_funding_funding_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFundingRequestByTransactionReferenceResponse); i {
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
		file_kappa_services_funding_funding_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFundingRequestRequest); i {
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
		file_kappa_services_funding_funding_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFundingRequestResponse); i {
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
		file_kappa_services_funding_funding_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateFundingRequest); i {
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
		file_kappa_services_funding_funding_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateFundingResponse); i {
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
	file_kappa_services_funding_funding_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kappa_services_funding_funding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kappa_services_funding_funding_proto_goTypes,
		DependencyIndexes: file_kappa_services_funding_funding_proto_depIdxs,
		MessageInfos:      file_kappa_services_funding_funding_proto_msgTypes,
	}.Build()
	File_kappa_services_funding_funding_proto = out.File
	file_kappa_services_funding_funding_proto_rawDesc = nil
	file_kappa_services_funding_funding_proto_goTypes = nil
	file_kappa_services_funding_funding_proto_depIdxs = nil
}

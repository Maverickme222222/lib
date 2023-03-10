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
// source: kappa/services/payout/payout.proto

package payout

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

type ProcessSilaTransactionUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *models.SilaTransactionUpdateEvent `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *ProcessSilaTransactionUpdateRequest) Reset() {
	*x = ProcessSilaTransactionUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessSilaTransactionUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessSilaTransactionUpdateRequest) ProtoMessage() {}

func (x *ProcessSilaTransactionUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[0]
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
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessSilaTransactionUpdateRequest) GetEvent() *models.SilaTransactionUpdateEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type CreatePayoutRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayoutRequest *models.PayoutRequest `protobuf:"bytes,1,opt,name=payout_request,json=payoutRequest,proto3" json:"payout_request,omitempty"`
}

func (x *CreatePayoutRequestRequest) Reset() {
	*x = CreatePayoutRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePayoutRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePayoutRequestRequest) ProtoMessage() {}

func (x *CreatePayoutRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePayoutRequestRequest.ProtoReflect.Descriptor instead.
func (*CreatePayoutRequestRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePayoutRequestRequest) GetPayoutRequest() *models.PayoutRequest {
	if x != nil {
		return x.PayoutRequest
	}
	return nil
}

type CreatePayoutRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayoutRequest *models.PayoutRequest `protobuf:"bytes,1,opt,name=payout_request,json=payoutRequest,proto3" json:"payout_request,omitempty"`
}

func (x *CreatePayoutRequestResponse) Reset() {
	*x = CreatePayoutRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePayoutRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePayoutRequestResponse) ProtoMessage() {}

func (x *CreatePayoutRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePayoutRequestResponse.ProtoReflect.Descriptor instead.
func (*CreatePayoutRequestResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePayoutRequestResponse) GetPayoutRequest() *models.PayoutRequest {
	if x != nil {
		return x.PayoutRequest
	}
	return nil
}

type UpdatePayoutRequestStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status models.TransactionPayoutStatus `protobuf:"varint,2,opt,name=status,proto3,enum=kappa.TransactionPayoutStatus" json:"status,omitempty"`
	Data   string                         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"` // json text to save webhook event details
}

func (x *UpdatePayoutRequestStatusRequest) Reset() {
	*x = UpdatePayoutRequestStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePayoutRequestStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePayoutRequestStatusRequest) ProtoMessage() {}

func (x *UpdatePayoutRequestStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePayoutRequestStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdatePayoutRequestStatusRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePayoutRequestStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePayoutRequestStatusRequest) GetStatus() models.TransactionPayoutStatus {
	if x != nil {
		return x.Status
	}
	return models.TransactionPayoutStatus(0)
}

func (x *UpdatePayoutRequestStatusRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type UpdatePayoutRequestStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayoutRequest *models.PayoutRequest `protobuf:"bytes,1,opt,name=payout_request,json=payoutRequest,proto3" json:"payout_request,omitempty"`
}

func (x *UpdatePayoutRequestStatusResponse) Reset() {
	*x = UpdatePayoutRequestStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePayoutRequestStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePayoutRequestStatusResponse) ProtoMessage() {}

func (x *UpdatePayoutRequestStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePayoutRequestStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdatePayoutRequestStatusResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePayoutRequestStatusResponse) GetPayoutRequest() *models.PayoutRequest {
	if x != nil {
		return x.PayoutRequest
	}
	return nil
}

type GetPayoutRequestByTransactionReferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionReference string `protobuf:"bytes,1,opt,name=transaction_reference,json=transactionReference,proto3" json:"transaction_reference,omitempty"`
	PaymentGatewayId     string `protobuf:"bytes,2,opt,name=payment_gateway_id,json=paymentGatewayId,proto3" json:"payment_gateway_id,omitempty"`
}

func (x *GetPayoutRequestByTransactionReferenceRequest) Reset() {
	*x = GetPayoutRequestByTransactionReferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPayoutRequestByTransactionReferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPayoutRequestByTransactionReferenceRequest) ProtoMessage() {}

func (x *GetPayoutRequestByTransactionReferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPayoutRequestByTransactionReferenceRequest.ProtoReflect.Descriptor instead.
func (*GetPayoutRequestByTransactionReferenceRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{5}
}

func (x *GetPayoutRequestByTransactionReferenceRequest) GetTransactionReference() string {
	if x != nil {
		return x.TransactionReference
	}
	return ""
}

func (x *GetPayoutRequestByTransactionReferenceRequest) GetPaymentGatewayId() string {
	if x != nil {
		return x.PaymentGatewayId
	}
	return ""
}

type GetPayoutRequestByTransactionReferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayoutRequest *models.PayoutRequest `protobuf:"bytes,1,opt,name=payout_request,json=payoutRequest,proto3" json:"payout_request,omitempty"`
}

func (x *GetPayoutRequestByTransactionReferenceResponse) Reset() {
	*x = GetPayoutRequestByTransactionReferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPayoutRequestByTransactionReferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPayoutRequestByTransactionReferenceResponse) ProtoMessage() {}

func (x *GetPayoutRequestByTransactionReferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPayoutRequestByTransactionReferenceResponse.ProtoReflect.Descriptor instead.
func (*GetPayoutRequestByTransactionReferenceResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{6}
}

func (x *GetPayoutRequestByTransactionReferenceResponse) GetPayoutRequest() *models.PayoutRequest {
	if x != nil {
		return x.PayoutRequest
	}
	return nil
}

type ProcessCMRBankAccountWithdrawalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *models.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *ProcessCMRBankAccountWithdrawalRequest) Reset() {
	*x = ProcessCMRBankAccountWithdrawalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessCMRBankAccountWithdrawalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessCMRBankAccountWithdrawalRequest) ProtoMessage() {}

func (x *ProcessCMRBankAccountWithdrawalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessCMRBankAccountWithdrawalRequest.ProtoReflect.Descriptor instead.
func (*ProcessCMRBankAccountWithdrawalRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{7}
}

func (x *ProcessCMRBankAccountWithdrawalRequest) GetTransaction() *models.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type ProcessCMRBankAccountWithdrawalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayoutRequest *models.PayoutRequest `protobuf:"bytes,1,opt,name=payout_request,json=payoutRequest,proto3" json:"payout_request,omitempty"`
}

func (x *ProcessCMRBankAccountWithdrawalResponse) Reset() {
	*x = ProcessCMRBankAccountWithdrawalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_payout_payout_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessCMRBankAccountWithdrawalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessCMRBankAccountWithdrawalResponse) ProtoMessage() {}

func (x *ProcessCMRBankAccountWithdrawalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_payout_payout_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessCMRBankAccountWithdrawalResponse.ProtoReflect.Descriptor instead.
func (*ProcessCMRBankAccountWithdrawalResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_payout_payout_proto_rawDescGZIP(), []int{8}
}

func (x *ProcessCMRBankAccountWithdrawalResponse) GetPayoutRequest() *models.PayoutRequest {
	if x != nil {
		return x.PayoutRequest
	}
	return nil
}

var File_kappa_services_payout_payout_proto protoreflect.FileDescriptor

var file_kappa_services_payout_payout_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x1a, 0x25, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x73, 0x69, 0x6c, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x23, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x6c,
	0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2e, 0x53, 0x69, 0x6c, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x59, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61,
	0x2e, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d,
	0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5a, 0x0a,
	0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0e,
	0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x50, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7e, 0x0a, 0x20, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x60, 0x0a, 0x21, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x50,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x70, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x2d,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64,
	0x22, 0x6d, 0x0a, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x61, 0x70,
	0x70, 0x61, 0x2e, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x0d, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x5e, 0x0a, 0x26, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x4d, 0x52, 0x42, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x66, 0x0a, 0x27, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x4d, 0x52, 0x42, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x50, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xe5, 0x04, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x2e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x97, 0x01, 0x0a, 0x26, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x35, 0x2e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x70, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x28, 0x2e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x69, 0x6c, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x69, 0x6c, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x82, 0x01, 0x0a, 0x1f, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x4d, 0x52, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x2e,
	0x2e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x4d, 0x52, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x70, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x4d, 0x52, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6c,
	0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_services_payout_payout_proto_rawDescOnce sync.Once
	file_kappa_services_payout_payout_proto_rawDescData = file_kappa_services_payout_payout_proto_rawDesc
)

func file_kappa_services_payout_payout_proto_rawDescGZIP() []byte {
	file_kappa_services_payout_payout_proto_rawDescOnce.Do(func() {
		file_kappa_services_payout_payout_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_services_payout_payout_proto_rawDescData)
	})
	return file_kappa_services_payout_payout_proto_rawDescData
}

var file_kappa_services_payout_payout_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_kappa_services_payout_payout_proto_goTypes = []interface{}{
	(*ProcessSilaTransactionUpdateRequest)(nil),            // 0: payout.ProcessSilaTransactionUpdateRequest
	(*CreatePayoutRequestRequest)(nil),                     // 1: payout.CreatePayoutRequestRequest
	(*CreatePayoutRequestResponse)(nil),                    // 2: payout.CreatePayoutRequestResponse
	(*UpdatePayoutRequestStatusRequest)(nil),               // 3: payout.UpdatePayoutRequestStatusRequest
	(*UpdatePayoutRequestStatusResponse)(nil),              // 4: payout.UpdatePayoutRequestStatusResponse
	(*GetPayoutRequestByTransactionReferenceRequest)(nil),  // 5: payout.GetPayoutRequestByTransactionReferenceRequest
	(*GetPayoutRequestByTransactionReferenceResponse)(nil), // 6: payout.GetPayoutRequestByTransactionReferenceResponse
	(*ProcessCMRBankAccountWithdrawalRequest)(nil),         // 7: payout.ProcessCMRBankAccountWithdrawalRequest
	(*ProcessCMRBankAccountWithdrawalResponse)(nil),        // 8: payout.ProcessCMRBankAccountWithdrawalResponse
	(*models.SilaTransactionUpdateEvent)(nil),              // 9: kappa.SilaTransactionUpdateEvent
	(*models.PayoutRequest)(nil),                           // 10: kappa.PayoutRequest
	(models.TransactionPayoutStatus)(0),                    // 11: kappa.TransactionPayoutStatus
	(*models.Transaction)(nil),                             // 12: kappa.Transaction
	(*emptypb.Empty)(nil),                                  // 13: google.protobuf.Empty
}
var file_kappa_services_payout_payout_proto_depIdxs = []int32{
	9,  // 0: payout.ProcessSilaTransactionUpdateRequest.event:type_name -> kappa.SilaTransactionUpdateEvent
	10, // 1: payout.CreatePayoutRequestRequest.payout_request:type_name -> kappa.PayoutRequest
	10, // 2: payout.CreatePayoutRequestResponse.payout_request:type_name -> kappa.PayoutRequest
	11, // 3: payout.UpdatePayoutRequestStatusRequest.status:type_name -> kappa.TransactionPayoutStatus
	10, // 4: payout.UpdatePayoutRequestStatusResponse.payout_request:type_name -> kappa.PayoutRequest
	10, // 5: payout.GetPayoutRequestByTransactionReferenceResponse.payout_request:type_name -> kappa.PayoutRequest
	12, // 6: payout.ProcessCMRBankAccountWithdrawalRequest.transaction:type_name -> kappa.Transaction
	10, // 7: payout.ProcessCMRBankAccountWithdrawalResponse.payout_request:type_name -> kappa.PayoutRequest
	1,  // 8: payout.PayoutService.CreatePayoutRequest:input_type -> payout.CreatePayoutRequestRequest
	5,  // 9: payout.PayoutService.GetPayoutRequestByTransactionReference:input_type -> payout.GetPayoutRequestByTransactionReferenceRequest
	3,  // 10: payout.PayoutService.UpdatePayoutRequestStatus:input_type -> payout.UpdatePayoutRequestStatusRequest
	0,  // 11: payout.PayoutService.ProcessSilaTransactionUpdate:input_type -> payout.ProcessSilaTransactionUpdateRequest
	7,  // 12: payout.PayoutService.ProcessCMRBankAccountWithdrawal:input_type -> payout.ProcessCMRBankAccountWithdrawalRequest
	2,  // 13: payout.PayoutService.CreatePayoutRequest:output_type -> payout.CreatePayoutRequestResponse
	6,  // 14: payout.PayoutService.GetPayoutRequestByTransactionReference:output_type -> payout.GetPayoutRequestByTransactionReferenceResponse
	4,  // 15: payout.PayoutService.UpdatePayoutRequestStatus:output_type -> payout.UpdatePayoutRequestStatusResponse
	13, // 16: payout.PayoutService.ProcessSilaTransactionUpdate:output_type -> google.protobuf.Empty
	8,  // 17: payout.PayoutService.ProcessCMRBankAccountWithdrawal:output_type -> payout.ProcessCMRBankAccountWithdrawalResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_kappa_services_payout_payout_proto_init() }
func file_kappa_services_payout_payout_proto_init() {
	if File_kappa_services_payout_payout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kappa_services_payout_payout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_kappa_services_payout_payout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePayoutRequestRequest); i {
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
		file_kappa_services_payout_payout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePayoutRequestResponse); i {
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
		file_kappa_services_payout_payout_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePayoutRequestStatusRequest); i {
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
		file_kappa_services_payout_payout_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePayoutRequestStatusResponse); i {
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
		file_kappa_services_payout_payout_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPayoutRequestByTransactionReferenceRequest); i {
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
		file_kappa_services_payout_payout_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPayoutRequestByTransactionReferenceResponse); i {
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
		file_kappa_services_payout_payout_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessCMRBankAccountWithdrawalRequest); i {
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
		file_kappa_services_payout_payout_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessCMRBankAccountWithdrawalResponse); i {
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
			RawDescriptor: file_kappa_services_payout_payout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kappa_services_payout_payout_proto_goTypes,
		DependencyIndexes: file_kappa_services_payout_payout_proto_depIdxs,
		MessageInfos:      file_kappa_services_payout_payout_proto_msgTypes,
	}.Build()
	File_kappa_services_payout_payout_proto = out.File
	file_kappa_services_payout_payout_proto_rawDesc = nil
	file_kappa_services_payout_payout_proto_goTypes = nil
	file_kappa_services_payout_payout_proto_depIdxs = nil
}

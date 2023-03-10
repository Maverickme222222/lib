// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: kappa/services/phones/phones.proto

package phones

import (
	models "github.com/kappapay/backend/lib/golang/src/kappa/common/models"
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

type CreatePhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone *models.Phone `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CreatePhoneNumberRequest) Reset() {
	*x = CreatePhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhoneNumberRequest) ProtoMessage() {}

func (x *CreatePhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*CreatePhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePhoneNumberRequest) GetPhone() *models.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

type CreatePhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone *models.Phone `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CreatePhoneNumberResponse) Reset() {
	*x = CreatePhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhoneNumberResponse) ProtoMessage() {}

func (x *CreatePhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*CreatePhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePhoneNumberResponse) GetPhone() *models.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

type GetPhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *GetPhoneNumberRequest) Reset() {
	*x = GetPhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhoneNumberRequest) ProtoMessage() {}

func (x *GetPhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*GetPhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{2}
}

func (x *GetPhoneNumberRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type GetPhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone *models.Phone `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetPhoneNumberResponse) Reset() {
	*x = GetPhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhoneNumberResponse) ProtoMessage() {}

func (x *GetPhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*GetPhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{3}
}

func (x *GetPhoneNumberResponse) GetPhone() *models.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

type GetPhoneNumberByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPhoneNumberByIdRequest) Reset() {
	*x = GetPhoneNumberByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhoneNumberByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhoneNumberByIdRequest) ProtoMessage() {}

func (x *GetPhoneNumberByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhoneNumberByIdRequest.ProtoReflect.Descriptor instead.
func (*GetPhoneNumberByIdRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{4}
}

func (x *GetPhoneNumberByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPhoneNumberByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone *models.Phone `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetPhoneNumberByIdResponse) Reset() {
	*x = GetPhoneNumberByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhoneNumberByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhoneNumberByIdResponse) ProtoMessage() {}

func (x *GetPhoneNumberByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhoneNumberByIdResponse.ProtoReflect.Descriptor instead.
func (*GetPhoneNumberByIdResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{5}
}

func (x *GetPhoneNumberByIdResponse) GetPhone() *models.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

type UpdatePhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone *models.Phone `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UpdatePhoneNumberRequest) Reset() {
	*x = UpdatePhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePhoneNumberRequest) ProtoMessage() {}

func (x *UpdatePhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*UpdatePhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePhoneNumberRequest) GetPhone() *models.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

type UpdatePhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone *models.Phone `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UpdatePhoneNumberResponse) Reset() {
	*x = UpdatePhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePhoneNumberResponse) ProtoMessage() {}

func (x *UpdatePhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*UpdatePhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePhoneNumberResponse) GetPhone() *models.Phone {
	if x != nil {
		return x.Phone
	}
	return nil
}

type DeletePhoneNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *DeletePhoneNumberRequest) Reset() {
	*x = DeletePhoneNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePhoneNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePhoneNumberRequest) ProtoMessage() {}

func (x *DeletePhoneNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePhoneNumberRequest.ProtoReflect.Descriptor instead.
func (*DeletePhoneNumberRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePhoneNumberRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type DeletePhoneNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *DeletePhoneNumberResponse) Reset() {
	*x = DeletePhoneNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePhoneNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePhoneNumberResponse) ProtoMessage() {}

func (x *DeletePhoneNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePhoneNumberResponse.ProtoReflect.Descriptor instead.
func (*DeletePhoneNumberResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{9}
}

func (x *DeletePhoneNumberResponse) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type DeletePhoneNumberByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePhoneNumberByIDRequest) Reset() {
	*x = DeletePhoneNumberByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePhoneNumberByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePhoneNumberByIDRequest) ProtoMessage() {}

func (x *DeletePhoneNumberByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePhoneNumberByIDRequest.ProtoReflect.Descriptor instead.
func (*DeletePhoneNumberByIDRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{10}
}

func (x *DeletePhoneNumberByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePhoneNumberByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *DeletePhoneNumberByIDResponse) Reset() {
	*x = DeletePhoneNumberByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePhoneNumberByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePhoneNumberByIDResponse) ProtoMessage() {}

func (x *DeletePhoneNumberByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePhoneNumberByIDResponse.ProtoReflect.Descriptor instead.
func (*DeletePhoneNumberByIDResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{11}
}

func (x *DeletePhoneNumberByIDResponse) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type MarkPhoneVerifiedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VendorHandle string `protobuf:"bytes,1,opt,name=vendor_handle,json=vendorHandle,proto3" json:"vendor_handle,omitempty"`
}

func (x *MarkPhoneVerifiedRequest) Reset() {
	*x = MarkPhoneVerifiedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkPhoneVerifiedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkPhoneVerifiedRequest) ProtoMessage() {}

func (x *MarkPhoneVerifiedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkPhoneVerifiedRequest.ProtoReflect.Descriptor instead.
func (*MarkPhoneVerifiedRequest) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{12}
}

func (x *MarkPhoneVerifiedRequest) GetVendorHandle() string {
	if x != nil {
		return x.VendorHandle
	}
	return ""
}

type MarkPhoneVerifiedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkPhoneVerifiedResponse) Reset() {
	*x = MarkPhoneVerifiedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_services_phones_phones_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkPhoneVerifiedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkPhoneVerifiedResponse) ProtoMessage() {}

func (x *MarkPhoneVerifiedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_services_phones_phones_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkPhoneVerifiedResponse.ProtoReflect.Descriptor instead.
func (*MarkPhoneVerifiedResponse) Descriptor() ([]byte, []int) {
	return file_kappa_services_phones_phones_proto_rawDescGZIP(), []int{13}
}

var File_kappa_services_phones_phones_proto protoreflect.FileDescriptor

var file_kappa_services_phones_phones_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x1a, 0x1f, 0x6b, 0x61,
	0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x61, 0x70, 0x70,
	0x61, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3f,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x61, 0x70,
	0x70, 0x61, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0x3a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x2b, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3e, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3f, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3d, 0x0a, 0x18, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x2e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x5a, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x18,
	0x4d, 0x61, 0x72, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x1b, 0x0a,
	0x19, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf3, 0x05, 0x0a, 0x0c, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x20, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x21, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20,
	0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x66, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b,
	0x61, 0x70, 0x70, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kappa_services_phones_phones_proto_rawDescOnce sync.Once
	file_kappa_services_phones_phones_proto_rawDescData = file_kappa_services_phones_phones_proto_rawDesc
)

func file_kappa_services_phones_phones_proto_rawDescGZIP() []byte {
	file_kappa_services_phones_phones_proto_rawDescOnce.Do(func() {
		file_kappa_services_phones_phones_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_services_phones_phones_proto_rawDescData)
	})
	return file_kappa_services_phones_phones_proto_rawDescData
}

var file_kappa_services_phones_phones_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_kappa_services_phones_phones_proto_goTypes = []interface{}{
	(*CreatePhoneNumberRequest)(nil),      // 0: phones.CreatePhoneNumberRequest
	(*CreatePhoneNumberResponse)(nil),     // 1: phones.CreatePhoneNumberResponse
	(*GetPhoneNumberRequest)(nil),         // 2: phones.GetPhoneNumberRequest
	(*GetPhoneNumberResponse)(nil),        // 3: phones.GetPhoneNumberResponse
	(*GetPhoneNumberByIdRequest)(nil),     // 4: phones.GetPhoneNumberByIdRequest
	(*GetPhoneNumberByIdResponse)(nil),    // 5: phones.GetPhoneNumberByIdResponse
	(*UpdatePhoneNumberRequest)(nil),      // 6: phones.UpdatePhoneNumberRequest
	(*UpdatePhoneNumberResponse)(nil),     // 7: phones.UpdatePhoneNumberResponse
	(*DeletePhoneNumberRequest)(nil),      // 8: phones.DeletePhoneNumberRequest
	(*DeletePhoneNumberResponse)(nil),     // 9: phones.DeletePhoneNumberResponse
	(*DeletePhoneNumberByIDRequest)(nil),  // 10: phones.DeletePhoneNumberByIDRequest
	(*DeletePhoneNumberByIDResponse)(nil), // 11: phones.DeletePhoneNumberByIDResponse
	(*MarkPhoneVerifiedRequest)(nil),      // 12: phones.MarkPhoneVerifiedRequest
	(*MarkPhoneVerifiedResponse)(nil),     // 13: phones.MarkPhoneVerifiedResponse
	(*models.Phone)(nil),                  // 14: kappa.Phone
	(*timestamppb.Timestamp)(nil),         // 15: google.protobuf.Timestamp
}
var file_kappa_services_phones_phones_proto_depIdxs = []int32{
	14, // 0: phones.CreatePhoneNumberRequest.phone:type_name -> kappa.Phone
	14, // 1: phones.CreatePhoneNumberResponse.phone:type_name -> kappa.Phone
	14, // 2: phones.GetPhoneNumberResponse.phone:type_name -> kappa.Phone
	14, // 3: phones.GetPhoneNumberByIdResponse.phone:type_name -> kappa.Phone
	14, // 4: phones.UpdatePhoneNumberRequest.phone:type_name -> kappa.Phone
	14, // 5: phones.UpdatePhoneNumberResponse.phone:type_name -> kappa.Phone
	15, // 6: phones.DeletePhoneNumberResponse.deleted_at:type_name -> google.protobuf.Timestamp
	15, // 7: phones.DeletePhoneNumberByIDResponse.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 8: phones.PhoneService.CreatePhoneNumber:input_type -> phones.CreatePhoneNumberRequest
	2,  // 9: phones.PhoneService.GetPhoneNumber:input_type -> phones.GetPhoneNumberRequest
	4,  // 10: phones.PhoneService.GetPhoneNumberByID:input_type -> phones.GetPhoneNumberByIdRequest
	6,  // 11: phones.PhoneService.UpdatePhoneNumber:input_type -> phones.UpdatePhoneNumberRequest
	8,  // 12: phones.PhoneService.DeletePhoneNumber:input_type -> phones.DeletePhoneNumberRequest
	10, // 13: phones.PhoneService.DeletePhoneNumberByID:input_type -> phones.DeletePhoneNumberByIDRequest
	2,  // 14: phones.PhoneService.GetOrCreatePhoneNumber:input_type -> phones.GetPhoneNumberRequest
	12, // 15: phones.PhoneService.MarkPhoneVerified:input_type -> phones.MarkPhoneVerifiedRequest
	1,  // 16: phones.PhoneService.CreatePhoneNumber:output_type -> phones.CreatePhoneNumberResponse
	3,  // 17: phones.PhoneService.GetPhoneNumber:output_type -> phones.GetPhoneNumberResponse
	5,  // 18: phones.PhoneService.GetPhoneNumberByID:output_type -> phones.GetPhoneNumberByIdResponse
	7,  // 19: phones.PhoneService.UpdatePhoneNumber:output_type -> phones.UpdatePhoneNumberResponse
	9,  // 20: phones.PhoneService.DeletePhoneNumber:output_type -> phones.DeletePhoneNumberResponse
	11, // 21: phones.PhoneService.DeletePhoneNumberByID:output_type -> phones.DeletePhoneNumberByIDResponse
	3,  // 22: phones.PhoneService.GetOrCreatePhoneNumber:output_type -> phones.GetPhoneNumberResponse
	13, // 23: phones.PhoneService.MarkPhoneVerified:output_type -> phones.MarkPhoneVerifiedResponse
	16, // [16:24] is the sub-list for method output_type
	8,  // [8:16] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_kappa_services_phones_phones_proto_init() }
func file_kappa_services_phones_phones_proto_init() {
	if File_kappa_services_phones_phones_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kappa_services_phones_phones_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhoneNumberRequest); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhoneNumberResponse); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhoneNumberRequest); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhoneNumberResponse); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhoneNumberByIdRequest); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhoneNumberByIdResponse); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePhoneNumberRequest); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePhoneNumberResponse); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePhoneNumberRequest); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePhoneNumberResponse); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePhoneNumberByIDRequest); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePhoneNumberByIDResponse); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkPhoneVerifiedRequest); i {
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
		file_kappa_services_phones_phones_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkPhoneVerifiedResponse); i {
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
			RawDescriptor: file_kappa_services_phones_phones_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kappa_services_phones_phones_proto_goTypes,
		DependencyIndexes: file_kappa_services_phones_phones_proto_depIdxs,
		MessageInfos:      file_kappa_services_phones_phones_proto_msgTypes,
	}.Build()
	File_kappa_services_phones_phones_proto = out.File
	file_kappa_services_phones_phones_proto_rawDesc = nil
	file_kappa_services_phones_phones_proto_goTypes = nil
	file_kappa_services_phones_phones_proto_depIdxs = nil
}

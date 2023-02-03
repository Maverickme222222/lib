// ------------------------------------------------------------------------------
// <copyright company="Kappa Pay Inc.">
// Copyright (C) Kappa Pay Inc.  All rights reserved.
// Unauthorised copying of this file, via any medium, is strictly prohibited.
// Proprietary and confidential
// </copyright>
// <author>Shashank A.</author>
// ------------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: kappa/common/models/time.proto

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

type HolidayCalendar int32

const (
	HolidayCalendar_HOLIDAY_CALENDAR_UNSPECIFIED  HolidayCalendar = 0
	HolidayCalendar_HOLIDAY_CALENDAR_US_FEDERAL   HolidayCalendar = 1
	HolidayCalendar_HOLIDAY_CALENDAR_CMR_NATIONAL HolidayCalendar = 2
)

// Enum value maps for HolidayCalendar.
var (
	HolidayCalendar_name = map[int32]string{
		0: "HOLIDAY_CALENDAR_UNSPECIFIED",
		1: "HOLIDAY_CALENDAR_US_FEDERAL",
		2: "HOLIDAY_CALENDAR_CMR_NATIONAL",
	}
	HolidayCalendar_value = map[string]int32{
		"HOLIDAY_CALENDAR_UNSPECIFIED":  0,
		"HOLIDAY_CALENDAR_US_FEDERAL":   1,
		"HOLIDAY_CALENDAR_CMR_NATIONAL": 2,
	}
)

func (x HolidayCalendar) Enum() *HolidayCalendar {
	p := new(HolidayCalendar)
	*p = x
	return p
}

func (x HolidayCalendar) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HolidayCalendar) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_time_proto_enumTypes[0].Descriptor()
}

func (HolidayCalendar) Type() protoreflect.EnumType {
	return &file_kappa_common_models_time_proto_enumTypes[0]
}

func (x HolidayCalendar) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HolidayCalendar.Descriptor instead.
func (HolidayCalendar) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_time_proto_rawDescGZIP(), []int{0}
}

type DayOfWeek int32

const (
	DayOfWeek_DAY_OF_WEEK_UNSPECIFIED DayOfWeek = 0
	DayOfWeek_DAY_OF_WEEK_MONDAY      DayOfWeek = 1
	DayOfWeek_DAY_OF_WEEK_TUESDAY     DayOfWeek = 2
	DayOfWeek_DAY_OF_WEEK_WEDNESDAY   DayOfWeek = 3
	DayOfWeek_DAY_OF_WEEK_THURSDAY    DayOfWeek = 4
	DayOfWeek_DAY_OF_WEEK_FRIDAY      DayOfWeek = 5
	DayOfWeek_DAY_OF_WEEK_SATURDAY    DayOfWeek = 6
	DayOfWeek_DAY_OF_WEEK_SUNDAY      DayOfWeek = 7
)

// Enum value maps for DayOfWeek.
var (
	DayOfWeek_name = map[int32]string{
		0: "DAY_OF_WEEK_UNSPECIFIED",
		1: "DAY_OF_WEEK_MONDAY",
		2: "DAY_OF_WEEK_TUESDAY",
		3: "DAY_OF_WEEK_WEDNESDAY",
		4: "DAY_OF_WEEK_THURSDAY",
		5: "DAY_OF_WEEK_FRIDAY",
		6: "DAY_OF_WEEK_SATURDAY",
		7: "DAY_OF_WEEK_SUNDAY",
	}
	DayOfWeek_value = map[string]int32{
		"DAY_OF_WEEK_UNSPECIFIED": 0,
		"DAY_OF_WEEK_MONDAY":      1,
		"DAY_OF_WEEK_TUESDAY":     2,
		"DAY_OF_WEEK_WEDNESDAY":   3,
		"DAY_OF_WEEK_THURSDAY":    4,
		"DAY_OF_WEEK_FRIDAY":      5,
		"DAY_OF_WEEK_SATURDAY":    6,
		"DAY_OF_WEEK_SUNDAY":      7,
	}
)

func (x DayOfWeek) Enum() *DayOfWeek {
	p := new(DayOfWeek)
	*p = x
	return p
}

func (x DayOfWeek) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DayOfWeek) Descriptor() protoreflect.EnumDescriptor {
	return file_kappa_common_models_time_proto_enumTypes[1].Descriptor()
}

func (DayOfWeek) Type() protoreflect.EnumType {
	return &file_kappa_common_models_time_proto_enumTypes[1]
}

func (x DayOfWeek) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DayOfWeek.Descriptor instead.
func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return file_kappa_common_models_time_proto_rawDescGZIP(), []int{1}
}

type TimeOfDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hours   int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds int32 `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,4,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *TimeOfDay) Reset() {
	*x = TimeOfDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeOfDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeOfDay) ProtoMessage() {}

func (x *TimeOfDay) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeOfDay.ProtoReflect.Descriptor instead.
func (*TimeOfDay) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_time_proto_rawDescGZIP(), []int{0}
}

func (x *TimeOfDay) GetHours() int32 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *TimeOfDay) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *TimeOfDay) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *TimeOfDay) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type DayTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DayOfWeek DayOfWeek  `protobuf:"varint,1,opt,name=day_of_week,json=dayOfWeek,proto3,enum=kappa.DayOfWeek" json:"day_of_week,omitempty"`
	TimeOfDay *TimeOfDay `protobuf:"bytes,2,opt,name=time_of_day,json=timeOfDay,proto3" json:"time_of_day,omitempty"`
}

func (x *DayTime) Reset() {
	*x = DayTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kappa_common_models_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DayTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DayTime) ProtoMessage() {}

func (x *DayTime) ProtoReflect() protoreflect.Message {
	mi := &file_kappa_common_models_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DayTime.ProtoReflect.Descriptor instead.
func (*DayTime) Descriptor() ([]byte, []int) {
	return file_kappa_common_models_time_proto_rawDescGZIP(), []int{1}
}

func (x *DayTime) GetDayOfWeek() DayOfWeek {
	if x != nil {
		return x.DayOfWeek
	}
	return DayOfWeek_DAY_OF_WEEK_UNSPECIFIED
}

func (x *DayTime) GetTimeOfDay() *TimeOfDay {
	if x != nil {
		return x.TimeOfDay
	}
	return nil
}

var File_kappa_common_models_time_proto protoreflect.FileDescriptor

var file_kappa_common_models_time_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x22, 0x6b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x4f,
	0x66, 0x44, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x22, 0x6d, 0x0a, 0x07, 0x44, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x0b, 0x64, 0x61, 0x79, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x44, 0x61, 0x79,
	0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x09, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65,
	0x6b, 0x12, 0x30, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x64, 0x61, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66,
	0x44, 0x61, 0x79, 0x2a, 0x77, 0x0a, 0x0f, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x1c, 0x48, 0x4f, 0x4c, 0x49, 0x44, 0x41,
	0x59, 0x5f, 0x43, 0x41, 0x4c, 0x45, 0x4e, 0x44, 0x41, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x48, 0x4f, 0x4c, 0x49,
	0x44, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x4c, 0x45, 0x4e, 0x44, 0x41, 0x52, 0x5f, 0x55, 0x53, 0x5f,
	0x46, 0x45, 0x44, 0x45, 0x52, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x48, 0x4f, 0x4c,
	0x49, 0x44, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x4c, 0x45, 0x4e, 0x44, 0x41, 0x52, 0x5f, 0x43, 0x4d,
	0x52, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x2a, 0xd8, 0x01, 0x0a,
	0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41,
	0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x41, 0x59, 0x5f, 0x4f,
	0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12,
	0x17, 0x0a, 0x13, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x54,
	0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x59, 0x5f,
	0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41,
	0x59, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45,
	0x45, 0x4b, 0x5f, 0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x16, 0x0a,
	0x12, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x46, 0x52, 0x49,
	0x44, 0x41, 0x59, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f,
	0x57, 0x45, 0x45, 0x4b, 0x5f, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59, 0x10, 0x06, 0x12,
	0x16, 0x0a, 0x12, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x5f, 0x53,
	0x55, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x07, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x70, 0x61, 0x79, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6b, 0x61, 0x70, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kappa_common_models_time_proto_rawDescOnce sync.Once
	file_kappa_common_models_time_proto_rawDescData = file_kappa_common_models_time_proto_rawDesc
)

func file_kappa_common_models_time_proto_rawDescGZIP() []byte {
	file_kappa_common_models_time_proto_rawDescOnce.Do(func() {
		file_kappa_common_models_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_kappa_common_models_time_proto_rawDescData)
	})
	return file_kappa_common_models_time_proto_rawDescData
}

var file_kappa_common_models_time_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_kappa_common_models_time_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kappa_common_models_time_proto_goTypes = []interface{}{
	(HolidayCalendar)(0), // 0: kappa.HolidayCalendar
	(DayOfWeek)(0),       // 1: kappa.DayOfWeek
	(*TimeOfDay)(nil),    // 2: kappa.TimeOfDay
	(*DayTime)(nil),      // 3: kappa.DayTime
}
var file_kappa_common_models_time_proto_depIdxs = []int32{
	1, // 0: kappa.DayTime.day_of_week:type_name -> kappa.DayOfWeek
	2, // 1: kappa.DayTime.time_of_day:type_name -> kappa.TimeOfDay
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kappa_common_models_time_proto_init() }
func file_kappa_common_models_time_proto_init() {
	if File_kappa_common_models_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kappa_common_models_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeOfDay); i {
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
		file_kappa_common_models_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DayTime); i {
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
			RawDescriptor: file_kappa_common_models_time_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kappa_common_models_time_proto_goTypes,
		DependencyIndexes: file_kappa_common_models_time_proto_depIdxs,
		EnumInfos:         file_kappa_common_models_time_proto_enumTypes,
		MessageInfos:      file_kappa_common_models_time_proto_msgTypes,
	}.Build()
	File_kappa_common_models_time_proto = out.File
	file_kappa_common_models_time_proto_rawDesc = nil
	file_kappa_common_models_time_proto_goTypes = nil
	file_kappa_common_models_time_proto_depIdxs = nil
}

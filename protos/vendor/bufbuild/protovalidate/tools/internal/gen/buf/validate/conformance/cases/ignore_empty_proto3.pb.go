// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/ignore_empty_proto3.proto

package cases

import (
	_ "github.com/bufbuild/protovalidate/tools/internal/gen/buf/validate"
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

type IgnoreEmptyProto3Scalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *IgnoreEmptyProto3Scalar) Reset() {
	*x = IgnoreEmptyProto3Scalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyProto3Scalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Scalar) ProtoMessage() {}

func (x *IgnoreEmptyProto3Scalar) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyProto3Scalar.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyProto3Scalar) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{0}
}

func (x *IgnoreEmptyProto3Scalar) GetVal() int32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type IgnoreEmptyProto3OptionalScalar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *int32 `protobuf:"varint,1,opt,name=val,proto3,oneof" json:"val,omitempty"`
}

func (x *IgnoreEmptyProto3OptionalScalar) Reset() {
	*x = IgnoreEmptyProto3OptionalScalar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyProto3OptionalScalar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3OptionalScalar) ProtoMessage() {}

func (x *IgnoreEmptyProto3OptionalScalar) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyProto3OptionalScalar.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyProto3OptionalScalar) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{1}
}

func (x *IgnoreEmptyProto3OptionalScalar) GetVal() int32 {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return 0
}

type IgnoreEmptyProto3Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *IgnoreEmptyProto3Message_Msg `protobuf:"bytes,1,opt,name=val,proto3,oneof" json:"val,omitempty"`
}

func (x *IgnoreEmptyProto3Message) Reset() {
	*x = IgnoreEmptyProto3Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyProto3Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Message) ProtoMessage() {}

func (x *IgnoreEmptyProto3Message) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyProto3Message.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyProto3Message) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{2}
}

func (x *IgnoreEmptyProto3Message) GetVal() *IgnoreEmptyProto3Message_Msg {
	if x != nil {
		return x.Val
	}
	return nil
}

type IgnoreEmptyProto3Oneof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to O:
	//
	//	*IgnoreEmptyProto3Oneof_Val
	O isIgnoreEmptyProto3Oneof_O `protobuf_oneof:"o"`
}

func (x *IgnoreEmptyProto3Oneof) Reset() {
	*x = IgnoreEmptyProto3Oneof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyProto3Oneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Oneof) ProtoMessage() {}

func (x *IgnoreEmptyProto3Oneof) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyProto3Oneof.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyProto3Oneof) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{3}
}

func (m *IgnoreEmptyProto3Oneof) GetO() isIgnoreEmptyProto3Oneof_O {
	if m != nil {
		return m.O
	}
	return nil
}

func (x *IgnoreEmptyProto3Oneof) GetVal() int32 {
	if x, ok := x.GetO().(*IgnoreEmptyProto3Oneof_Val); ok {
		return x.Val
	}
	return 0
}

type isIgnoreEmptyProto3Oneof_O interface {
	isIgnoreEmptyProto3Oneof_O()
}

type IgnoreEmptyProto3Oneof_Val struct {
	Val int32 `protobuf:"varint,1,opt,name=val,proto3,oneof"`
}

func (*IgnoreEmptyProto3Oneof_Val) isIgnoreEmptyProto3Oneof_O() {}

type IgnoreEmptyProto3Repeated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []int32 `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *IgnoreEmptyProto3Repeated) Reset() {
	*x = IgnoreEmptyProto3Repeated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyProto3Repeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Repeated) ProtoMessage() {}

func (x *IgnoreEmptyProto3Repeated) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyProto3Repeated.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyProto3Repeated) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{4}
}

func (x *IgnoreEmptyProto3Repeated) GetVal() []int32 {
	if x != nil {
		return x.Val
	}
	return nil
}

type IgnoreEmptyProto3Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[int32]int32 `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *IgnoreEmptyProto3Map) Reset() {
	*x = IgnoreEmptyProto3Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyProto3Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Map) ProtoMessage() {}

func (x *IgnoreEmptyProto3Map) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyProto3Map.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyProto3Map) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{5}
}

func (x *IgnoreEmptyProto3Map) GetVal() map[int32]int32 {
	if x != nil {
		return x.Val
	}
	return nil
}

type IgnoreEmptyRepeatedItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []int32 `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *IgnoreEmptyRepeatedItems) Reset() {
	*x = IgnoreEmptyRepeatedItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyRepeatedItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyRepeatedItems) ProtoMessage() {}

func (x *IgnoreEmptyRepeatedItems) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyRepeatedItems.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyRepeatedItems) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{6}
}

func (x *IgnoreEmptyRepeatedItems) GetVal() []int32 {
	if x != nil {
		return x.Val
	}
	return nil
}

type IgnoreEmptyMapPairs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[string]int32 `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *IgnoreEmptyMapPairs) Reset() {
	*x = IgnoreEmptyMapPairs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyMapPairs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyMapPairs) ProtoMessage() {}

func (x *IgnoreEmptyMapPairs) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyMapPairs.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyMapPairs) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{7}
}

func (x *IgnoreEmptyMapPairs) GetVal() map[string]int32 {
	if x != nil {
		return x.Val
	}
	return nil
}

type IgnoreEmptyProto3Message_Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val string `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *IgnoreEmptyProto3Message_Msg) Reset() {
	*x = IgnoreEmptyProto3Message_Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IgnoreEmptyProto3Message_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgnoreEmptyProto3Message_Msg) ProtoMessage() {}

func (x *IgnoreEmptyProto3Message_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgnoreEmptyProto3Message_Msg.ProtoReflect.Descriptor instead.
func (*IgnoreEmptyProto3Message_Msg) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP(), []int{2, 0}
}

func (x *IgnoreEmptyProto3Message_Msg) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_buf_validate_conformance_cases_ignore_empty_proto3_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDesc = []byte{
	0x0a, 0x38, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x17, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x53, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x12, 0x1c, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0a, 0xba, 0x48, 0x07, 0xd0, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0x4c, 0x0a, 0x1f, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x61,
	0x6c, 0x61, 0x72, 0x12, 0x21, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0xd0, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x76, 0x61, 0x6c, 0x22, 0xd4,
	0x01, 0x0a, 0x18, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x41, 0xba, 0x48, 0x3e, 0xba, 0x01, 0x38, 0x0a,
	0x1b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x06, 0x66, 0x6f,
	0x6f, 0x62, 0x61, 0x72, 0x1a, 0x11, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x76, 0x61, 0x6c, 0x20, 0x3d,
	0x3d, 0x20, 0x27, 0x66, 0x6f, 0x6f, 0x27, 0xd0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x88, 0x01, 0x01, 0x1a, 0x17, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x76, 0x61, 0x6c, 0x22, 0x3d, 0x0a, 0x16, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12,
	0x1e, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xba, 0x48,
	0x07, 0xd0, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x42,
	0x03, 0x0a, 0x01, 0x6f, 0x22, 0x3a, 0x0a, 0x19, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x0b,
	0xba, 0x48, 0x08, 0xd0, 0x01, 0x01, 0x92, 0x01, 0x02, 0x08, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0xac, 0x01, 0x0a, 0x14, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x61, 0x70, 0x12, 0x5c, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xd0, 0x01, 0x01, 0x9a, 0x01, 0x02,
	0x08, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x3d, 0x0a, 0x18, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0x92, 0x01, 0x09,
	0x22, 0x07, 0xd0, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0xb7,
	0x01, 0x0a, 0x13, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x61,
	0x70, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x68, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x61, 0x70, 0x50, 0x61, 0x69, 0x72, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x18, 0xba, 0x48, 0x15, 0x9a, 0x01, 0x12, 0x22, 0x07, 0xd0, 0x01, 0x01, 0x72, 0x02,
	0x10, 0x03, 0x2a, 0x07, 0xd0, 0x01, 0x01, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xad, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x42,
	0x16, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0xa2, 0x02,
	0x04, 0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescData = file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDesc
)

func file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDescData
}

var file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_buf_validate_conformance_cases_ignore_empty_proto3_proto_goTypes = []interface{}{
	(*IgnoreEmptyProto3Scalar)(nil),         // 0: buf.validate.conformance.cases.IgnoreEmptyProto3Scalar
	(*IgnoreEmptyProto3OptionalScalar)(nil), // 1: buf.validate.conformance.cases.IgnoreEmptyProto3OptionalScalar
	(*IgnoreEmptyProto3Message)(nil),        // 2: buf.validate.conformance.cases.IgnoreEmptyProto3Message
	(*IgnoreEmptyProto3Oneof)(nil),          // 3: buf.validate.conformance.cases.IgnoreEmptyProto3Oneof
	(*IgnoreEmptyProto3Repeated)(nil),       // 4: buf.validate.conformance.cases.IgnoreEmptyProto3Repeated
	(*IgnoreEmptyProto3Map)(nil),            // 5: buf.validate.conformance.cases.IgnoreEmptyProto3Map
	(*IgnoreEmptyRepeatedItems)(nil),        // 6: buf.validate.conformance.cases.IgnoreEmptyRepeatedItems
	(*IgnoreEmptyMapPairs)(nil),             // 7: buf.validate.conformance.cases.IgnoreEmptyMapPairs
	(*IgnoreEmptyProto3Message_Msg)(nil),    // 8: buf.validate.conformance.cases.IgnoreEmptyProto3Message.Msg
	nil,                                     // 9: buf.validate.conformance.cases.IgnoreEmptyProto3Map.ValEntry
	nil,                                     // 10: buf.validate.conformance.cases.IgnoreEmptyMapPairs.ValEntry
}
var file_buf_validate_conformance_cases_ignore_empty_proto3_proto_depIdxs = []int32{
	8,  // 0: buf.validate.conformance.cases.IgnoreEmptyProto3Message.val:type_name -> buf.validate.conformance.cases.IgnoreEmptyProto3Message.Msg
	9,  // 1: buf.validate.conformance.cases.IgnoreEmptyProto3Map.val:type_name -> buf.validate.conformance.cases.IgnoreEmptyProto3Map.ValEntry
	10, // 2: buf.validate.conformance.cases.IgnoreEmptyMapPairs.val:type_name -> buf.validate.conformance.cases.IgnoreEmptyMapPairs.ValEntry
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_ignore_empty_proto3_proto_init() }
func file_buf_validate_conformance_cases_ignore_empty_proto3_proto_init() {
	if File_buf_validate_conformance_cases_ignore_empty_proto3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyProto3Scalar); i {
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
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyProto3OptionalScalar); i {
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
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyProto3Message); i {
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
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyProto3Oneof); i {
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
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyProto3Repeated); i {
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
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyProto3Map); i {
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
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyRepeatedItems); i {
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
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyMapPairs); i {
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
		file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IgnoreEmptyProto3Message_Msg); i {
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
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*IgnoreEmptyProto3Oneof_Val)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_ignore_empty_proto3_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_ignore_empty_proto3_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_cases_ignore_empty_proto3_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_ignore_empty_proto3_proto = out.File
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_rawDesc = nil
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_goTypes = nil
	file_buf_validate_conformance_cases_ignore_empty_proto3_proto_depIdxs = nil
}

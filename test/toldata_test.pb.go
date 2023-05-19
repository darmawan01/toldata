// Copyright 2019 Citra Digital Lintas
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: toldata_test.proto

package test

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Id    int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TestARequest) Reset() {
	*x = TestARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toldata_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestARequest) ProtoMessage() {}

func (x *TestARequest) ProtoReflect() protoreflect.Message {
	mi := &file_toldata_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestARequest.ProtoReflect.Descriptor instead.
func (*TestARequest) Descriptor() ([]byte, []int) {
	return file_toldata_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestARequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TestARequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TestAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Id     int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TestAResponse) Reset() {
	*x = TestAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toldata_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAResponse) ProtoMessage() {}

func (x *TestAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_toldata_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAResponse.ProtoReflect.Descriptor instead.
func (*TestAResponse) Descriptor() ([]byte, []int) {
	return file_toldata_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestAResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *TestAResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FeedDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data int64 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FeedDataRequest) Reset() {
	*x = FeedDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toldata_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedDataRequest) ProtoMessage() {}

func (x *FeedDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_toldata_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedDataRequest.ProtoReflect.Descriptor instead.
func (*FeedDataRequest) Descriptor() ([]byte, []int) {
	return file_toldata_test_proto_rawDescGZIP(), []int{2}
}

func (x *FeedDataRequest) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

type FeedDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *FeedDataResponse) Reset() {
	*x = FeedDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toldata_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedDataResponse) ProtoMessage() {}

func (x *FeedDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_toldata_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedDataResponse.ProtoReflect.Descriptor instead.
func (*FeedDataResponse) Descriptor() ([]byte, []int) {
	return file_toldata_test_proto_rawDescGZIP(), []int{3}
}

func (x *FeedDataResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type StreamDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StreamDataRequest) Reset() {
	*x = StreamDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toldata_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamDataRequest) ProtoMessage() {}

func (x *StreamDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_toldata_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamDataRequest.ProtoReflect.Descriptor instead.
func (*StreamDataRequest) Descriptor() ([]byte, []int) {
	return file_toldata_test_proto_rawDescGZIP(), []int{4}
}

func (x *StreamDataRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type StreamDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data int64 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamDataResponse) Reset() {
	*x = StreamDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toldata_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamDataResponse) ProtoMessage() {}

func (x *StreamDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_toldata_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamDataResponse.ProtoReflect.Descriptor instead.
func (*StreamDataResponse) Descriptor() ([]byte, []int) {
	return file_toldata_test_proto_rawDescGZIP(), []int{5}
}

func (x *StreamDataResponse) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

type TestGetIPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *TestGetIPResponse) Reset() {
	*x = TestGetIPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toldata_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestGetIPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestGetIPResponse) ProtoMessage() {}

func (x *TestGetIPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_toldata_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestGetIPResponse.ProtoReflect.Descriptor instead.
func (*TestGetIPResponse) Descriptor() ([]byte, []int) {
	return file_toldata_test_proto_rawDescGZIP(), []int{6}
}

func (x *TestGetIPResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_toldata_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_toldata_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_toldata_test_proto_rawDescGZIP(), []int{7}
}

var file_toldata_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         999999,
		Name:          "cdl.toldatatest.rest_mount",
		Tag:           "bytes,999999,opt,name=rest_mount",
		Filename:      "toldata_test.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional string rest_mount = 999999;
	E_RestMount = &file_toldata_test_proto_extTypes[0]
)

var File_toldata_test_proto protoreflect.FileDescriptor

var file_toldata_test_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x41,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a,
	0x0d, 0x54, 0x65, 0x73, 0x74, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x46, 0x65, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a,
	0x10, 0x46, 0x65, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x73, 0x75, 0x6d, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x23, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x49, 0x50, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0xd4, 0x04, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x41, 0x12, 0x1d, 0x2e, 0x63,
	0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x64,
	0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x41, 0x42, 0x12, 0x1d, 0x2e, 0x63, 0x64, 0x6c,
	0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x64, 0x6c, 0x2e,
	0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x49, 0x50, 0x12, 0x16, 0x2e, 0x63, 0x64,
	0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x63, 0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x47, 0x65, 0x74, 0x49, 0x50, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x08, 0x46, 0x65, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x2e, 0x63, 0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x64, 0x6c, 0x2e, 0x74, 0x6f,
	0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x59,
	0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x2e, 0x63,
	0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x63, 0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5d, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x41, 0x6c, 0x74, 0x31, 0x12, 0x22, 0x2e, 0x63, 0x64,
	0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x63, 0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x63, 0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x63, 0x64, 0x6c, 0x2e, 0x74, 0x6f, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x1a, 0x0e, 0xfa, 0xa3, 0xe8, 0x03, 0x09, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x40, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbf, 0x84, 0x3d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b,
	0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_toldata_test_proto_rawDescOnce sync.Once
	file_toldata_test_proto_rawDescData = file_toldata_test_proto_rawDesc
)

func file_toldata_test_proto_rawDescGZIP() []byte {
	file_toldata_test_proto_rawDescOnce.Do(func() {
		file_toldata_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_toldata_test_proto_rawDescData)
	})
	return file_toldata_test_proto_rawDescData
}

var file_toldata_test_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_toldata_test_proto_goTypes = []interface{}{
	(*TestARequest)(nil),                // 0: cdl.toldatatest.TestARequest
	(*TestAResponse)(nil),               // 1: cdl.toldatatest.TestAResponse
	(*FeedDataRequest)(nil),             // 2: cdl.toldatatest.FeedDataRequest
	(*FeedDataResponse)(nil),            // 3: cdl.toldatatest.FeedDataResponse
	(*StreamDataRequest)(nil),           // 4: cdl.toldatatest.StreamDataRequest
	(*StreamDataResponse)(nil),          // 5: cdl.toldatatest.StreamDataResponse
	(*TestGetIPResponse)(nil),           // 6: cdl.toldatatest.TestGetIPResponse
	(*Empty)(nil),                       // 7: cdl.toldatatest.Empty
	(*descriptorpb.ServiceOptions)(nil), // 8: google.protobuf.ServiceOptions
}
var file_toldata_test_proto_depIdxs = []int32{
	8, // 0: cdl.toldatatest.rest_mount:extendee -> google.protobuf.ServiceOptions
	0, // 1: cdl.toldatatest.TestService.GetTestA:input_type -> cdl.toldatatest.TestARequest
	0, // 2: cdl.toldatatest.TestService.GetTestAB:input_type -> cdl.toldatatest.TestARequest
	7, // 3: cdl.toldatatest.TestService.GetTestGetIP:input_type -> cdl.toldatatest.Empty
	2, // 4: cdl.toldatatest.TestService.FeedData:input_type -> cdl.toldatatest.FeedDataRequest
	4, // 5: cdl.toldatatest.TestService.StreamData:input_type -> cdl.toldatatest.StreamDataRequest
	4, // 6: cdl.toldatatest.TestService.StreamDataAlt1:input_type -> cdl.toldatatest.StreamDataRequest
	7, // 7: cdl.toldatatest.TestService.TestEmpty:input_type -> cdl.toldatatest.Empty
	1, // 8: cdl.toldatatest.TestService.GetTestA:output_type -> cdl.toldatatest.TestAResponse
	1, // 9: cdl.toldatatest.TestService.GetTestAB:output_type -> cdl.toldatatest.TestAResponse
	6, // 10: cdl.toldatatest.TestService.GetTestGetIP:output_type -> cdl.toldatatest.TestGetIPResponse
	3, // 11: cdl.toldatatest.TestService.FeedData:output_type -> cdl.toldatatest.FeedDataResponse
	5, // 12: cdl.toldatatest.TestService.StreamData:output_type -> cdl.toldatatest.StreamDataResponse
	5, // 13: cdl.toldatatest.TestService.StreamDataAlt1:output_type -> cdl.toldatatest.StreamDataResponse
	7, // 14: cdl.toldatatest.TestService.TestEmpty:output_type -> cdl.toldatatest.Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_toldata_test_proto_init() }
func file_toldata_test_proto_init() {
	if File_toldata_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_toldata_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestARequest); i {
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
		file_toldata_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAResponse); i {
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
		file_toldata_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedDataRequest); i {
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
		file_toldata_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedDataResponse); i {
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
		file_toldata_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamDataRequest); i {
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
		file_toldata_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamDataResponse); i {
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
		file_toldata_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestGetIPResponse); i {
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
		file_toldata_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_toldata_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_toldata_test_proto_goTypes,
		DependencyIndexes: file_toldata_test_proto_depIdxs,
		MessageInfos:      file_toldata_test_proto_msgTypes,
		ExtensionInfos:    file_toldata_test_proto_extTypes,
	}.Build()
	File_toldata_test_proto = out.File
	file_toldata_test_proto_rawDesc = nil
	file_toldata_test_proto_goTypes = nil
	file_toldata_test_proto_depIdxs = nil
}

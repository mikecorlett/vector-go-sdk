// Copyright (c) 2018 Anki, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License in the file LICENSE.txt or at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Alexa messages

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.4
// source: alexa.proto

package vectorpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AlexaAuthState int32

const (
	// Invalid/error/versioning issue
	AlexaAuthState_ALEXA_AUTH_INVALID AlexaAuthState = 0
	// Not opted in, or opt-in attempted but failed
	AlexaAuthState_ALEXA_AUTH_UNINITIALIZED AlexaAuthState = 1
	// Opted in, and attempting to authorize
	AlexaAuthState_ALEXA_AUTH_REQUESTING_AUTH AlexaAuthState = 2
	// Opted in, and waiting on the user to enter a code
	AlexaAuthState_ALEXA_AUTH_WAITING_FOR_CODE AlexaAuthState = 3
	// Opted in, and authorized / in use
	AlexaAuthState_ALEXA_AUTH_AUTHORIZED AlexaAuthState = 4
)

// Enum value maps for AlexaAuthState.
var (
	AlexaAuthState_name = map[int32]string{
		0: "ALEXA_AUTH_INVALID",
		1: "ALEXA_AUTH_UNINITIALIZED",
		2: "ALEXA_AUTH_REQUESTING_AUTH",
		3: "ALEXA_AUTH_WAITING_FOR_CODE",
		4: "ALEXA_AUTH_AUTHORIZED",
	}
	AlexaAuthState_value = map[string]int32{
		"ALEXA_AUTH_INVALID":          0,
		"ALEXA_AUTH_UNINITIALIZED":    1,
		"ALEXA_AUTH_REQUESTING_AUTH":  2,
		"ALEXA_AUTH_WAITING_FOR_CODE": 3,
		"ALEXA_AUTH_AUTHORIZED":       4,
	}
)

func (x AlexaAuthState) Enum() *AlexaAuthState {
	p := new(AlexaAuthState)
	*p = x
	return p
}

func (x AlexaAuthState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlexaAuthState) Descriptor() protoreflect.EnumDescriptor {
	return file_alexa_proto_enumTypes[0].Descriptor()
}

func (AlexaAuthState) Type() protoreflect.EnumType {
	return &file_alexa_proto_enumTypes[0]
}

func (x AlexaAuthState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlexaAuthState.Descriptor instead.
func (AlexaAuthState) EnumDescriptor() ([]byte, []int) {
	return file_alexa_proto_rawDescGZIP(), []int{0}
}

type AlexaAuthStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AlexaAuthStateRequest) Reset() {
	*x = AlexaAuthStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlexaAuthStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlexaAuthStateRequest) ProtoMessage() {}

func (x *AlexaAuthStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alexa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlexaAuthStateRequest.ProtoReflect.Descriptor instead.
func (*AlexaAuthStateRequest) Descriptor() ([]byte, []int) {
	return file_alexa_proto_rawDescGZIP(), []int{0}
}

type AlexaAuthStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    *ResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	AuthState AlexaAuthState  `protobuf:"varint,2,opt,name=auth_state,json=authState,proto3,enum=Anki.Vector.external_interface.AlexaAuthState" json:"auth_state,omitempty"`
	Extra     string          `protobuf:"bytes,3,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *AlexaAuthStateResponse) Reset() {
	*x = AlexaAuthStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlexaAuthStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlexaAuthStateResponse) ProtoMessage() {}

func (x *AlexaAuthStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alexa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlexaAuthStateResponse.ProtoReflect.Descriptor instead.
func (*AlexaAuthStateResponse) Descriptor() ([]byte, []int) {
	return file_alexa_proto_rawDescGZIP(), []int{1}
}

func (x *AlexaAuthStateResponse) GetStatus() *ResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AlexaAuthStateResponse) GetAuthState() AlexaAuthState {
	if x != nil {
		return x.AuthState
	}
	return AlexaAuthState_ALEXA_AUTH_INVALID
}

func (x *AlexaAuthStateResponse) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type AlexaOptInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptIn bool `protobuf:"varint,1,opt,name=opt_in,json=optIn,proto3" json:"opt_in,omitempty"`
}

func (x *AlexaOptInRequest) Reset() {
	*x = AlexaOptInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlexaOptInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlexaOptInRequest) ProtoMessage() {}

func (x *AlexaOptInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alexa_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlexaOptInRequest.ProtoReflect.Descriptor instead.
func (*AlexaOptInRequest) Descriptor() ([]byte, []int) {
	return file_alexa_proto_rawDescGZIP(), []int{2}
}

func (x *AlexaOptInRequest) GetOptIn() bool {
	if x != nil {
		return x.OptIn
	}
	return false
}

type AlexaOptInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *ResponseStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AlexaOptInResponse) Reset() {
	*x = AlexaOptInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexa_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlexaOptInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlexaOptInResponse) ProtoMessage() {}

func (x *AlexaOptInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alexa_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlexaOptInResponse.ProtoReflect.Descriptor instead.
func (*AlexaOptInResponse) Descriptor() ([]byte, []int) {
	return file_alexa_proto_rawDescGZIP(), []int{3}
}

func (x *AlexaOptInResponse) GetStatus() *ResponseStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type AlexaAuthEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthState AlexaAuthState `protobuf:"varint,1,opt,name=auth_state,json=authState,proto3,enum=Anki.Vector.external_interface.AlexaAuthState" json:"auth_state,omitempty"`
	Extra     string         `protobuf:"bytes,2,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *AlexaAuthEvent) Reset() {
	*x = AlexaAuthEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alexa_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlexaAuthEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlexaAuthEvent) ProtoMessage() {}

func (x *AlexaAuthEvent) ProtoReflect() protoreflect.Message {
	mi := &file_alexa_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlexaAuthEvent.ProtoReflect.Descriptor instead.
func (*AlexaAuthEvent) Descriptor() ([]byte, []int) {
	return file_alexa_proto_rawDescGZIP(), []int{4}
}

func (x *AlexaAuthEvent) GetAuthState() AlexaAuthState {
	if x != nil {
		return x.AuthState
	}
	return AlexaAuthState_ALEXA_AUTH_INVALID
}

func (x *AlexaAuthEvent) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

var File_alexa_proto protoreflect.FileDescriptor

var file_alexa_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x41,
	0x6e, 0x6b, 0x69, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x1a, 0x15, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc5, 0x01,
	0x0a, 0x16, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x41, 0x6e, 0x6b, 0x69, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x4d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x41, 0x6e, 0x6b, 0x69, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x2a, 0x0a, 0x11, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x4f, 0x70,
	0x74, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x70,
	0x74, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6f, 0x70, 0x74, 0x49,
	0x6e, 0x22, 0x5c, 0x0a, 0x12, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x4f, 0x70, 0x74, 0x49, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x41, 0x6e, 0x6b, 0x69, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x75, 0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x41, 0x75, 0x74, 0x68, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x41, 0x6e, 0x6b, 0x69, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x6c, 0x65, 0x78, 0x61, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2a, 0xa2, 0x01, 0x0a, 0x0e, 0x41, 0x6c, 0x65, 0x78, 0x61,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x4c, 0x45,
	0x58, 0x41, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x4c, 0x45, 0x58, 0x41, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f,
	0x55, 0x4e, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x1e, 0x0a, 0x1a, 0x41, 0x4c, 0x45, 0x58, 0x41, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x10, 0x02, 0x12,
	0x1f, 0x0a, 0x1b, 0x41, 0x4c, 0x45, 0x58, 0x41, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x57, 0x41,
	0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x03,
	0x12, 0x19, 0x0a, 0x15, 0x41, 0x4c, 0x45, 0x58, 0x41, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x41,
	0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x04, 0x42, 0x3a, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x2d, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alexa_proto_rawDescOnce sync.Once
	file_alexa_proto_rawDescData = file_alexa_proto_rawDesc
)

func file_alexa_proto_rawDescGZIP() []byte {
	file_alexa_proto_rawDescOnce.Do(func() {
		file_alexa_proto_rawDescData = protoimpl.X.CompressGZIP(file_alexa_proto_rawDescData)
	})
	return file_alexa_proto_rawDescData
}

var file_alexa_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_alexa_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_alexa_proto_goTypes = []interface{}{
	(AlexaAuthState)(0),            // 0: Anki.Vector.external_interface.AlexaAuthState
	(*AlexaAuthStateRequest)(nil),  // 1: Anki.Vector.external_interface.AlexaAuthStateRequest
	(*AlexaAuthStateResponse)(nil), // 2: Anki.Vector.external_interface.AlexaAuthStateResponse
	(*AlexaOptInRequest)(nil),      // 3: Anki.Vector.external_interface.AlexaOptInRequest
	(*AlexaOptInResponse)(nil),     // 4: Anki.Vector.external_interface.AlexaOptInResponse
	(*AlexaAuthEvent)(nil),         // 5: Anki.Vector.external_interface.AlexaAuthEvent
	(*ResponseStatus)(nil),         // 6: Anki.Vector.external_interface.ResponseStatus
}
var file_alexa_proto_depIdxs = []int32{
	6, // 0: Anki.Vector.external_interface.AlexaAuthStateResponse.status:type_name -> Anki.Vector.external_interface.ResponseStatus
	0, // 1: Anki.Vector.external_interface.AlexaAuthStateResponse.auth_state:type_name -> Anki.Vector.external_interface.AlexaAuthState
	6, // 2: Anki.Vector.external_interface.AlexaOptInResponse.status:type_name -> Anki.Vector.external_interface.ResponseStatus
	0, // 3: Anki.Vector.external_interface.AlexaAuthEvent.auth_state:type_name -> Anki.Vector.external_interface.AlexaAuthState
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_alexa_proto_init() }
func file_alexa_proto_init() {
	if File_alexa_proto != nil {
		return
	}
	file_response_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alexa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlexaAuthStateRequest); i {
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
		file_alexa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlexaAuthStateResponse); i {
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
		file_alexa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlexaOptInRequest); i {
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
		file_alexa_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlexaOptInResponse); i {
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
		file_alexa_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlexaAuthEvent); i {
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
			RawDescriptor: file_alexa_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alexa_proto_goTypes,
		DependencyIndexes: file_alexa_proto_depIdxs,
		EnumInfos:         file_alexa_proto_enumTypes,
		MessageInfos:      file_alexa_proto_msgTypes,
	}.Build()
	File_alexa_proto = out.File
	file_alexa_proto_rawDesc = nil
	file_alexa_proto_goTypes = nil
	file_alexa_proto_depIdxs = nil
}

//
// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: bgp/bgp.proto

package bgp

import (
	_ "github.com/openconfig/gnoi/types"
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

type ClearBGPNeighborRequest_Mode int32

const (
	ClearBGPNeighborRequest_SOFT   ClearBGPNeighborRequest_Mode = 0 // Send route-refresh and reapply policy.
	ClearBGPNeighborRequest_SOFTIN ClearBGPNeighborRequest_Mode = 1 // Re-apply inbound policy on stored Adj-RIB-In.
	ClearBGPNeighborRequest_HARD   ClearBGPNeighborRequest_Mode = 2 // Teardown and restart TCP connection.
)

// Enum value maps for ClearBGPNeighborRequest_Mode.
var (
	ClearBGPNeighborRequest_Mode_name = map[int32]string{
		0: "SOFT",
		1: "SOFTIN",
		2: "HARD",
	}
	ClearBGPNeighborRequest_Mode_value = map[string]int32{
		"SOFT":   0,
		"SOFTIN": 1,
		"HARD":   2,
	}
)

func (x ClearBGPNeighborRequest_Mode) Enum() *ClearBGPNeighborRequest_Mode {
	p := new(ClearBGPNeighborRequest_Mode)
	*p = x
	return p
}

func (x ClearBGPNeighborRequest_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClearBGPNeighborRequest_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_bgp_bgp_proto_enumTypes[0].Descriptor()
}

func (ClearBGPNeighborRequest_Mode) Type() protoreflect.EnumType {
	return &file_bgp_bgp_proto_enumTypes[0]
}

func (x ClearBGPNeighborRequest_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClearBGPNeighborRequest_Mode.Descriptor instead.
func (ClearBGPNeighborRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return file_bgp_bgp_proto_rawDescGZIP(), []int{0, 0}
}

type ClearBGPNeighborRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // IPv4 or IPv6 BGP neighbor address to clear.
	// Routing instance containing the neighbor. Defaults to the global routing
	// table.
	RoutingInstance string                       `protobuf:"bytes,2,opt,name=routing_instance,json=routingInstance,proto3" json:"routing_instance,omitempty"`
	Mode            ClearBGPNeighborRequest_Mode `protobuf:"varint,3,opt,name=mode,proto3,enum=gnoi.bgp.ClearBGPNeighborRequest_Mode" json:"mode,omitempty"` // Which mode to use for the clear operation.
}

func (x *ClearBGPNeighborRequest) Reset() {
	*x = ClearBGPNeighborRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgp_bgp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearBGPNeighborRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearBGPNeighborRequest) ProtoMessage() {}

func (x *ClearBGPNeighborRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgp_bgp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearBGPNeighborRequest.ProtoReflect.Descriptor instead.
func (*ClearBGPNeighborRequest) Descriptor() ([]byte, []int) {
	return file_bgp_bgp_proto_rawDescGZIP(), []int{0}
}

func (x *ClearBGPNeighborRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ClearBGPNeighborRequest) GetRoutingInstance() string {
	if x != nil {
		return x.RoutingInstance
	}
	return ""
}

func (x *ClearBGPNeighborRequest) GetMode() ClearBGPNeighborRequest_Mode {
	if x != nil {
		return x.Mode
	}
	return ClearBGPNeighborRequest_SOFT
}

type ClearBGPNeighborResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearBGPNeighborResponse) Reset() {
	*x = ClearBGPNeighborResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgp_bgp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearBGPNeighborResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearBGPNeighborResponse) ProtoMessage() {}

func (x *ClearBGPNeighborResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgp_bgp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearBGPNeighborResponse.ProtoReflect.Descriptor instead.
func (*ClearBGPNeighborResponse) Descriptor() ([]byte, []int) {
	return file_bgp_bgp_proto_rawDescGZIP(), []int{1}
}

var File_bgp_bgp_proto protoreflect.FileDescriptor

var file_bgp_bgp_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x67, 0x70, 0x2f, 0x62, 0x67, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x62, 0x67, 0x70, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x17, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x42, 0x47, 0x50, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x62, 0x67,
	0x70, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x42, 0x47, 0x50, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x22, 0x26, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x53, 0x4f, 0x46, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x46, 0x54, 0x49, 0x4e,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x41, 0x52, 0x44, 0x10, 0x02, 0x22, 0x1a, 0x0a, 0x18,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x42, 0x47, 0x50, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x62, 0x0a, 0x03, 0x42, 0x47, 0x50, 0x12,
	0x5b, 0x0a, 0x10, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x42, 0x47, 0x50, 0x4e, 0x65, 0x69, 0x67, 0x68,
	0x62, 0x6f, 0x72, 0x12, 0x21, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x62, 0x67, 0x70, 0x2e, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x42, 0x47, 0x50, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x62, 0x67,
	0x70, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x42, 0x47, 0x50, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x1e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x62, 0x67, 0x70, 0xd2, 0x3e,
	0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bgp_bgp_proto_rawDescOnce sync.Once
	file_bgp_bgp_proto_rawDescData = file_bgp_bgp_proto_rawDesc
)

func file_bgp_bgp_proto_rawDescGZIP() []byte {
	file_bgp_bgp_proto_rawDescOnce.Do(func() {
		file_bgp_bgp_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgp_bgp_proto_rawDescData)
	})
	return file_bgp_bgp_proto_rawDescData
}

var file_bgp_bgp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bgp_bgp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bgp_bgp_proto_goTypes = []interface{}{
	(ClearBGPNeighborRequest_Mode)(0), // 0: gnoi.bgp.ClearBGPNeighborRequest.Mode
	(*ClearBGPNeighborRequest)(nil),   // 1: gnoi.bgp.ClearBGPNeighborRequest
	(*ClearBGPNeighborResponse)(nil),  // 2: gnoi.bgp.ClearBGPNeighborResponse
}
var file_bgp_bgp_proto_depIdxs = []int32{
	0, // 0: gnoi.bgp.ClearBGPNeighborRequest.mode:type_name -> gnoi.bgp.ClearBGPNeighborRequest.Mode
	1, // 1: gnoi.bgp.BGP.ClearBGPNeighbor:input_type -> gnoi.bgp.ClearBGPNeighborRequest
	2, // 2: gnoi.bgp.BGP.ClearBGPNeighbor:output_type -> gnoi.bgp.ClearBGPNeighborResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bgp_bgp_proto_init() }
func file_bgp_bgp_proto_init() {
	if File_bgp_bgp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgp_bgp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearBGPNeighborRequest); i {
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
		file_bgp_bgp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearBGPNeighborResponse); i {
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
			RawDescriptor: file_bgp_bgp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bgp_bgp_proto_goTypes,
		DependencyIndexes: file_bgp_bgp_proto_depIdxs,
		EnumInfos:         file_bgp_bgp_proto_enumTypes,
		MessageInfos:      file_bgp_bgp_proto_msgTypes,
	}.Build()
	File_bgp_bgp_proto = out.File
	file_bgp_bgp_proto_rawDesc = nil
	file_bgp_bgp_proto_goTypes = nil
	file_bgp_bgp_proto_depIdxs = nil
}

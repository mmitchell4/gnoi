//
// Copyright 2018 Google Inc. All Rights Reserved.
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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.9
// source: common/common.proto

package common

import (
	types "github.com/openconfig/gnoi/types"
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

type RemoteDownload_Protocol int32

const (
	RemoteDownload_UNKNOWN RemoteDownload_Protocol = 0
	RemoteDownload_SFTP    RemoteDownload_Protocol = 1
	RemoteDownload_HTTP    RemoteDownload_Protocol = 2
	RemoteDownload_HTTPS   RemoteDownload_Protocol = 3
	RemoteDownload_SCP     RemoteDownload_Protocol = 4
)

// Enum value maps for RemoteDownload_Protocol.
var (
	RemoteDownload_Protocol_name = map[int32]string{
		0: "UNKNOWN",
		1: "SFTP",
		2: "HTTP",
		3: "HTTPS",
		4: "SCP",
	}
	RemoteDownload_Protocol_value = map[string]int32{
		"UNKNOWN": 0,
		"SFTP":    1,
		"HTTP":    2,
		"HTTPS":   3,
		"SCP":     4,
	}
)

func (x RemoteDownload_Protocol) Enum() *RemoteDownload_Protocol {
	p := new(RemoteDownload_Protocol)
	*p = x
	return p
}

func (x RemoteDownload_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoteDownload_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[0].Descriptor()
}

func (RemoteDownload_Protocol) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[0]
}

func (x RemoteDownload_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RemoteDownload_Protocol.Descriptor instead.
func (RemoteDownload_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0, 0}
}

// RemoteDownload defines the details for a device to initiate a file transfer
// from or to a remote location.
type RemoteDownload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path information containing where to download the data from or to.
	// For HTTP(S), this will be the URL (i.e. foo.com/file.tbz2).
	// For SFTP and SCP, this will be the address:/path/to/file
	// (i.e. host.foo.com:/bar/baz).
	Path        string                  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Protocol    RemoteDownload_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=gnoi.common.RemoteDownload_Protocol" json:"protocol,omitempty"`
	Credentials *types.Credentials      `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// Optional source address used to initiate connections from the device.
	// It can be either an IPv4 address or an IPv6 address, depending on the
	// connection's destination address.
	SourceAddress string `protobuf:"bytes,4,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
}

func (x *RemoteDownload) Reset() {
	*x = RemoteDownload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteDownload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteDownload) ProtoMessage() {}

func (x *RemoteDownload) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteDownload.ProtoReflect.Descriptor instead.
func (*RemoteDownload) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteDownload) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RemoteDownload) GetProtocol() RemoteDownload_Protocol {
	if x != nil {
		return x.Protocol
	}
	return RemoteDownload_UNKNOWN
}

func (x *RemoteDownload) GetCredentials() *types.Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *RemoteDownload) GetSourceAddress() string {
	if x != nil {
		return x.SourceAddress
	}
	return ""
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x39,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x3f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x46, 0x54,
	0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x48, 0x54, 0x54, 0x50, 0x53, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x43, 0x50, 0x10,
	0x04, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_common_proto_goTypes = []interface{}{
	(RemoteDownload_Protocol)(0), // 0: gnoi.common.RemoteDownload.Protocol
	(*RemoteDownload)(nil),       // 1: gnoi.common.RemoteDownload
	(*types.Credentials)(nil),    // 2: gnoi.types.Credentials
}
var file_common_common_proto_depIdxs = []int32{
	0, // 0: gnoi.common.RemoteDownload.protocol:type_name -> gnoi.common.RemoteDownload.Protocol
	2, // 1: gnoi.common.RemoteDownload.credentials:type_name -> gnoi.types.Credentials
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteDownload); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		EnumInfos:         file_common_common_proto_enumTypes,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}

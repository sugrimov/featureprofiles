// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ocrps.proto defines a specification of OpenConfig RPC support or requirements
// of a networking entity.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ocrpcs.proto

package ocrpcs_go_proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OCRPCs is the complete list of all OpenConfig RPCs associated with some
// entity (e.g. NOS, or RPC requirements list for a particular device role).
type OCRPCs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required OpenConfig RPC service support for this entity.
	//
	// The key of this map is the full name of the gRPC protocol in lower caps.
	// Examples: gnoi, gnmi, gnsi, gribi.
	OcProtocols map[string]*OCProtocol `protobuf:"bytes,1,rep,name=oc_protocols,json=ocProtocols,proto3" json:"oc_protocols,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *OCRPCs) Reset() {
	*x = OCRPCs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocrpcs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCRPCs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCRPCs) ProtoMessage() {}

func (x *OCRPCs) ProtoReflect() protoreflect.Message {
	mi := &file_ocrpcs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCRPCs.ProtoReflect.Descriptor instead.
func (*OCRPCs) Descriptor() ([]byte, []int) {
	return file_ocrpcs_proto_rawDescGZIP(), []int{0}
}

func (x *OCRPCs) GetOcProtocols() map[string]*OCProtocol {
	if x != nil {
		return x.OcProtocols
	}
	return nil
}

// OCProtocol is the list of OpenConfig RPC methods supported or required that
// belong to the same OpenConfig RPC protocol (e.g. gNOI, gNMI, gRIBI, gNSI).
type OCProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full name of the gRPC method specification.
	// Format: <package>.<service>.<rpc>
	// Example: gnoi.healthz.Healthz.Get
	// Example: gnmi.gNMI.Subscribe
	MethodName []string `protobuf:"bytes,1,rep,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// The semantic version of the gRPC protocol release.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *OCProtocol) Reset() {
	*x = OCProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocrpcs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OCProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCProtocol) ProtoMessage() {}

func (x *OCProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_ocrpcs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCProtocol.ProtoReflect.Descriptor instead.
func (*OCProtocol) Descriptor() ([]byte, []int) {
	return file_ocrpcs_proto_rawDescGZIP(), []int{1}
}

func (x *OCProtocol) GetMethodName() []string {
	if x != nil {
		return x.MethodName
	}
	return nil
}

func (x *OCProtocol) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_ocrpcs_proto protoreflect.FileDescriptor

var file_ocrpcs_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x63, 0x72, 0x70, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x63, 0x72, 0x70, 0x63,
	0x73, 0x22, 0xb6, 0x01, 0x0a, 0x06, 0x4f, 0x43, 0x52, 0x50, 0x43, 0x73, 0x12, 0x4d, 0x0a, 0x0c,
	0x6f, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x6f, 0x63, 0x72, 0x70, 0x63, 0x73, 0x2e, 0x4f, 0x43, 0x52, 0x50, 0x43, 0x73, 0x2e, 0x4f, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x6f, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x1a, 0x5d, 0x0a, 0x10, 0x4f,
	0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x63,
	0x72, 0x70, 0x63, 0x73, 0x2e, 0x4f, 0x43, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x0a, 0x4f, 0x43,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocrpcs_proto_rawDescOnce sync.Once
	file_ocrpcs_proto_rawDescData = file_ocrpcs_proto_rawDesc
)

func file_ocrpcs_proto_rawDescGZIP() []byte {
	file_ocrpcs_proto_rawDescOnce.Do(func() {
		file_ocrpcs_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocrpcs_proto_rawDescData)
	})
	return file_ocrpcs_proto_rawDescData
}

var file_ocrpcs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ocrpcs_proto_goTypes = []interface{}{
	(*OCRPCs)(nil),     // 0: openconfig.ocrpcs.OCRPCs
	(*OCProtocol)(nil), // 1: openconfig.ocrpcs.OCProtocol
	nil,                // 2: openconfig.ocrpcs.OCRPCs.OcProtocolsEntry
}
var file_ocrpcs_proto_depIdxs = []int32{
	2, // 0: openconfig.ocrpcs.OCRPCs.oc_protocols:type_name -> openconfig.ocrpcs.OCRPCs.OcProtocolsEntry
	1, // 1: openconfig.ocrpcs.OCRPCs.OcProtocolsEntry.value:type_name -> openconfig.ocrpcs.OCProtocol
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ocrpcs_proto_init() }
func file_ocrpcs_proto_init() {
	if File_ocrpcs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocrpcs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCRPCs); i {
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
		file_ocrpcs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OCProtocol); i {
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
			RawDescriptor: file_ocrpcs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ocrpcs_proto_goTypes,
		DependencyIndexes: file_ocrpcs_proto_depIdxs,
		MessageInfos:      file_ocrpcs_proto_msgTypes,
	}.Build()
	File_ocrpcs_proto = out.File
	file_ocrpcs_proto_rawDesc = nil
	file_ocrpcs_proto_goTypes = nil
	file_ocrpcs_proto_depIdxs = nil
}

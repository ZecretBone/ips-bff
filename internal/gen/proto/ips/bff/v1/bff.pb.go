// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ips/bff/v1/bff.proto

package bffv1

import (
	v1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/shared/bff/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestMessage string `protobuf:"bytes,1,opt,name=testMessage,proto3" json:"testMessage,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_bff_v1_bff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ips_bff_v1_bff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_ips_bff_v1_bff_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckRequest) GetTestMessage() string {
	if x != nil {
		return x.TestMessage
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string          `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Status   v1.HealthStatus `protobuf:"varint,2,opt,name=status,proto3,enum=ips.shared.bff.v1.HealthStatus" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ips_bff_v1_bff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ips_bff_v1_bff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_ips_bff_v1_bff_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *HealthCheckResponse) GetStatus() v1.HealthStatus {
	if x != nil {
		return x.Status
	}
	return v1.HealthStatus(0)
}

var File_ips_bff_v1_bff_proto protoreflect.FileDescriptor

var file_ips_bff_v1_bff_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x70, 0x73, 0x2f, 0x62, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x66, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x70, 0x73, 0x2e, 0x62, 0x66, 0x66, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x69, 0x70, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f,
	0x62, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x36, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6a, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x69, 0x70,
	0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0x5c, 0x0a, 0x0a, 0x42, 0x46, 0x46, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x1e, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x62, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xa7, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x70, 0x73, 0x2e, 0x62,
	0x66, 0x66, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x42, 0x66, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x42, 0x6f, 0x6e, 0x65, 0x2f, 0x69, 0x70, 0x73, 0x2d, 0x62, 0x66, 0x66,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x70, 0x73, 0x2f, 0x62, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x62,
	0x66, 0x66, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x42, 0x58, 0xaa, 0x02, 0x0a, 0x49, 0x70, 0x73,
	0x2e, 0x42, 0x66, 0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x49, 0x70, 0x73, 0x5c, 0x42, 0x66,
	0x66, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x49, 0x70, 0x73, 0x5c, 0x42, 0x66, 0x66, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x49, 0x70, 0x73, 0x3a, 0x3a, 0x42, 0x66, 0x66, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ips_bff_v1_bff_proto_rawDescOnce sync.Once
	file_ips_bff_v1_bff_proto_rawDescData = file_ips_bff_v1_bff_proto_rawDesc
)

func file_ips_bff_v1_bff_proto_rawDescGZIP() []byte {
	file_ips_bff_v1_bff_proto_rawDescOnce.Do(func() {
		file_ips_bff_v1_bff_proto_rawDescData = protoimpl.X.CompressGZIP(file_ips_bff_v1_bff_proto_rawDescData)
	})
	return file_ips_bff_v1_bff_proto_rawDescData
}

var file_ips_bff_v1_bff_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ips_bff_v1_bff_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),  // 0: ips.bff.v1.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 1: ips.bff.v1.HealthCheckResponse
	(v1.HealthStatus)(0),        // 2: ips.shared.bff.v1.HealthStatus
}
var file_ips_bff_v1_bff_proto_depIdxs = []int32{
	2, // 0: ips.bff.v1.HealthCheckResponse.status:type_name -> ips.shared.bff.v1.HealthStatus
	0, // 1: ips.bff.v1.BFFService.HealthCheck:input_type -> ips.bff.v1.HealthCheckRequest
	1, // 2: ips.bff.v1.BFFService.HealthCheck:output_type -> ips.bff.v1.HealthCheckResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ips_bff_v1_bff_proto_init() }
func file_ips_bff_v1_bff_proto_init() {
	if File_ips_bff_v1_bff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ips_bff_v1_bff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_ips_bff_v1_bff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_ips_bff_v1_bff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ips_bff_v1_bff_proto_goTypes,
		DependencyIndexes: file_ips_bff_v1_bff_proto_depIdxs,
		MessageInfos:      file_ips_bff_v1_bff_proto_msgTypes,
	}.Build()
	File_ips_bff_v1_bff_proto = out.File
	file_ips_bff_v1_bff_proto_rawDesc = nil
	file_ips_bff_v1_bff_proto_goTypes = nil
	file_ips_bff_v1_bff_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ips/shared/rssi/v1/stat.proto

package rssiv1

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

type StatCollectionStage int32

const (
	StatCollectionStage_STAT_COLLECTION_STAGE_UNDEFINED StatCollectionStage = 0
	StatCollectionStage_STAT_COLLECTION_STAGE_SINGLE    StatCollectionStage = 1
	StatCollectionStage_STAT_COLLECTION_STAGE_MULTIPLE  StatCollectionStage = 2
)

// Enum value maps for StatCollectionStage.
var (
	StatCollectionStage_name = map[int32]string{
		0: "STAT_COLLECTION_STAGE_UNDEFINED",
		1: "STAT_COLLECTION_STAGE_SINGLE",
		2: "STAT_COLLECTION_STAGE_MULTIPLE",
	}
	StatCollectionStage_value = map[string]int32{
		"STAT_COLLECTION_STAGE_UNDEFINED": 0,
		"STAT_COLLECTION_STAGE_SINGLE":    1,
		"STAT_COLLECTION_STAGE_MULTIPLE":  2,
	}
)

func (x StatCollectionStage) Enum() *StatCollectionStage {
	p := new(StatCollectionStage)
	*p = x
	return p
}

func (x StatCollectionStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatCollectionStage) Descriptor() protoreflect.EnumDescriptor {
	return file_ips_shared_rssi_v1_stat_proto_enumTypes[0].Descriptor()
}

func (StatCollectionStage) Type() protoreflect.EnumType {
	return &file_ips_shared_rssi_v1_stat_proto_enumTypes[0]
}

func (x StatCollectionStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatCollectionStage.Descriptor instead.
func (StatCollectionStage) EnumDescriptor() ([]byte, []int) {
	return file_ips_shared_rssi_v1_stat_proto_rawDescGZIP(), []int{0}
}

var File_ips_shared_rssi_v1_stat_proto protoreflect.FileDescriptor

var file_ips_shared_rssi_v1_stat_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x70, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x72, 0x73, 0x73,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x69, 0x70, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x73, 0x73, 0x69,
	0x2e, 0x76, 0x31, 0x2a, 0x80, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x53,
	0x54, 0x41, 0x54, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x41, 0x54, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x54, 0x41, 0x54, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x4d, 0x55, 0x4c, 0x54,
	0x49, 0x50, 0x4c, 0x45, 0x10, 0x02, 0x42, 0xda, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x70, 0x73, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x72, 0x73, 0x73, 0x69, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x53, 0x74, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x42, 0x6f, 0x6e, 0x65, 0x2f, 0x69, 0x70, 0x73, 0x2d, 0x62, 0x66, 0x66, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x69, 0x70, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x72, 0x73, 0x73, 0x69,
	0x2f, 0x76, 0x31, 0x3b, 0x72, 0x73, 0x73, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x53, 0x52,
	0xaa, 0x02, 0x12, 0x49, 0x70, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x73,
	0x73, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x49, 0x70, 0x73, 0x5c, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x5c, 0x52, 0x73, 0x73, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x49, 0x70, 0x73,
	0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5c, 0x52, 0x73, 0x73, 0x69, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x49, 0x70,
	0x73, 0x3a, 0x3a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x3a, 0x3a, 0x52, 0x73, 0x73, 0x69, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ips_shared_rssi_v1_stat_proto_rawDescOnce sync.Once
	file_ips_shared_rssi_v1_stat_proto_rawDescData = file_ips_shared_rssi_v1_stat_proto_rawDesc
)

func file_ips_shared_rssi_v1_stat_proto_rawDescGZIP() []byte {
	file_ips_shared_rssi_v1_stat_proto_rawDescOnce.Do(func() {
		file_ips_shared_rssi_v1_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_ips_shared_rssi_v1_stat_proto_rawDescData)
	})
	return file_ips_shared_rssi_v1_stat_proto_rawDescData
}

var file_ips_shared_rssi_v1_stat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ips_shared_rssi_v1_stat_proto_goTypes = []interface{}{
	(StatCollectionStage)(0), // 0: ips.shared.rssi.v1.StatCollectionStage
}
var file_ips_shared_rssi_v1_stat_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ips_shared_rssi_v1_stat_proto_init() }
func file_ips_shared_rssi_v1_stat_proto_init() {
	if File_ips_shared_rssi_v1_stat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ips_shared_rssi_v1_stat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ips_shared_rssi_v1_stat_proto_goTypes,
		DependencyIndexes: file_ips_shared_rssi_v1_stat_proto_depIdxs,
		EnumInfos:         file_ips_shared_rssi_v1_stat_proto_enumTypes,
	}.Build()
	File_ips_shared_rssi_v1_stat_proto = out.File
	file_ips_shared_rssi_v1_stat_proto_rawDesc = nil
	file_ips_shared_rssi_v1_stat_proto_goTypes = nil
	file_ips_shared_rssi_v1_stat_proto_depIdxs = nil
}

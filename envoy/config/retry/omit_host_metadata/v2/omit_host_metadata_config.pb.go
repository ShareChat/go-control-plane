// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/config/retry/omit_host_metadata/v2/omit_host_metadata_config.proto

package omit_host_metadatav2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
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

// A retry host predicate that can be used to reject a host based on
// predefined metadata match criteria.
// [#extension: envoy.retry_host_predicates.omit_host_metadata]
type OmitHostMetadataConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Retry host predicate metadata match criteria. The hosts in
	// the upstream cluster with matching metadata will be omitted while
	// attempting a retry of a failed request. The metadata should be specified
	// under the *envoy.lb* key.
	MetadataMatch *core.Metadata `protobuf:"bytes,1,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
}

func (x *OmitHostMetadataConfig) Reset() {
	*x = OmitHostMetadataConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OmitHostMetadataConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OmitHostMetadataConfig) ProtoMessage() {}

func (x *OmitHostMetadataConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OmitHostMetadataConfig.ProtoReflect.Descriptor instead.
func (*OmitHostMetadataConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDescGZIP(), []int{0}
}

func (x *OmitHostMetadataConfig) GetMetadataMatch() *core.Metadata {
	if x != nil {
		return x.MetadataMatch
	}
	return nil
}

var File_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto protoreflect.FileDescriptor

var file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDesc = []byte{
	0x0a, 0x48, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x2f, 0x6f, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x6d, 0x69, 0x74, 0x5f,
	0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x6f,
	0x6d, 0x69, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5c, 0x0a, 0x16, 0x4f, 0x6d, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x0a, 0x0e, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42,
	0xfe, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x33, 0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x6f, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x01, 0x0a, 0x36, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x6f, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x32, 0x42, 0x1b, 0x4f, 0x6d, 0x69,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x6f, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x32, 0x3b, 0x6f, 0x6d, 0x69, 0x74,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDescOnce sync.Once
	file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDescData = file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDesc
)

func file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDescGZIP() []byte {
	file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDescData)
	})
	return file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDescData
}

var file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_goTypes = []interface{}{
	(*OmitHostMetadataConfig)(nil), // 0: envoy.config.retry.omit_host_metadata.v2.OmitHostMetadataConfig
	(*core.Metadata)(nil),          // 1: envoy.api.v2.core.Metadata
}
var file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_depIdxs = []int32{
	1, // 0: envoy.config.retry.omit_host_metadata.v2.OmitHostMetadataConfig.metadata_match:type_name -> envoy.api.v2.core.Metadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_init() }
func file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_init() {
	if File_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OmitHostMetadataConfig); i {
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
			RawDescriptor: file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_msgTypes,
	}.Build()
	File_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto = out.File
	file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_rawDesc = nil
	file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_goTypes = nil
	file_envoy_config_retry_omit_host_metadata_v2_omit_host_metadata_config_proto_depIdxs = nil
}

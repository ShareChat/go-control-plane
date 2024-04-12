// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/load_balancing_policies/random/v3/random.proto

package randomv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/load_balancing_policies/common/v3"
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

// This configuration allows the built-in Random LB policy to be configured via the LB policy
// extension point. See the :ref:`load balancing architecture overview
// <arch_overview_load_balancing_types>` for more information.
type Random struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for local zone aware load balancing or locality weighted load balancing.
	LocalityLbConfig *v3.LocalityLbConfig `protobuf:"bytes,1,opt,name=locality_lb_config,json=localityLbConfig,proto3" json:"locality_lb_config,omitempty"`
}

func (x *Random) Reset() {
	*x = Random{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_load_balancing_policies_random_v3_random_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Random) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Random) ProtoMessage() {}

func (x *Random) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_load_balancing_policies_random_v3_random_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Random.ProtoReflect.Descriptor instead.
func (*Random) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDescGZIP(), []int{0}
}

func (x *Random) GetLocalityLbConfig() *v3.LocalityLbConfig {
	if x != nil {
		return x.LocalityLbConfig
	}
	return nil
}

var File_envoy_extensions_load_balancing_policies_random_v3_random_proto protoreflect.FileDescriptor

var file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x72, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x2e, 0x76, 0x33, 0x1a, 0x3f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12,
	0x72, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x62, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0xbd, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x40,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x2e, 0x76, 0x33,
	0x42, 0x0b, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x62, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x2f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x2f, 0x76, 0x33, 0x3b, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDescOnce sync.Once
	file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDescData = file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDesc
)

func file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDescGZIP() []byte {
	file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDescData)
	})
	return file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDescData
}

var file_envoy_extensions_load_balancing_policies_random_v3_random_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_load_balancing_policies_random_v3_random_proto_goTypes = []interface{}{
	(*Random)(nil),              // 0: envoy.extensions.load_balancing_policies.random.v3.Random
	(*v3.LocalityLbConfig)(nil), // 1: envoy.extensions.load_balancing_policies.common.v3.LocalityLbConfig
}
var file_envoy_extensions_load_balancing_policies_random_v3_random_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.load_balancing_policies.random.v3.Random.locality_lb_config:type_name -> envoy.extensions.load_balancing_policies.common.v3.LocalityLbConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_load_balancing_policies_random_v3_random_proto_init() }
func file_envoy_extensions_load_balancing_policies_random_v3_random_proto_init() {
	if File_envoy_extensions_load_balancing_policies_random_v3_random_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_load_balancing_policies_random_v3_random_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Random); i {
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
			RawDescriptor: file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_load_balancing_policies_random_v3_random_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_load_balancing_policies_random_v3_random_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_load_balancing_policies_random_v3_random_proto_msgTypes,
	}.Build()
	File_envoy_extensions_load_balancing_policies_random_v3_random_proto = out.File
	file_envoy_extensions_load_balancing_policies_random_v3_random_proto_rawDesc = nil
	file_envoy_extensions_load_balancing_policies_random_v3_random_proto_goTypes = nil
	file_envoy_extensions_load_balancing_policies_random_v3_random_proto_depIdxs = nil
}

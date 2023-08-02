// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/quic/server_preferred_address/v3/fixed_server_preferred_address_config.proto

package server_preferred_addressv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
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

// Configuration for FixedServerPreferredAddressConfig.
type FixedServerPreferredAddressConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Ipv4Type:
	//	*FixedServerPreferredAddressConfig_Ipv4Address
	Ipv4Type isFixedServerPreferredAddressConfig_Ipv4Type `protobuf_oneof:"ipv4_type"`
	// Types that are assignable to Ipv6Type:
	//	*FixedServerPreferredAddressConfig_Ipv6Address
	Ipv6Type isFixedServerPreferredAddressConfig_Ipv6Type `protobuf_oneof:"ipv6_type"`
}

func (x *FixedServerPreferredAddressConfig) Reset() {
	*x = FixedServerPreferredAddressConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FixedServerPreferredAddressConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedServerPreferredAddressConfig) ProtoMessage() {}

func (x *FixedServerPreferredAddressConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedServerPreferredAddressConfig.ProtoReflect.Descriptor instead.
func (*FixedServerPreferredAddressConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDescGZIP(), []int{0}
}

func (m *FixedServerPreferredAddressConfig) GetIpv4Type() isFixedServerPreferredAddressConfig_Ipv4Type {
	if m != nil {
		return m.Ipv4Type
	}
	return nil
}

func (x *FixedServerPreferredAddressConfig) GetIpv4Address() string {
	if x, ok := x.GetIpv4Type().(*FixedServerPreferredAddressConfig_Ipv4Address); ok {
		return x.Ipv4Address
	}
	return ""
}

func (m *FixedServerPreferredAddressConfig) GetIpv6Type() isFixedServerPreferredAddressConfig_Ipv6Type {
	if m != nil {
		return m.Ipv6Type
	}
	return nil
}

func (x *FixedServerPreferredAddressConfig) GetIpv6Address() string {
	if x, ok := x.GetIpv6Type().(*FixedServerPreferredAddressConfig_Ipv6Address); ok {
		return x.Ipv6Address
	}
	return ""
}

type isFixedServerPreferredAddressConfig_Ipv4Type interface {
	isFixedServerPreferredAddressConfig_Ipv4Type()
}

type FixedServerPreferredAddressConfig_Ipv4Address struct {
	// String representation of IPv4 address, i.e. "127.0.0.2".
	// If not specified, none will be configured.
	Ipv4Address string `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3,oneof"`
}

func (*FixedServerPreferredAddressConfig_Ipv4Address) isFixedServerPreferredAddressConfig_Ipv4Type() {
}

type isFixedServerPreferredAddressConfig_Ipv6Type interface {
	isFixedServerPreferredAddressConfig_Ipv6Type()
}

type FixedServerPreferredAddressConfig_Ipv6Address struct {
	// String representation of IPv6 address, i.e. "::1".
	// If not specified, none will be configured.
	Ipv6Address string `protobuf:"bytes,2,opt,name=ipv6_address,json=ipv6Address,proto3,oneof"`
}

func (*FixedServerPreferredAddressConfig_Ipv6Address) isFixedServerPreferredAddressConfig_Ipv6Type() {
}

var File_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x33, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x21, 0x46, 0x69, 0x78, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0c, 0x69, 0x70, 0x76, 0x34,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x69, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a,
	0x0c, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x69, 0x70, 0x76, 0x36, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x3a, 0x08, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x42, 0x0b, 0x0a, 0x09,
	0x69, 0x70, 0x76, 0x34, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x69, 0x70, 0x76,
	0x36, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xe8, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x0a, 0x3f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x33, 0x42, 0x26, 0x46, 0x69, 0x78, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x73, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x76,
	0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDescData = file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDesc
)

func file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDescData)
	})
	return file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDescData
}

var file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_goTypes = []interface{}{
	(*FixedServerPreferredAddressConfig)(nil), // 0: envoy.extensions.quic.server_preferred_address.v3.FixedServerPreferredAddressConfig
}
var file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_init()
}
func file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_init() {
	if File_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FixedServerPreferredAddressConfig); i {
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
	file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FixedServerPreferredAddressConfig_Ipv4Address)(nil),
		(*FixedServerPreferredAddressConfig_Ipv6Address)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto = out.File
	file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_rawDesc = nil
	file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_goTypes = nil
	file_envoy_extensions_quic_server_preferred_address_v3_fixed_server_preferred_address_config_proto_depIdxs = nil
}

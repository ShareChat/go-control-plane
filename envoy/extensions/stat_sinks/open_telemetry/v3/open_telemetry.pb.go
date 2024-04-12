// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/stat_sinks/open_telemetry/v3/open_telemetry.proto

package open_telemetryv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// [#next-free-field: 7]
type SinkConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ProtocolSpecifier:
	//	*SinkConfig_GrpcService
	ProtocolSpecifier isSinkConfig_ProtocolSpecifier `protobuf_oneof:"protocol_specifier"`
	// If set to true, counters will be emitted as deltas, and the OTLP message will have
	// ``AGGREGATION_TEMPORALITY_DELTA`` set as AggregationTemporality.
	ReportCountersAsDeltas bool `protobuf:"varint,2,opt,name=report_counters_as_deltas,json=reportCountersAsDeltas,proto3" json:"report_counters_as_deltas,omitempty"`
	// If set to true, histograms will be emitted as deltas, and the OTLP message will have
	// ``AGGREGATION_TEMPORALITY_DELTA`` set as AggregationTemporality.
	ReportHistogramsAsDeltas bool `protobuf:"varint,3,opt,name=report_histograms_as_deltas,json=reportHistogramsAsDeltas,proto3" json:"report_histograms_as_deltas,omitempty"`
	// If set to true, metrics will have their tags emitted as OTLP attributes, which may
	// contain values used by the tag extractor or additional tags added during stats creation.
	// Otherwise, no attributes will be associated with the export message. Default value is true.
	EmitTagsAsAttributes *wrappers.BoolValue `protobuf:"bytes,4,opt,name=emit_tags_as_attributes,json=emitTagsAsAttributes,proto3" json:"emit_tags_as_attributes,omitempty"`
	// If set to true, metric names will be represented as the tag extracted name instead
	// of the full metric name. Default value is true.
	UseTagExtractedName *wrappers.BoolValue `protobuf:"bytes,5,opt,name=use_tag_extracted_name,json=useTagExtractedName,proto3" json:"use_tag_extracted_name,omitempty"`
	// If set, emitted stats names will be prepended with a prefix, so full stat name will be
	// <prefix>.<stats_name>. For example, if the stat name is "foo.bar" and prefix is
	// "pre", the full stat name will be "pre.foo.bar". If this field is not set, there is no
	// prefix added. According to the example, the full stat name will remain "foo.bar".
	Prefix string `protobuf:"bytes,6,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *SinkConfig) Reset() {
	*x = SinkConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SinkConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SinkConfig) ProtoMessage() {}

func (x *SinkConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SinkConfig.ProtoReflect.Descriptor instead.
func (*SinkConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDescGZIP(), []int{0}
}

func (m *SinkConfig) GetProtocolSpecifier() isSinkConfig_ProtocolSpecifier {
	if m != nil {
		return m.ProtocolSpecifier
	}
	return nil
}

func (x *SinkConfig) GetGrpcService() *v3.GrpcService {
	if x, ok := x.GetProtocolSpecifier().(*SinkConfig_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (x *SinkConfig) GetReportCountersAsDeltas() bool {
	if x != nil {
		return x.ReportCountersAsDeltas
	}
	return false
}

func (x *SinkConfig) GetReportHistogramsAsDeltas() bool {
	if x != nil {
		return x.ReportHistogramsAsDeltas
	}
	return false
}

func (x *SinkConfig) GetEmitTagsAsAttributes() *wrappers.BoolValue {
	if x != nil {
		return x.EmitTagsAsAttributes
	}
	return nil
}

func (x *SinkConfig) GetUseTagExtractedName() *wrappers.BoolValue {
	if x != nil {
		return x.UseTagExtractedName
	}
	return nil
}

func (x *SinkConfig) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type isSinkConfig_ProtocolSpecifier interface {
	isSinkConfig_ProtocolSpecifier()
}

type SinkConfig_GrpcService struct {
	// The upstream gRPC cluster that implements the OTLP/gRPC collector.
	GrpcService *v3.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

func (*SinkConfig_GrpcService) isSinkConfig_ProtocolSpecifier() {}

var File_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto protoreflect.FileDescriptor

var file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDesc = []byte{
	0x0a, 0x42, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x73, 0x69, 0x6e, 0x6b,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x33, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x03, 0x0a, 0x0a, 0x53, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x41, 0x73, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73,
	0x12, 0x3d, 0x0a, 0x1b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x61, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x41, 0x73, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x12,
	0x51, 0x0a, 0x17, 0x65, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x5f, 0x61, 0x73, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x65, 0x6d,
	0x69, 0x74, 0x54, 0x61, 0x67, 0x73, 0x41, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x4f, 0x0a, 0x16, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13,
	0x75, 0x73, 0x65, 0x54, 0x61, 0x67, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x19, 0x0a, 0x12, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0xc2, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x12,
	0x4f, 0x70, 0x65, 0x6e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x6f, 0x70, 0x65, 0x6e, 0x5f,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDescOnce sync.Once
	file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDescData = file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDesc
)

func file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDescGZIP() []byte {
	file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDescData)
	})
	return file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDescData
}

var file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_goTypes = []interface{}{
	(*SinkConfig)(nil),         // 0: envoy.extensions.stat_sinks.open_telemetry.v3.SinkConfig
	(*v3.GrpcService)(nil),     // 1: envoy.config.core.v3.GrpcService
	(*wrappers.BoolValue)(nil), // 2: google.protobuf.BoolValue
}
var file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.stat_sinks.open_telemetry.v3.SinkConfig.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	2, // 1: envoy.extensions.stat_sinks.open_telemetry.v3.SinkConfig.emit_tags_as_attributes:type_name -> google.protobuf.BoolValue
	2, // 2: envoy.extensions.stat_sinks.open_telemetry.v3.SinkConfig.use_tag_extracted_name:type_name -> google.protobuf.BoolValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_init() }
func file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_init() {
	if File_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SinkConfig); i {
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
	file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SinkConfig_GrpcService)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_msgTypes,
	}.Build()
	File_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto = out.File
	file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_rawDesc = nil
	file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_goTypes = nil
	file_envoy_extensions_stat_sinks_open_telemetry_v3_open_telemetry_proto_depIdxs = nil
}

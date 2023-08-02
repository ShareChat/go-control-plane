// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/http/header_validators/envoy_default/v3/header_validator.proto

package envoy_defaultv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Action to take when Envoy receives client request with header names containing underscore
// characters.
// Underscore character is allowed in header names by the RFC-7230 and this behavior is implemented
// as a security measure due to systems that treat '_' and '-' as interchangeable. Envoy by default allows client request headers with underscore
// characters.
type HeaderValidatorConfig_HeadersWithUnderscoresAction int32

const (
	// Allow headers with underscores. This is the default behavior.
	HeaderValidatorConfig_ALLOW HeaderValidatorConfig_HeadersWithUnderscoresAction = 0
	// Reject client request. HTTP/1 requests are rejected with the 400 status. HTTP/2 requests
	// end with the stream reset. The
	// :ref:`httpN.requests_rejected_with_underscores_in_headers <config_http_conn_man_stats_per_codec>` counter
	// is incremented for each rejected request.
	HeaderValidatorConfig_REJECT_REQUEST HeaderValidatorConfig_HeadersWithUnderscoresAction = 1
	// Drop the client header with name containing underscores. The header is dropped before the filter chain is
	// invoked and as such filters will not see dropped headers. The
	// :ref:`httpN.dropped_headers_with_underscores <config_http_conn_man_stats_per_codec>` is incremented for
	// each dropped header.
	HeaderValidatorConfig_DROP_HEADER HeaderValidatorConfig_HeadersWithUnderscoresAction = 2
)

// Enum value maps for HeaderValidatorConfig_HeadersWithUnderscoresAction.
var (
	HeaderValidatorConfig_HeadersWithUnderscoresAction_name = map[int32]string{
		0: "ALLOW",
		1: "REJECT_REQUEST",
		2: "DROP_HEADER",
	}
	HeaderValidatorConfig_HeadersWithUnderscoresAction_value = map[string]int32{
		"ALLOW":          0,
		"REJECT_REQUEST": 1,
		"DROP_HEADER":    2,
	}
)

func (x HeaderValidatorConfig_HeadersWithUnderscoresAction) Enum() *HeaderValidatorConfig_HeadersWithUnderscoresAction {
	p := new(HeaderValidatorConfig_HeadersWithUnderscoresAction)
	*p = x
	return p
}

func (x HeaderValidatorConfig_HeadersWithUnderscoresAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HeaderValidatorConfig_HeadersWithUnderscoresAction) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_enumTypes[0].Descriptor()
}

func (HeaderValidatorConfig_HeadersWithUnderscoresAction) Type() protoreflect.EnumType {
	return &file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_enumTypes[0]
}

func (x HeaderValidatorConfig_HeadersWithUnderscoresAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HeaderValidatorConfig_HeadersWithUnderscoresAction.Descriptor instead.
func (HeaderValidatorConfig_HeadersWithUnderscoresAction) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescGZIP(), []int{0, 0}
}

// Determines the action for requests that contain ``%2F``, ``%2f``, ``%5C`` or ``%5c`` sequences in the URI path.
// This operation occurs before URL normalization and the merge slashes transformations if they were enabled.
type HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction int32

const (
	// Default behavior specific to implementation (i.e. Envoy) of this configuration option.
	// Envoy, by default, takes the ``KEEP_UNCHANGED`` action.
	// NOTE: the implementation may change the default behavior at-will.
	HeaderValidatorConfig_UriPathNormalizationOptions_IMPLEMENTATION_SPECIFIC_DEFAULT HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction = 0
	// Keep escaped slashes.
	HeaderValidatorConfig_UriPathNormalizationOptions_KEEP_UNCHANGED HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction = 1
	// Reject client request with the 400 status. gRPC requests will be rejected with the ``INTERNAL`` (13) error code.
	// The ``http#.downstream_rq_failed_path_normalization`` counter is incremented for each rejected request.
	HeaderValidatorConfig_UriPathNormalizationOptions_REJECT_REQUEST HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction = 2
	// Unescape ``%2F`` and ``%5C`` sequences and redirect the request to the new path if these sequences were present.
	// The redirect occurs after path normalization and merge slashes transformations if they were configured.
	// NOTE: gRPC requests will be rejected with the ``INTERNAL`` (13) error code.
	// This option minimizes possibility of path confusion exploits by forcing request with unescaped slashes to
	// traverse all parties: downstream client, intermediate proxies, Envoy and upstream server.
	// The ``http#.downstream_rq_redirected_with_normalized_path`` counter is incremented for each
	// redirected request.
	HeaderValidatorConfig_UriPathNormalizationOptions_UNESCAPE_AND_REDIRECT HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction = 3
	// Unescape ``%2F`` and ``%5C`` sequences.
	// Note: this option should not be enabled if intermediaries perform path based access control as
	// it may lead to path confusion vulnerabilities.
	HeaderValidatorConfig_UriPathNormalizationOptions_UNESCAPE_AND_FORWARD HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction = 4
)

// Enum value maps for HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction.
var (
	HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction_name = map[int32]string{
		0: "IMPLEMENTATION_SPECIFIC_DEFAULT",
		1: "KEEP_UNCHANGED",
		2: "REJECT_REQUEST",
		3: "UNESCAPE_AND_REDIRECT",
		4: "UNESCAPE_AND_FORWARD",
	}
	HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction_value = map[string]int32{
		"IMPLEMENTATION_SPECIFIC_DEFAULT": 0,
		"KEEP_UNCHANGED":                  1,
		"REJECT_REQUEST":                  2,
		"UNESCAPE_AND_REDIRECT":           3,
		"UNESCAPE_AND_FORWARD":            4,
	}
)

func (x HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction) Enum() *HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction {
	p := new(HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction)
	*p = x
	return p
}

func (x HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_enumTypes[1].Descriptor()
}

func (HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction) Type() protoreflect.EnumType {
	return &file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_enumTypes[1]
}

func (x HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction.Descriptor instead.
func (HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescGZIP(), []int{0, 0, 0}
}

// This extension validates that HTTP request and response headers are well formed according to respective RFCs.
//
// #. HTTP/1 header map validity according to `RFC 7230 section 3.2 <https://datatracker.ietf.org/doc/html/rfc7230#section-3.2>`_
// #. Syntax of HTTP/1 request target URI and response status
// #. HTTP/2 header map validity according to `RFC 7540 section 8.1.2 <https://datatracker.ietf.org/doc/html/rfc7540#section-8.1.2>`_
// #. Syntax of HTTP/2 pseudo headers
// #. HTTP/3 header map validity according to `RFC 9114 section 4.3  <https://www.rfc-editor.org/rfc/rfc9114.html>`_
// #. Syntax of HTTP/3 pseudo headers
// #. Syntax of Content-Length and Transfer-Encoding
// #. Validation of HTTP/1 requests with both ``Content-Length`` and ``Transfer-Encoding`` headers
// #. Normalization of the URI path according to `Normalization and Comparison <https://datatracker.ietf.org/doc/html/rfc3986#section-6>`_
//    without `case normalization <https://datatracker.ietf.org/doc/html/rfc3986#section-6.2.2.1>`_
//
// [#comment:TODO(yanavlasov): Put #extension: envoy.http.header_validators.envoy_default after it is not hidden any more]
// [#next-free-field: 6]
type HeaderValidatorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http1ProtocolOptions *HeaderValidatorConfig_Http1ProtocolOptions `protobuf:"bytes,1,opt,name=http1_protocol_options,json=http1ProtocolOptions,proto3" json:"http1_protocol_options,omitempty"`
	// The URI path normalization options.
	// By default Envoy normalizes URI path using the default values of the :ref:`UriPathNormalizationOptions
	// <envoy_v3_api_msg_extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.UriPathNormalizationOptions>`.
	// URI path transformations specified by the ``uri_path_normalization_options`` configuration can be applied to a portion
	// of requests by setting the ``envoy_default_header_validator.uri_path_transformations`` runtime value.
	// Caution: disabling path normalization may lead to path confusion vulnerabilities in access control or incorrect service
	// selection.
	UriPathNormalizationOptions *HeaderValidatorConfig_UriPathNormalizationOptions `protobuf:"bytes,2,opt,name=uri_path_normalization_options,json=uriPathNormalizationOptions,proto3" json:"uri_path_normalization_options,omitempty"`
	// Restrict HTTP methods to these defined in the `RFC 7231 section 4.1 <https://datatracker.ietf.org/doc/html/rfc7231#section-4.1>`_
	// Envoy will respond with 400 to requests with disallowed methods.
	// By default methods with arbitrary names are accepted.
	RestrictHttpMethods bool `protobuf:"varint,3,opt,name=restrict_http_methods,json=restrictHttpMethods,proto3" json:"restrict_http_methods,omitempty"`
	// Action to take when a client request with a header name containing underscore characters is received.
	// If this setting is not specified, the value defaults to ALLOW.
	HeadersWithUnderscoresAction HeaderValidatorConfig_HeadersWithUnderscoresAction `protobuf:"varint,4,opt,name=headers_with_underscores_action,json=headersWithUnderscoresAction,proto3,enum=envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig_HeadersWithUnderscoresAction" json:"headers_with_underscores_action,omitempty"`
	// Allow requests with fragment in URL path and strip the fragment before request processing.
	// By default Envoy rejects requests with fragment in URL path.
	StripFragmentFromPath bool `protobuf:"varint,5,opt,name=strip_fragment_from_path,json=stripFragmentFromPath,proto3" json:"strip_fragment_from_path,omitempty"`
}

func (x *HeaderValidatorConfig) Reset() {
	*x = HeaderValidatorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValidatorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValidatorConfig) ProtoMessage() {}

func (x *HeaderValidatorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValidatorConfig.ProtoReflect.Descriptor instead.
func (*HeaderValidatorConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderValidatorConfig) GetHttp1ProtocolOptions() *HeaderValidatorConfig_Http1ProtocolOptions {
	if x != nil {
		return x.Http1ProtocolOptions
	}
	return nil
}

func (x *HeaderValidatorConfig) GetUriPathNormalizationOptions() *HeaderValidatorConfig_UriPathNormalizationOptions {
	if x != nil {
		return x.UriPathNormalizationOptions
	}
	return nil
}

func (x *HeaderValidatorConfig) GetRestrictHttpMethods() bool {
	if x != nil {
		return x.RestrictHttpMethods
	}
	return false
}

func (x *HeaderValidatorConfig) GetHeadersWithUnderscoresAction() HeaderValidatorConfig_HeadersWithUnderscoresAction {
	if x != nil {
		return x.HeadersWithUnderscoresAction
	}
	return HeaderValidatorConfig_ALLOW
}

func (x *HeaderValidatorConfig) GetStripFragmentFromPath() bool {
	if x != nil {
		return x.StripFragmentFromPath
	}
	return false
}

type HeaderValidatorConfig_UriPathNormalizationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Should paths be normalized according to RFC 3986?
	// This operation overwrites the original request URI path and the new path is used for processing of
	// the request by HTTP filters and proxied to the upstream service.
	// Envoy will respond with 400 to requests with malformed paths that fail path normalization.
	// The default behavior is to normalize the path.
	// This value may be overridden by the runtime variable
	// :ref:`http_connection_manager.normalize_path<config_http_conn_man_runtime_normalize_path>`.
	// See `Normalization and Comparison <https://datatracker.ietf.org/doc/html/rfc3986#section-6>`_
	// for details of normalization.
	// Note that Envoy does not perform
	// `case normalization <https://datatracker.ietf.org/doc/html/rfc3986#section-6.2.2.1>`_
	// URI path normalization can be applied to a portion of requests by setting the
	// ``envoy_default_header_validator.path_normalization`` runtime value.
	SkipPathNormalization bool `protobuf:"varint,1,opt,name=skip_path_normalization,json=skipPathNormalization,proto3" json:"skip_path_normalization,omitempty"`
	// Determines if adjacent slashes in the path are merged into one.
	// This operation overwrites the original request URI path and the new path is used for processing of
	// the request by HTTP filters and proxied to the upstream service.
	// Setting this option to true will cause incoming requests with path ``//dir///file`` to not match against
	// route with ``prefix`` match set to ``/dir``. Defaults to ``false``. Note that slash merging is not part of
	// `HTTP spec <https://datatracker.ietf.org/doc/html/rfc3986>`_ and is provided for convenience.
	// Merging of slashes in URI path can be applied to a portion of requests by setting the
	// ``envoy_default_header_validator.merge_slashes`` runtime value.
	SkipMergingSlashes bool `protobuf:"varint,2,opt,name=skip_merging_slashes,json=skipMergingSlashes,proto3" json:"skip_merging_slashes,omitempty"`
	// The action to take when request URL path contains escaped slash sequences (``%2F``, ``%2f``, ``%5C`` and ``%5c``).
	// This operation may overwrite the original request URI path and the new path is used for processing of
	// the request by HTTP filters and proxied to the upstream service.
	PathWithEscapedSlashesAction HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction `protobuf:"varint,3,opt,name=path_with_escaped_slashes_action,json=pathWithEscapedSlashesAction,proto3,enum=envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction" json:"path_with_escaped_slashes_action,omitempty"`
}

func (x *HeaderValidatorConfig_UriPathNormalizationOptions) Reset() {
	*x = HeaderValidatorConfig_UriPathNormalizationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValidatorConfig_UriPathNormalizationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValidatorConfig_UriPathNormalizationOptions) ProtoMessage() {}

func (x *HeaderValidatorConfig_UriPathNormalizationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValidatorConfig_UriPathNormalizationOptions.ProtoReflect.Descriptor instead.
func (*HeaderValidatorConfig_UriPathNormalizationOptions) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HeaderValidatorConfig_UriPathNormalizationOptions) GetSkipPathNormalization() bool {
	if x != nil {
		return x.SkipPathNormalization
	}
	return false
}

func (x *HeaderValidatorConfig_UriPathNormalizationOptions) GetSkipMergingSlashes() bool {
	if x != nil {
		return x.SkipMergingSlashes
	}
	return false
}

func (x *HeaderValidatorConfig_UriPathNormalizationOptions) GetPathWithEscapedSlashesAction() HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction {
	if x != nil {
		return x.PathWithEscapedSlashesAction
	}
	return HeaderValidatorConfig_UriPathNormalizationOptions_IMPLEMENTATION_SPECIFIC_DEFAULT
}

type HeaderValidatorConfig_Http1ProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Allows Envoy to process HTTP/1 requests/responses with both ``Content-Length`` and ``Transfer-Encoding``
	// headers set. By default such messages are rejected, but if option is enabled - Envoy will
	// remove the ``Content-Length`` header and process the message.
	// See `RFC7230, sec. 3.3.3 <https://datatracker.ietf.org/doc/html/rfc7230#section-3.3.3>`_ for details.
	//
	// .. attention::
	//   Enabling this option might lead to request smuggling vulnerabilities, especially if traffic
	//   is proxied via multiple layers of proxies.
	AllowChunkedLength bool `protobuf:"varint,1,opt,name=allow_chunked_length,json=allowChunkedLength,proto3" json:"allow_chunked_length,omitempty"`
}

func (x *HeaderValidatorConfig_Http1ProtocolOptions) Reset() {
	*x = HeaderValidatorConfig_Http1ProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValidatorConfig_Http1ProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValidatorConfig_Http1ProtocolOptions) ProtoMessage() {}

func (x *HeaderValidatorConfig_Http1ProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValidatorConfig_Http1ProtocolOptions.ProtoReflect.Descriptor instead.
func (*HeaderValidatorConfig_Http1ProtocolOptions) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescGZIP(), []int{0, 1}
}

func (x *HeaderValidatorConfig_Http1ProtocolOptions) GetAllowChunkedLength() bool {
	if x != nil {
		return x.AllowChunkedLength
	}
	return false
}

var File_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto protoreflect.FileDescriptor

var file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x0a, 0x0a, 0x15, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x9a, 0x01,
	0x0a, 0x16, 0x68, 0x74, 0x74, 0x70, 0x31, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x64,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x31, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x14, 0x68, 0x74, 0x74, 0x70, 0x31, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xb0, 0x01, 0x0a, 0x1e, 0x75,
	0x72, 0x69, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x6b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x72, 0x69, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x1b, 0x75, 0x72, 0x69, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x12, 0xb3, 0x01, 0x0a, 0x1f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x6c, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1c, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x74, 0x72, 0x69, 0x70,
	0x5f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x73, 0x74, 0x72, 0x69, 0x70,
	0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x74, 0x68,
	0x1a, 0x88, 0x04, 0x0a, 0x1b, 0x55, 0x72, 0x69, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x36, 0x0a, 0x17, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x6e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x15, 0x73, 0x6b, 0x69, 0x70, 0x50, 0x61, 0x74, 0x68, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x6b, 0x69, 0x70,
	0x5f, 0x6d, 0x65, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73, 0x6b, 0x69, 0x70, 0x4d, 0x65, 0x72, 0x67,
	0x69, 0x6e, 0x67, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0xdb, 0x01, 0x0a, 0x20, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x64,
	0x5f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x88, 0x01, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76,
	0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x72, 0x69, 0x50, 0x61, 0x74, 0x68, 0x4e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x57, 0x69, 0x74, 0x68, 0x45, 0x73, 0x63, 0x61,
	0x70, 0x65, 0x64, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x1c, 0x70, 0x61, 0x74, 0x68,
	0x57, 0x69, 0x74, 0x68, 0x45, 0x73, 0x63, 0x61, 0x70, 0x65, 0x64, 0x53, 0x6c, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x1c, 0x50, 0x61, 0x74,
	0x68, 0x57, 0x69, 0x74, 0x68, 0x45, 0x73, 0x63, 0x61, 0x70, 0x65, 0x64, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x1f, 0x49, 0x4d, 0x50,
	0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x43, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x4b, 0x45, 0x45, 0x50, 0x5f, 0x55, 0x4e, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x4e, 0x45, 0x53, 0x43, 0x41,
	0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10,
	0x03, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x45, 0x53, 0x43, 0x41, 0x50, 0x45, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x04, 0x1a, 0x48, 0x0a, 0x14, 0x48,
	0x74, 0x74, 0x70, 0x31, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x65, 0x64, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x4e, 0x0a, 0x1c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x48, 0x45, 0x41,
	0x44, 0x45, 0x52, 0x10, 0x02, 0x42, 0xd9, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02,
	0x0a, 0x46, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x14, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76,
	0x33, 0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x76,
	0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescOnce sync.Once
	file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescData = file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDesc
)

func file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescGZIP() []byte {
	file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescData)
	})
	return file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDescData
}

var file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_goTypes = []interface{}{
	(HeaderValidatorConfig_HeadersWithUnderscoresAction)(0),                             // 0: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.HeadersWithUnderscoresAction
	(HeaderValidatorConfig_UriPathNormalizationOptions_PathWithEscapedSlashesAction)(0), // 1: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.UriPathNormalizationOptions.PathWithEscapedSlashesAction
	(*HeaderValidatorConfig)(nil),                                                       // 2: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig
	(*HeaderValidatorConfig_UriPathNormalizationOptions)(nil),                           // 3: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.UriPathNormalizationOptions
	(*HeaderValidatorConfig_Http1ProtocolOptions)(nil),                                  // 4: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.Http1ProtocolOptions
}
var file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.http1_protocol_options:type_name -> envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.Http1ProtocolOptions
	3, // 1: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.uri_path_normalization_options:type_name -> envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.UriPathNormalizationOptions
	0, // 2: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.headers_with_underscores_action:type_name -> envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.HeadersWithUnderscoresAction
	1, // 3: envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.UriPathNormalizationOptions.path_with_escaped_slashes_action:type_name -> envoy.extensions.http.header_validators.envoy_default.v3.HeaderValidatorConfig.UriPathNormalizationOptions.PathWithEscapedSlashesAction
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_init()
}
func file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_init() {
	if File_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValidatorConfig); i {
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
		file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValidatorConfig_UriPathNormalizationOptions); i {
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
		file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValidatorConfig_Http1ProtocolOptions); i {
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
			RawDescriptor: file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_msgTypes,
	}.Build()
	File_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto = out.File
	file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_rawDesc = nil
	file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_goTypes = nil
	file_envoy_extensions_http_header_validators_envoy_default_v3_header_validator_proto_depIdxs = nil
}

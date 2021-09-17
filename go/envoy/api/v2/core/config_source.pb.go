// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: envoy/api/v2/core/config_source.proto

package envoy_api_v2_core

import (
	_ "github.com/cilium/proxy/go/envoy/annotations"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// xDS API version. This is used to describe both resource and transport
// protocol versions (in distinct configuration fields).
type ApiVersion int32

const (
	// When not specified, we assume v2, to ease migration to Envoy's stable API
	// versioning. If a client does not support v2 (e.g. due to deprecation), this
	// is an invalid value.
	//
	// Deprecated: Do not use.
	ApiVersion_AUTO ApiVersion = 0
	// Use xDS v2 API.
	//
	// Deprecated: Do not use.
	ApiVersion_V2 ApiVersion = 1
	// Use xDS v3 API.
	ApiVersion_V3 ApiVersion = 2
)

// Enum value maps for ApiVersion.
var (
	ApiVersion_name = map[int32]string{
		0: "AUTO",
		1: "V2",
		2: "V3",
	}
	ApiVersion_value = map[string]int32{
		"AUTO": 0,
		"V2":   1,
		"V3":   2,
	}
)

func (x ApiVersion) Enum() *ApiVersion {
	p := new(ApiVersion)
	*p = x
	return p
}

func (x ApiVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_api_v2_core_config_source_proto_enumTypes[0].Descriptor()
}

func (ApiVersion) Type() protoreflect.EnumType {
	return &file_envoy_api_v2_core_config_source_proto_enumTypes[0]
}

func (x ApiVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiVersion.Descriptor instead.
func (ApiVersion) EnumDescriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{0}
}

// APIs may be fetched via either REST or gRPC.
type ApiConfigSource_ApiType int32

const (
	// Ideally this would be 'reserved 0' but one can't reserve the default
	// value. Instead we throw an exception if this is ever used.
	//
	// Deprecated: Do not use.
	ApiConfigSource_UNSUPPORTED_REST_LEGACY ApiConfigSource_ApiType = 0
	// REST-JSON v2 API. The `canonical JSON encoding
	// <https://developers.google.com/protocol-buffers/docs/proto3#json>`_ for
	// the v2 protos is used.
	ApiConfigSource_REST ApiConfigSource_ApiType = 1
	// gRPC v2 API.
	ApiConfigSource_GRPC ApiConfigSource_ApiType = 2
	// Using the delta xDS gRPC service, i.e. DeltaDiscovery{Request,Response}
	// rather than Discovery{Request,Response}. Rather than sending Envoy the entire state
	// with every update, the xDS server only sends what has changed since the last update.
	ApiConfigSource_DELTA_GRPC ApiConfigSource_ApiType = 3
)

// Enum value maps for ApiConfigSource_ApiType.
var (
	ApiConfigSource_ApiType_name = map[int32]string{
		0: "UNSUPPORTED_REST_LEGACY",
		1: "REST",
		2: "GRPC",
		3: "DELTA_GRPC",
	}
	ApiConfigSource_ApiType_value = map[string]int32{
		"UNSUPPORTED_REST_LEGACY": 0,
		"REST":                    1,
		"GRPC":                    2,
		"DELTA_GRPC":              3,
	}
)

func (x ApiConfigSource_ApiType) Enum() *ApiConfigSource_ApiType {
	p := new(ApiConfigSource_ApiType)
	*p = x
	return p
}

func (x ApiConfigSource_ApiType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiConfigSource_ApiType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_api_v2_core_config_source_proto_enumTypes[1].Descriptor()
}

func (ApiConfigSource_ApiType) Type() protoreflect.EnumType {
	return &file_envoy_api_v2_core_config_source_proto_enumTypes[1]
}

func (x ApiConfigSource_ApiType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiConfigSource_ApiType.Descriptor instead.
func (ApiConfigSource_ApiType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{0, 0}
}

// API configuration source. This identifies the API type and cluster that Envoy
// will use to fetch an xDS API.
// [#next-free-field: 9]
type ApiConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API type (gRPC, REST, delta gRPC)
	ApiType ApiConfigSource_ApiType `protobuf:"varint,1,opt,name=api_type,json=apiType,proto3,enum=envoy.api.v2.core.ApiConfigSource_ApiType" json:"api_type,omitempty"`
	// API version for xDS transport protocol. This describes the xDS gRPC/REST
	// endpoint and version of [Delta]DiscoveryRequest/Response used on the wire.
	TransportApiVersion ApiVersion `protobuf:"varint,8,opt,name=transport_api_version,json=transportApiVersion,proto3,enum=envoy.api.v2.core.ApiVersion" json:"transport_api_version,omitempty"`
	// Cluster names should be used only with REST. If > 1
	// cluster is defined, clusters will be cycled through if any kind of failure
	// occurs.
	//
	// .. note::
	//
	//  The cluster with name ``cluster_name`` must be statically defined and its
	//  type must not be ``EDS``.
	ClusterNames []string `protobuf:"bytes,2,rep,name=cluster_names,json=clusterNames,proto3" json:"cluster_names,omitempty"`
	// Multiple gRPC services be provided for GRPC. If > 1 cluster is defined,
	// services will be cycled through if any kind of failure occurs.
	GrpcServices []*GrpcService `protobuf:"bytes,4,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	// For REST APIs, the delay between successive polls.
	RefreshDelay *durationpb.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	// For REST APIs, the request timeout. If not set, a default value of 1s will be used.
	RequestTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// For GRPC APIs, the rate limit settings. If present, discovery requests made by Envoy will be
	// rate limited.
	RateLimitSettings *RateLimitSettings `protobuf:"bytes,6,opt,name=rate_limit_settings,json=rateLimitSettings,proto3" json:"rate_limit_settings,omitempty"`
	// Skip the node identifier in subsequent discovery requests for streaming gRPC config types.
	SetNodeOnFirstMessageOnly bool `protobuf:"varint,7,opt,name=set_node_on_first_message_only,json=setNodeOnFirstMessageOnly,proto3" json:"set_node_on_first_message_only,omitempty"`
}

func (x *ApiConfigSource) Reset() {
	*x = ApiConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiConfigSource) ProtoMessage() {}

func (x *ApiConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiConfigSource.ProtoReflect.Descriptor instead.
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{0}
}

func (x *ApiConfigSource) GetApiType() ApiConfigSource_ApiType {
	if x != nil {
		return x.ApiType
	}
	return ApiConfigSource_UNSUPPORTED_REST_LEGACY
}

func (x *ApiConfigSource) GetTransportApiVersion() ApiVersion {
	if x != nil {
		return x.TransportApiVersion
	}
	return ApiVersion_AUTO
}

func (x *ApiConfigSource) GetClusterNames() []string {
	if x != nil {
		return x.ClusterNames
	}
	return nil
}

func (x *ApiConfigSource) GetGrpcServices() []*GrpcService {
	if x != nil {
		return x.GrpcServices
	}
	return nil
}

func (x *ApiConfigSource) GetRefreshDelay() *durationpb.Duration {
	if x != nil {
		return x.RefreshDelay
	}
	return nil
}

func (x *ApiConfigSource) GetRequestTimeout() *durationpb.Duration {
	if x != nil {
		return x.RequestTimeout
	}
	return nil
}

func (x *ApiConfigSource) GetRateLimitSettings() *RateLimitSettings {
	if x != nil {
		return x.RateLimitSettings
	}
	return nil
}

func (x *ApiConfigSource) GetSetNodeOnFirstMessageOnly() bool {
	if x != nil {
		return x.SetNodeOnFirstMessageOnly
	}
	return false
}

// Aggregated Discovery Service (ADS) options. This is currently empty, but when
// set in :ref:`ConfigSource <envoy_api_msg_core.ConfigSource>` can be used to
// specify that ADS is to be used.
type AggregatedConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AggregatedConfigSource) Reset() {
	*x = AggregatedConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregatedConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregatedConfigSource) ProtoMessage() {}

func (x *AggregatedConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregatedConfigSource.ProtoReflect.Descriptor instead.
func (*AggregatedConfigSource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{1}
}

// [#not-implemented-hide:]
// Self-referencing config source options. This is currently empty, but when
// set in :ref:`ConfigSource <envoy_api_msg_core.ConfigSource>` can be used to
// specify that other data can be obtained from the same server.
type SelfConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API version for xDS transport protocol. This describes the xDS gRPC/REST
	// endpoint and version of [Delta]DiscoveryRequest/Response used on the wire.
	TransportApiVersion ApiVersion `protobuf:"varint,1,opt,name=transport_api_version,json=transportApiVersion,proto3,enum=envoy.api.v2.core.ApiVersion" json:"transport_api_version,omitempty"`
}

func (x *SelfConfigSource) Reset() {
	*x = SelfConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfConfigSource) ProtoMessage() {}

func (x *SelfConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfConfigSource.ProtoReflect.Descriptor instead.
func (*SelfConfigSource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{2}
}

func (x *SelfConfigSource) GetTransportApiVersion() ApiVersion {
	if x != nil {
		return x.TransportApiVersion
	}
	return ApiVersion_AUTO
}

// Rate Limit settings to be applied for discovery requests made by Envoy.
type RateLimitSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maximum number of tokens to be used for rate limiting discovery request calls. If not set, a
	// default value of 100 will be used.
	MaxTokens *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// Rate at which tokens will be filled per second. If not set, a default fill rate of 10 tokens
	// per second will be used.
	FillRate *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=fill_rate,json=fillRate,proto3" json:"fill_rate,omitempty"`
}

func (x *RateLimitSettings) Reset() {
	*x = RateLimitSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitSettings) ProtoMessage() {}

func (x *RateLimitSettings) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitSettings.ProtoReflect.Descriptor instead.
func (*RateLimitSettings) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{3}
}

func (x *RateLimitSettings) GetMaxTokens() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxTokens
	}
	return nil
}

func (x *RateLimitSettings) GetFillRate() *wrapperspb.DoubleValue {
	if x != nil {
		return x.FillRate
	}
	return nil
}

// Configuration for :ref:`listeners <config_listeners>`, :ref:`clusters
// <config_cluster_manager>`, :ref:`routes
// <envoy_api_msg_RouteConfiguration>`, :ref:`endpoints
// <arch_overview_service_discovery>` etc. may either be sourced from the
// filesystem or from an xDS API source. Filesystem configs are watched with
// inotify for updates.
// [#next-free-field: 7]
type ConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigSourceSpecifier:
	//	*ConfigSource_Path
	//	*ConfigSource_ApiConfigSource
	//	*ConfigSource_Ads
	//	*ConfigSource_Self
	ConfigSourceSpecifier isConfigSource_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	// When this timeout is specified, Envoy will wait no longer than the specified time for first
	// config response on this xDS subscription during the :ref:`initialization process
	// <arch_overview_initialization>`. After reaching the timeout, Envoy will move to the next
	// initialization phase, even if the first config is not delivered yet. The timer is activated
	// when the xDS API subscription starts, and is disarmed on first config update or on error. 0
	// means no timeout - Envoy will wait indefinitely for the first xDS config (unless another
	// timeout applies). The default is 15s.
	InitialFetchTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=initial_fetch_timeout,json=initialFetchTimeout,proto3" json:"initial_fetch_timeout,omitempty"`
	// API version for xDS resources. This implies the type URLs that the client
	// will request for resources and the resource type that the client will in
	// turn expect to be delivered.
	ResourceApiVersion ApiVersion `protobuf:"varint,6,opt,name=resource_api_version,json=resourceApiVersion,proto3,enum=envoy.api.v2.core.ApiVersion" json:"resource_api_version,omitempty"`
}

func (x *ConfigSource) Reset() {
	*x = ConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSource) ProtoMessage() {}

func (x *ConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSource.ProtoReflect.Descriptor instead.
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{4}
}

func (m *ConfigSource) GetConfigSourceSpecifier() isConfigSource_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (x *ConfigSource) GetPath() string {
	if x, ok := x.GetConfigSourceSpecifier().(*ConfigSource_Path); ok {
		return x.Path
	}
	return ""
}

func (x *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if x, ok := x.GetConfigSourceSpecifier().(*ConfigSource_ApiConfigSource); ok {
		return x.ApiConfigSource
	}
	return nil
}

func (x *ConfigSource) GetAds() *AggregatedConfigSource {
	if x, ok := x.GetConfigSourceSpecifier().(*ConfigSource_Ads); ok {
		return x.Ads
	}
	return nil
}

func (x *ConfigSource) GetSelf() *SelfConfigSource {
	if x, ok := x.GetConfigSourceSpecifier().(*ConfigSource_Self); ok {
		return x.Self
	}
	return nil
}

func (x *ConfigSource) GetInitialFetchTimeout() *durationpb.Duration {
	if x != nil {
		return x.InitialFetchTimeout
	}
	return nil
}

func (x *ConfigSource) GetResourceApiVersion() ApiVersion {
	if x != nil {
		return x.ResourceApiVersion
	}
	return ApiVersion_AUTO
}

type isConfigSource_ConfigSourceSpecifier interface {
	isConfigSource_ConfigSourceSpecifier()
}

type ConfigSource_Path struct {
	// Path on the filesystem to source and watch for configuration updates.
	// When sourcing configuration for :ref:`secret <envoy_api_msg_auth.Secret>`,
	// the certificate and key files are also watched for updates.
	//
	// .. note::
	//
	//  The path to the source must exist at config load time.
	//
	// .. note::
	//
	//   Envoy will only watch the file path for *moves.* This is because in general only moves
	//   are atomic. The same method of swapping files as is demonstrated in the
	//   :ref:`runtime documentation <config_runtime_symbolic_link_swap>` can be used here also.
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ConfigSource_ApiConfigSource struct {
	// API configuration source.
	ApiConfigSource *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,proto3,oneof"`
}

type ConfigSource_Ads struct {
	// When set, ADS will be used to fetch resources. The ADS API configuration
	// source in the bootstrap configuration is used.
	Ads *AggregatedConfigSource `protobuf:"bytes,3,opt,name=ads,proto3,oneof"`
}

type ConfigSource_Self struct {
	// [#not-implemented-hide:]
	// When set, the client will access the resources from the same server it got the
	// ConfigSource from, although not necessarily from the same stream. This is similar to the
	// :ref:`ads<envoy_api_field.ConfigSource.ads>` field, except that the client may use a
	// different stream to the same server. As a result, this field can be used for things
	// like LRS that cannot be sent on an ADS stream. It can also be used to link from (e.g.)
	// LDS to RDS on the same server without requiring the management server to know its name
	// or required credentials.
	// [#next-major-version: In xDS v3, consider replacing the ads field with this one, since
	// this field can implicitly mean to use the same stream in the case where the ConfigSource
	// is provided via ADS and the specified data can also be obtained via ADS.]
	Self *SelfConfigSource `protobuf:"bytes,5,opt,name=self,proto3,oneof"`
}

func (*ConfigSource_Path) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_ApiConfigSource) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Ads) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Self) isConfigSource_ConfigSourceSpecifier() {}

var File_envoy_api_v2_core_config_source_proto protoreflect.FileDescriptor

var file_envoy_api_v2_core_config_source_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x05,
	0x0a, 0x0f, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x70, 0x69, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x5b, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x4c, 0x0a, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x54, 0x0a, 0x13, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x11, 0x72, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x41, 0x0a,
	0x1e, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x73, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x6e,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0x54, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x17, 0x55,
	0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x5f,
	0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x00, 0x1a, 0x08, 0x08, 0x01, 0xa8, 0xf7, 0xb4, 0x8b,
	0x02, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x47, 0x52, 0x50, 0x43, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x4c, 0x54, 0x41, 0x5f,
	0x47, 0x52, 0x50, 0x43, 0x10, 0x03, 0x22, 0x18, 0x0a, 0x16, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x6f, 0x0a, 0x10, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x9b, 0x01, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x21, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x22,
	0xba, 0x03, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x50, 0x0a, 0x11, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x61, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x48, 0x00, 0x52, 0x03, 0x61, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x04, 0x73, 0x65,
	0x6c, 0x66, 0x12, 0x4d, 0x0a, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x59, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x1e, 0x0a, 0x17,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x2a, 0x2e, 0x0a, 0x0a,
	0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x04, 0x41, 0x55,
	0x54, 0x4f, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0a, 0x0a, 0x02, 0x56, 0x32, 0x10, 0x01,
	0x1a, 0x02, 0x08, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x33, 0x10, 0x02, 0x42, 0x5a, 0x0a, 0x1f,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42,
	0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x16, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_api_v2_core_config_source_proto_rawDescOnce sync.Once
	file_envoy_api_v2_core_config_source_proto_rawDescData = file_envoy_api_v2_core_config_source_proto_rawDesc
)

func file_envoy_api_v2_core_config_source_proto_rawDescGZIP() []byte {
	file_envoy_api_v2_core_config_source_proto_rawDescOnce.Do(func() {
		file_envoy_api_v2_core_config_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_api_v2_core_config_source_proto_rawDescData)
	})
	return file_envoy_api_v2_core_config_source_proto_rawDescData
}

var file_envoy_api_v2_core_config_source_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_api_v2_core_config_source_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_api_v2_core_config_source_proto_goTypes = []interface{}{
	(ApiVersion)(0),                // 0: envoy.api.v2.core.ApiVersion
	(ApiConfigSource_ApiType)(0),   // 1: envoy.api.v2.core.ApiConfigSource.ApiType
	(*ApiConfigSource)(nil),        // 2: envoy.api.v2.core.ApiConfigSource
	(*AggregatedConfigSource)(nil), // 3: envoy.api.v2.core.AggregatedConfigSource
	(*SelfConfigSource)(nil),       // 4: envoy.api.v2.core.SelfConfigSource
	(*RateLimitSettings)(nil),      // 5: envoy.api.v2.core.RateLimitSettings
	(*ConfigSource)(nil),           // 6: envoy.api.v2.core.ConfigSource
	(*GrpcService)(nil),            // 7: envoy.api.v2.core.GrpcService
	(*durationpb.Duration)(nil),    // 8: google.protobuf.Duration
	(*wrapperspb.UInt32Value)(nil), // 9: google.protobuf.UInt32Value
	(*wrapperspb.DoubleValue)(nil), // 10: google.protobuf.DoubleValue
}
var file_envoy_api_v2_core_config_source_proto_depIdxs = []int32{
	1,  // 0: envoy.api.v2.core.ApiConfigSource.api_type:type_name -> envoy.api.v2.core.ApiConfigSource.ApiType
	0,  // 1: envoy.api.v2.core.ApiConfigSource.transport_api_version:type_name -> envoy.api.v2.core.ApiVersion
	7,  // 2: envoy.api.v2.core.ApiConfigSource.grpc_services:type_name -> envoy.api.v2.core.GrpcService
	8,  // 3: envoy.api.v2.core.ApiConfigSource.refresh_delay:type_name -> google.protobuf.Duration
	8,  // 4: envoy.api.v2.core.ApiConfigSource.request_timeout:type_name -> google.protobuf.Duration
	5,  // 5: envoy.api.v2.core.ApiConfigSource.rate_limit_settings:type_name -> envoy.api.v2.core.RateLimitSettings
	0,  // 6: envoy.api.v2.core.SelfConfigSource.transport_api_version:type_name -> envoy.api.v2.core.ApiVersion
	9,  // 7: envoy.api.v2.core.RateLimitSettings.max_tokens:type_name -> google.protobuf.UInt32Value
	10, // 8: envoy.api.v2.core.RateLimitSettings.fill_rate:type_name -> google.protobuf.DoubleValue
	2,  // 9: envoy.api.v2.core.ConfigSource.api_config_source:type_name -> envoy.api.v2.core.ApiConfigSource
	3,  // 10: envoy.api.v2.core.ConfigSource.ads:type_name -> envoy.api.v2.core.AggregatedConfigSource
	4,  // 11: envoy.api.v2.core.ConfigSource.self:type_name -> envoy.api.v2.core.SelfConfigSource
	8,  // 12: envoy.api.v2.core.ConfigSource.initial_fetch_timeout:type_name -> google.protobuf.Duration
	0,  // 13: envoy.api.v2.core.ConfigSource.resource_api_version:type_name -> envoy.api.v2.core.ApiVersion
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_envoy_api_v2_core_config_source_proto_init() }
func file_envoy_api_v2_core_config_source_proto_init() {
	if File_envoy_api_v2_core_config_source_proto != nil {
		return
	}
	file_envoy_api_v2_core_grpc_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_api_v2_core_config_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiConfigSource); i {
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
		file_envoy_api_v2_core_config_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregatedConfigSource); i {
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
		file_envoy_api_v2_core_config_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfConfigSource); i {
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
		file_envoy_api_v2_core_config_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitSettings); i {
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
		file_envoy_api_v2_core_config_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSource); i {
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
	file_envoy_api_v2_core_config_source_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ConfigSource_Path)(nil),
		(*ConfigSource_ApiConfigSource)(nil),
		(*ConfigSource_Ads)(nil),
		(*ConfigSource_Self)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_api_v2_core_config_source_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_api_v2_core_config_source_proto_goTypes,
		DependencyIndexes: file_envoy_api_v2_core_config_source_proto_depIdxs,
		EnumInfos:         file_envoy_api_v2_core_config_source_proto_enumTypes,
		MessageInfos:      file_envoy_api_v2_core_config_source_proto_msgTypes,
	}.Build()
	File_envoy_api_v2_core_config_source_proto = out.File
	file_envoy_api_v2_core_config_source_proto_rawDesc = nil
	file_envoy_api_v2_core_config_source_proto_goTypes = nil
	file_envoy_api_v2_core_config_source_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: envoy/extensions/filters/http/ratelimit/v3/rate_limit.proto

package envoy_extensions_filters_http_ratelimit_v3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/ratelimit/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// Defines the version of the standard to use for X-RateLimit headers.
type RateLimit_XRateLimitHeadersRFCVersion int32

const (
	// X-RateLimit headers disabled.
	RateLimit_OFF RateLimit_XRateLimitHeadersRFCVersion = 0
	// Use `draft RFC Version 03 <https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html>`_.
	RateLimit_DRAFT_VERSION_03 RateLimit_XRateLimitHeadersRFCVersion = 1
)

// Enum value maps for RateLimit_XRateLimitHeadersRFCVersion.
var (
	RateLimit_XRateLimitHeadersRFCVersion_name = map[int32]string{
		0: "OFF",
		1: "DRAFT_VERSION_03",
	}
	RateLimit_XRateLimitHeadersRFCVersion_value = map[string]int32{
		"OFF":              0,
		"DRAFT_VERSION_03": 1,
	}
)

func (x RateLimit_XRateLimitHeadersRFCVersion) Enum() *RateLimit_XRateLimitHeadersRFCVersion {
	p := new(RateLimit_XRateLimitHeadersRFCVersion)
	*p = x
	return p
}

func (x RateLimit_XRateLimitHeadersRFCVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimit_XRateLimitHeadersRFCVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_enumTypes[0].Descriptor()
}

func (RateLimit_XRateLimitHeadersRFCVersion) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_enumTypes[0]
}

func (x RateLimit_XRateLimitHeadersRFCVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimit_XRateLimitHeadersRFCVersion.Descriptor instead.
func (RateLimit_XRateLimitHeadersRFCVersion) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescGZIP(), []int{0, 0}
}

type RateLimitPerRoute_VhRateLimitsOptions int32

const (
	// Use the virtual host rate limits unless the route has a rate limit policy.
	RateLimitPerRoute_OVERRIDE RateLimitPerRoute_VhRateLimitsOptions = 0
	// Use the virtual host rate limits even if the route has a rate limit policy.
	RateLimitPerRoute_INCLUDE RateLimitPerRoute_VhRateLimitsOptions = 1
	// Ignore the virtual host rate limits even if the route does not have a rate limit policy.
	RateLimitPerRoute_IGNORE RateLimitPerRoute_VhRateLimitsOptions = 2
)

// Enum value maps for RateLimitPerRoute_VhRateLimitsOptions.
var (
	RateLimitPerRoute_VhRateLimitsOptions_name = map[int32]string{
		0: "OVERRIDE",
		1: "INCLUDE",
		2: "IGNORE",
	}
	RateLimitPerRoute_VhRateLimitsOptions_value = map[string]int32{
		"OVERRIDE": 0,
		"INCLUDE":  1,
		"IGNORE":   2,
	}
)

func (x RateLimitPerRoute_VhRateLimitsOptions) Enum() *RateLimitPerRoute_VhRateLimitsOptions {
	p := new(RateLimitPerRoute_VhRateLimitsOptions)
	*p = x
	return p
}

func (x RateLimitPerRoute_VhRateLimitsOptions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitPerRoute_VhRateLimitsOptions) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_enumTypes[1].Descriptor()
}

func (RateLimitPerRoute_VhRateLimitsOptions) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_enumTypes[1]
}

func (x RateLimitPerRoute_VhRateLimitsOptions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitPerRoute_VhRateLimitsOptions.Descriptor instead.
func (RateLimitPerRoute_VhRateLimitsOptions) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescGZIP(), []int{1, 0}
}

// [#next-free-field: 10]
type RateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rate limit domain to use when calling the rate limit service.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configurations to be applied with the same
	// stage number. If not set, the default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The type of requests the filter should apply to. The supported
	// types are *internal*, *external* or *both*. A request is considered internal if
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is set to true. If
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is not set or false, a
	// request is considered external. The filter defaults to *both*, and it will apply to all request
	// types.
	RequestType string `protobuf:"bytes,3,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *durationpb.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Specifies whether a `RESOURCE_EXHAUSTED` gRPC code must be returned instead
	// of the default `UNAVAILABLE` gRPC code for a rate limited gRPC call. The
	// HTTP code will be 200 for a gRPC response.
	RateLimitedAsResourceExhausted bool `protobuf:"varint,6,opt,name=rate_limited_as_resource_exhausted,json=rateLimitedAsResourceExhausted,proto3" json:"rate_limited_as_resource_exhausted,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService *v3.RateLimitServiceConfig `protobuf:"bytes,7,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	// Defines the standard version to use for X-RateLimit headers emitted by the filter:
	//
	// * ``X-RateLimit-Limit`` - indicates the request-quota associated to the
	//   client in the current time-window followed by the description of the
	//   quota policy. The values are returned by the rate limiting service in
	//   :ref:`current_limit<envoy_v3_api_field_service.ratelimit.v3.RateLimitResponse.DescriptorStatus.current_limit>`
	//   field. Example: `10, 10;w=1;name="per-ip", 1000;w=3600`.
	// * ``X-RateLimit-Remaining`` - indicates the remaining requests in the
	//   current time-window. The values are returned by the rate limiting service
	//   in :ref:`limit_remaining<envoy_v3_api_field_service.ratelimit.v3.RateLimitResponse.DescriptorStatus.limit_remaining>`
	//   field.
	// * ``X-RateLimit-Reset`` - indicates the number of seconds until reset of
	//   the current time-window. The values are returned by the rate limiting service
	//   in :ref:`duration_until_reset<envoy_v3_api_field_service.ratelimit.v3.RateLimitResponse.DescriptorStatus.duration_until_reset>`
	//   field.
	//
	// In case rate limiting policy specifies more then one time window, the values
	// above represent the window that is closest to reaching its limit.
	//
	// For more information about the headers specification see selected version of
	// the `draft RFC <https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html>`_.
	//
	// Disabled by default.
	EnableXRatelimitHeaders RateLimit_XRateLimitHeadersRFCVersion `protobuf:"varint,8,opt,name=enable_x_ratelimit_headers,json=enableXRatelimitHeaders,proto3,enum=envoy.extensions.filters.http.ratelimit.v3.RateLimit_XRateLimitHeadersRFCVersion" json:"enable_x_ratelimit_headers,omitempty"`
	// Disables emitting the :ref:`x-envoy-ratelimited<config_http_filters_router_x-envoy-ratelimited>` header
	// in case of rate limiting (i.e. 429 responses).
	// Having this header not present potentially makes the request retriable.
	DisableXEnvoyRatelimitedHeader bool `protobuf:"varint,9,opt,name=disable_x_envoy_ratelimited_header,json=disableXEnvoyRatelimitedHeader,proto3" json:"disable_x_envoy_ratelimited_header,omitempty"`
}

func (x *RateLimit) Reset() {
	*x = RateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimit) ProtoMessage() {}

func (x *RateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimit.ProtoReflect.Descriptor instead.
func (*RateLimit) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimit) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RateLimit) GetStage() uint32 {
	if x != nil {
		return x.Stage
	}
	return 0
}

func (x *RateLimit) GetRequestType() string {
	if x != nil {
		return x.RequestType
	}
	return ""
}

func (x *RateLimit) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RateLimit) GetFailureModeDeny() bool {
	if x != nil {
		return x.FailureModeDeny
	}
	return false
}

func (x *RateLimit) GetRateLimitedAsResourceExhausted() bool {
	if x != nil {
		return x.RateLimitedAsResourceExhausted
	}
	return false
}

func (x *RateLimit) GetRateLimitService() *v3.RateLimitServiceConfig {
	if x != nil {
		return x.RateLimitService
	}
	return nil
}

func (x *RateLimit) GetEnableXRatelimitHeaders() RateLimit_XRateLimitHeadersRFCVersion {
	if x != nil {
		return x.EnableXRatelimitHeaders
	}
	return RateLimit_OFF
}

func (x *RateLimit) GetDisableXEnvoyRatelimitedHeader() bool {
	if x != nil {
		return x.DisableXEnvoyRatelimitedHeader
	}
	return false
}

type RateLimitPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies if the rate limit filter should include the virtual host rate limits.
	VhRateLimits RateLimitPerRoute_VhRateLimitsOptions `protobuf:"varint,1,opt,name=vh_rate_limits,json=vhRateLimits,proto3,enum=envoy.extensions.filters.http.ratelimit.v3.RateLimitPerRoute_VhRateLimitsOptions" json:"vh_rate_limits,omitempty"`
}

func (x *RateLimitPerRoute) Reset() {
	*x = RateLimitPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitPerRoute) ProtoMessage() {}

func (x *RateLimitPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitPerRoute.ProtoReflect.Descriptor instead.
func (*RateLimitPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimitPerRoute) GetVhRateLimits() RateLimitPerRoute_VhRateLimitsOptions {
	if x != nil {
		return x.VhRateLimits
	}
	return RateLimitPerRoute_OVERRIDE
}

var File_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x06, 0x0a, 0x09, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x0a,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xfa,
	0x42, 0x1e, 0x72, 0x1c, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x08,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x04, 0x62, 0x6f, 0x74, 0x68, 0x52, 0x00,
	0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x5f, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x6e, 0x79, 0x12, 0x4a,
	0x0a, 0x22, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x68, 0x61, 0x75,
	0x73, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x72, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x41, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x45, 0x78, 0x68, 0x61, 0x75, 0x73, 0x74, 0x65, 0x64, 0x12, 0x69, 0x0a, 0x12, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x10, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x51, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x58, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x46, 0x43, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x17, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x58,
	0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x4a, 0x0a, 0x22, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x78, 0x5f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x58, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x52, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x1b,
	0x58, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x46, 0x43, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x4f,
	0x46, 0x46, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x52, 0x41, 0x46, 0x54, 0x5f, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x30, 0x33, 0x10, 0x01, 0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e,
	0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0xd5, 0x01, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x76, 0x68,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x51, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x56, 0x68, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x0c, 0x76, 0x68, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x3c, 0x0a,
	0x13, 0x56, 0x68, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x56, 0x45, 0x52, 0x52, 0x49, 0x44, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45, 0x10, 0x02, 0x42, 0x54, 0x0a, 0x38, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescData = file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDesc
)

func file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDescData
}

var file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_goTypes = []interface{}{
	(RateLimit_XRateLimitHeadersRFCVersion)(0), // 0: envoy.extensions.filters.http.ratelimit.v3.RateLimit.XRateLimitHeadersRFCVersion
	(RateLimitPerRoute_VhRateLimitsOptions)(0), // 1: envoy.extensions.filters.http.ratelimit.v3.RateLimitPerRoute.VhRateLimitsOptions
	(*RateLimit)(nil),                          // 2: envoy.extensions.filters.http.ratelimit.v3.RateLimit
	(*RateLimitPerRoute)(nil),                  // 3: envoy.extensions.filters.http.ratelimit.v3.RateLimitPerRoute
	(*durationpb.Duration)(nil),                // 4: google.protobuf.Duration
	(*v3.RateLimitServiceConfig)(nil),          // 5: envoy.config.ratelimit.v3.RateLimitServiceConfig
}
var file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.filters.http.ratelimit.v3.RateLimit.timeout:type_name -> google.protobuf.Duration
	5, // 1: envoy.extensions.filters.http.ratelimit.v3.RateLimit.rate_limit_service:type_name -> envoy.config.ratelimit.v3.RateLimitServiceConfig
	0, // 2: envoy.extensions.filters.http.ratelimit.v3.RateLimit.enable_x_ratelimit_headers:type_name -> envoy.extensions.filters.http.ratelimit.v3.RateLimit.XRateLimitHeadersRFCVersion
	1, // 3: envoy.extensions.filters.http.ratelimit.v3.RateLimitPerRoute.vh_rate_limits:type_name -> envoy.extensions.filters.http.ratelimit.v3.RateLimitPerRoute.VhRateLimitsOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_init() }
func file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_init() {
	if File_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimit); i {
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
		file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitPerRoute); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto = out.File
	file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_rawDesc = nil
	file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_goTypes = nil
	file_envoy_extensions_filters_http_ratelimit_v3_rate_limit_proto_depIdxs = nil
}

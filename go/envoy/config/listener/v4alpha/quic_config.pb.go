// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: envoy/config/listener/v4alpha/quic_config.proto

package envoy_config_listener_v4alpha

import (
	v4alpha "github.com/cilium/proxy/go/envoy/config/core/v4alpha"
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// Configuration specific to the UDP QUIC listener.
type QuicProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuicProtocolOptions *v4alpha.QuicProtocolOptions `protobuf:"bytes,1,opt,name=quic_protocol_options,json=quicProtocolOptions,proto3" json:"quic_protocol_options,omitempty"`
	// Maximum number of milliseconds that connection will be alive when there is
	// no network activity. 300000ms if not specified.
	IdleTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// Connection timeout in milliseconds before the crypto handshake is finished.
	// 20000ms if not specified.
	CryptoHandshakeTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
	// Runtime flag that controls whether the listener is enabled or not. If not specified, defaults
	// to enabled.
	Enabled *v4alpha.RuntimeFeatureFlag `protobuf:"bytes,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *QuicProtocolOptions) Reset() {
	*x = QuicProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v4alpha_quic_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuicProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuicProtocolOptions) ProtoMessage() {}

func (x *QuicProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v4alpha_quic_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuicProtocolOptions.ProtoReflect.Descriptor instead.
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v4alpha_quic_config_proto_rawDescGZIP(), []int{0}
}

func (x *QuicProtocolOptions) GetQuicProtocolOptions() *v4alpha.QuicProtocolOptions {
	if x != nil {
		return x.QuicProtocolOptions
	}
	return nil
}

func (x *QuicProtocolOptions) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *QuicProtocolOptions) GetCryptoHandshakeTimeout() *durationpb.Duration {
	if x != nil {
		return x.CryptoHandshakeTimeout
	}
	return nil
}

func (x *QuicProtocolOptions) GetEnabled() *v4alpha.RuntimeFeatureFlag {
	if x != nil {
		return x.Enabled
	}
	return nil
}

var File_envoy_config_listener_v4alpha_quic_config_proto protoreflect.FileDescriptor

var file_envoy_config_listener_v4alpha_quic_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x71, 0x75, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x13, 0x51, 0x75, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x62, 0x0a, 0x15, 0x71, 0x75,
	0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x51, 0x75, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x71, 0x75, 0x69, 0x63, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c,
	0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x53, 0x0a, 0x18,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x47, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61,
	0x67, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x3a, 0x33, 0x9a, 0xc5, 0x88, 0x1e,
	0x2e, 0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x75, 0x69, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x48, 0x0a, 0x2b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0f,
	0x51, 0x75, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_config_listener_v4alpha_quic_config_proto_rawDescOnce sync.Once
	file_envoy_config_listener_v4alpha_quic_config_proto_rawDescData = file_envoy_config_listener_v4alpha_quic_config_proto_rawDesc
)

func file_envoy_config_listener_v4alpha_quic_config_proto_rawDescGZIP() []byte {
	file_envoy_config_listener_v4alpha_quic_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_listener_v4alpha_quic_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_listener_v4alpha_quic_config_proto_rawDescData)
	})
	return file_envoy_config_listener_v4alpha_quic_config_proto_rawDescData
}

var file_envoy_config_listener_v4alpha_quic_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_listener_v4alpha_quic_config_proto_goTypes = []interface{}{
	(*QuicProtocolOptions)(nil),         // 0: envoy.config.listener.v4alpha.QuicProtocolOptions
	(*v4alpha.QuicProtocolOptions)(nil), // 1: envoy.config.core.v4alpha.QuicProtocolOptions
	(*durationpb.Duration)(nil),         // 2: google.protobuf.Duration
	(*v4alpha.RuntimeFeatureFlag)(nil),  // 3: envoy.config.core.v4alpha.RuntimeFeatureFlag
}
var file_envoy_config_listener_v4alpha_quic_config_proto_depIdxs = []int32{
	1, // 0: envoy.config.listener.v4alpha.QuicProtocolOptions.quic_protocol_options:type_name -> envoy.config.core.v4alpha.QuicProtocolOptions
	2, // 1: envoy.config.listener.v4alpha.QuicProtocolOptions.idle_timeout:type_name -> google.protobuf.Duration
	2, // 2: envoy.config.listener.v4alpha.QuicProtocolOptions.crypto_handshake_timeout:type_name -> google.protobuf.Duration
	3, // 3: envoy.config.listener.v4alpha.QuicProtocolOptions.enabled:type_name -> envoy.config.core.v4alpha.RuntimeFeatureFlag
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_config_listener_v4alpha_quic_config_proto_init() }
func file_envoy_config_listener_v4alpha_quic_config_proto_init() {
	if File_envoy_config_listener_v4alpha_quic_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_listener_v4alpha_quic_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuicProtocolOptions); i {
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
			RawDescriptor: file_envoy_config_listener_v4alpha_quic_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_listener_v4alpha_quic_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_listener_v4alpha_quic_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_listener_v4alpha_quic_config_proto_msgTypes,
	}.Build()
	File_envoy_config_listener_v4alpha_quic_config_proto = out.File
	file_envoy_config_listener_v4alpha_quic_config_proto_rawDesc = nil
	file_envoy_config_listener_v4alpha_quic_config_proto_goTypes = nil
	file_envoy_config_listener_v4alpha_quic_config_proto_depIdxs = nil
}

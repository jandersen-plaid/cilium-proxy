// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: envoy/extensions/transport_sockets/tls/v3/cert.proto

package envoy_extensions_transport_sockets_tls_v3

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_envoy_extensions_transport_sockets_tls_v3_cert_proto protoreflect.FileDescriptor

var file_envoy_extensions_transport_sockets_tls_v3_cert_proto_rawDesc = []byte{
	0x0a, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x65, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76,
	0x33, 0x1a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x46, 0x0a, 0x37, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76,
	0x33, 0x42, 0x09, 0x43, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x50, 0x00,
	0x50, 0x01, 0x50, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_envoy_extensions_transport_sockets_tls_v3_cert_proto_goTypes = []interface{}{}
var file_envoy_extensions_transport_sockets_tls_v3_cert_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_transport_sockets_tls_v3_cert_proto_init() }
func file_envoy_extensions_transport_sockets_tls_v3_cert_proto_init() {
	if File_envoy_extensions_transport_sockets_tls_v3_cert_proto != nil {
		return
	}
	file_envoy_extensions_transport_sockets_tls_v3_common_proto_init()
	file_envoy_extensions_transport_sockets_tls_v3_secret_proto_init()
	file_envoy_extensions_transport_sockets_tls_v3_tls_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_transport_sockets_tls_v3_cert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_transport_sockets_tls_v3_cert_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_transport_sockets_tls_v3_cert_proto_depIdxs,
	}.Build()
	File_envoy_extensions_transport_sockets_tls_v3_cert_proto = out.File
	file_envoy_extensions_transport_sockets_tls_v3_cert_proto_rawDesc = nil
	file_envoy_extensions_transport_sockets_tls_v3_cert_proto_goTypes = nil
	file_envoy_extensions_transport_sockets_tls_v3_cert_proto_depIdxs = nil
}

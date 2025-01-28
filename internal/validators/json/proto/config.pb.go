// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.27.5
// source: github.com/cloudprober/cloudprober/internal/validators/json/proto/config.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// JSON validator configuration.
type Validator struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// If jq filter is specified, validator passes only if applying jq_filter to
	// the probe output, e.g. HTTP API response, results in 'true' boolean.
	// See the following test file for some examples:
	// https://github.com/cloudprober/cloudprober/blob/master/validators/json/json_test.go
	JqFilter      string `protobuf:"bytes,1,opt,name=jq_filter,json=jqFilter,proto3" json:"jq_filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Validator) Reset() {
	*x = Validator{}
	mi := &file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *Validator) GetJqFilter() string {
	if x != nil {
		return x.JqFilter
	}
	return ""
}

var File_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDesc = string([]byte{
	0x0a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x28, 0x0a,
	0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x71,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a,
	0x71, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDescData []byte
)

func file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDesc), len(file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDesc)))
	})
	return file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_goTypes = []any{
	(*Validator)(nil), // 0: cloudprober.validators.json.Validator
}
var file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_init()
}
func file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDesc), len(file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_internal_validators_json_proto_config_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.27.5
// source: github.com/cloudprober/cloudprober/probes/testdata/testdata.proto

package testdata

import (
	proto "github.com/cloudprober/cloudprober/probes/proto"
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

type FancyProbe struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          *string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FancyProbe) Reset() {
	*x = FancyProbe{}
	mi := &file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FancyProbe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FancyProbe) ProtoMessage() {}

func (x *FancyProbe) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FancyProbe.ProtoReflect.Descriptor instead.
func (*FancyProbe) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescGZIP(), []int{0}
}

func (x *FancyProbe) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type AnotherFancyProbe struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          *string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnotherFancyProbe) Reset() {
	*x = AnotherFancyProbe{}
	mi := &file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnotherFancyProbe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnotherFancyProbe) ProtoMessage() {}

func (x *AnotherFancyProbe) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnotherFancyProbe.ProtoReflect.Descriptor instead.
func (*AnotherFancyProbe) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescGZIP(), []int{1}
}

func (x *AnotherFancyProbe) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*proto.ProbeDef)(nil),
		ExtensionType: (*FancyProbe)(nil),
		Field:         200,
		Name:          "cloudprober.probes.testdata.fancy_probe",
		Tag:           "bytes,200,opt,name=fancy_probe",
		Filename:      "github.com/cloudprober/cloudprober/probes/testdata/testdata.proto",
	},
	{
		ExtendedType:  (*proto.ProbeDef)(nil),
		ExtensionType: (*AnotherFancyProbe)(nil),
		Field:         201,
		Name:          "cloudprober.probes.testdata.another_fancy_probe",
		Tag:           "bytes,201,opt,name=another_fancy_probe",
		Filename:      "github.com/cloudprober/cloudprober/probes/testdata/testdata.proto",
	},
}

// Extension fields to proto.ProbeDef.
var (
	// optional cloudprober.probes.testdata.FancyProbe fancy_probe = 200;
	E_FancyProbe = &file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_extTypes[0]
	// optional cloudprober.probes.testdata.AnotherFancyProbe another_fancy_probe = 201;
	E_AnotherFancyProbe = &file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_extTypes[1]
)

var File_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20,
	0x0a, 0x0a, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x27, 0x0a, 0x11, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x61, 0x6e, 0x63, 0x79,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x67, 0x0a, 0x0b, 0x66, 0x61, 0x6e,
	0x63, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x44, 0x65, 0x66, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x61, 0x6e, 0x63,
	0x79, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x0a, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x3a, 0x7d, 0x0a, 0x13, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x61,
	0x6e, 0x63, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x44, 0x65, 0x66, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x11,
	0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
}

var (
	file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescData = file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_goTypes = []any{
	(*FancyProbe)(nil),        // 0: cloudprober.probes.testdata.FancyProbe
	(*AnotherFancyProbe)(nil), // 1: cloudprober.probes.testdata.AnotherFancyProbe
	(*proto.ProbeDef)(nil),    // 2: cloudprober.probes.ProbeDef
}
var file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_depIdxs = []int32{
	2, // 0: cloudprober.probes.testdata.fancy_probe:extendee -> cloudprober.probes.ProbeDef
	2, // 1: cloudprober.probes.testdata.another_fancy_probe:extendee -> cloudprober.probes.ProbeDef
	0, // 2: cloudprober.probes.testdata.fancy_probe:type_name -> cloudprober.probes.testdata.FancyProbe
	1, // 3: cloudprober.probes.testdata.another_fancy_probe:type_name -> cloudprober.probes.testdata.AnotherFancyProbe
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_init() }
func file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_init() {
	if File_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_msgTypes,
		ExtensionInfos:    file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_extTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto = out.File
	file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_probes_testdata_testdata_proto_depIdxs = nil
}

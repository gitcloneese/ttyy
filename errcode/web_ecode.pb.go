// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: errcode/web_ecode.proto

package v1

import (
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

type WebErrCode int32

const (
	WebErrCode_Successed                WebErrCode = 0     //成功
	WebErrCode_GetWarZoneListFailed     WebErrCode = 60001 //获取列表失败
	WebErrCode_WarZoneProhibitModify    WebErrCode = 60002 //战区信息禁止修改
	WebErrCode_AddWarZoneFailed         WebErrCode = 60003 //添加战区配置信息失败
	WebErrCode_GetServerListFailed      WebErrCode = 60004 //获取服务器列表失败
	WebErrCode_ServerInfoProhibitModify WebErrCode = 60005 //服务器信息禁止修改
	WebErrCode_AddServerInfoFailed      WebErrCode = 60006 //添加服务器配置信息失败
)

// Enum value maps for WebErrCode.
var (
	WebErrCode_name = map[int32]string{
		0:     "Successed",
		60001: "GetWarZoneListFailed",
		60002: "WarZoneProhibitModify",
		60003: "AddWarZoneFailed",
		60004: "GetServerListFailed",
		60005: "ServerInfoProhibitModify",
		60006: "AddServerInfoFailed",
	}
	WebErrCode_value = map[string]int32{
		"Successed":                0,
		"GetWarZoneListFailed":     60001,
		"WarZoneProhibitModify":    60002,
		"AddWarZoneFailed":         60003,
		"GetServerListFailed":      60004,
		"ServerInfoProhibitModify": 60005,
		"AddServerInfoFailed":      60006,
	}
)

func (x WebErrCode) Enum() *WebErrCode {
	p := new(WebErrCode)
	*p = x
	return p
}

func (x WebErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WebErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errcode_web_ecode_proto_enumTypes[0].Descriptor()
}

func (WebErrCode) Type() protoreflect.EnumType {
	return &file_errcode_web_ecode_proto_enumTypes[0]
}

func (x WebErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WebErrCode.Descriptor instead.
func (WebErrCode) EnumDescriptor() ([]byte, []int) {
	return file_errcode_web_ecode_proto_rawDescGZIP(), []int{0}
}

var File_errcode_web_ecode_proto protoreflect.FileDescriptor

var file_errcode_web_ecode_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x77, 0x65, 0x62, 0x5f, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x72, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x2a, 0xc2, 0x01, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x5a, 0x6f, 0x6e, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0xe1, 0xd4, 0x03, 0x12, 0x1b, 0x0a, 0x15,
	0x57, 0x61, 0x72, 0x5a, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x68, 0x69, 0x62, 0x69, 0x74, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x10, 0xe2, 0xd4, 0x03, 0x12, 0x16, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x57, 0x61, 0x72, 0x5a, 0x6f, 0x6e, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0xe3, 0xd4,
	0x03, 0x12, 0x19, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0xe4, 0xd4, 0x03, 0x12, 0x1e, 0x0a, 0x18,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x68, 0x69, 0x62,
	0x69, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x10, 0xe5, 0xd4, 0x03, 0x12, 0x19, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x10, 0xe6, 0xd4, 0x03, 0x42, 0x1a, 0x5a, 0x18, 0x78, 0x79, 0x33, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errcode_web_ecode_proto_rawDescOnce sync.Once
	file_errcode_web_ecode_proto_rawDescData = file_errcode_web_ecode_proto_rawDesc
)

func file_errcode_web_ecode_proto_rawDescGZIP() []byte {
	file_errcode_web_ecode_proto_rawDescOnce.Do(func() {
		file_errcode_web_ecode_proto_rawDescData = protoimpl.X.CompressGZIP(file_errcode_web_ecode_proto_rawDescData)
	})
	return file_errcode_web_ecode_proto_rawDescData
}

var file_errcode_web_ecode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errcode_web_ecode_proto_goTypes = []interface{}{
	(WebErrCode)(0), // 0: errcode.WebErrCode
}
var file_errcode_web_ecode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errcode_web_ecode_proto_init() }
func file_errcode_web_ecode_proto_init() {
	if File_errcode_web_ecode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errcode_web_ecode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errcode_web_ecode_proto_goTypes,
		DependencyIndexes: file_errcode_web_ecode_proto_depIdxs,
		EnumInfos:         file_errcode_web_ecode_proto_enumTypes,
	}.Build()
	File_errcode_web_ecode_proto = out.File
	file_errcode_web_ecode_proto_rawDesc = nil
	file_errcode_web_ecode_proto_goTypes = nil
	file_errcode_web_ecode_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.5.1
// source: ret/code.proto

package ret

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

type Code int32

const (
	Code_Ok      Code = 0
	Code_IsNil   Code = 1 //
	Code_IsFalse Code = 2
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0: "Ok",
		1: "IsNil",
		2: "IsFalse",
	}
	Code_value = map[string]int32{
		"Ok":      0,
		"IsNil":   1,
		"IsFalse": 2,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_ret_code_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_ret_code_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_ret_code_proto_rawDescGZIP(), []int{0}
}

var File_ret_code_proto protoreflect.FileDescriptor

var file_ret_code_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x65, 0x74, 0x2a, 0x26, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a,
	0x02, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x73, 0x4e, 0x69, 0x6c, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x49, 0x73, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x10, 0x02, 0x42, 0x1a, 0x5a,
	0x18, 0x67, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ret_code_proto_rawDescOnce sync.Once
	file_ret_code_proto_rawDescData = file_ret_code_proto_rawDesc
)

func file_ret_code_proto_rawDescGZIP() []byte {
	file_ret_code_proto_rawDescOnce.Do(func() {
		file_ret_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_ret_code_proto_rawDescData)
	})
	return file_ret_code_proto_rawDescData
}

var file_ret_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ret_code_proto_goTypes = []interface{}{
	(Code)(0), // 0: ret.Code
}
var file_ret_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ret_code_proto_init() }
func file_ret_code_proto_init() {
	if File_ret_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ret_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ret_code_proto_goTypes,
		DependencyIndexes: file_ret_code_proto_depIdxs,
		EnumInfos:         file_ret_code_proto_enumTypes,
	}.Build()
	File_ret_code_proto = out.File
	file_ret_code_proto_rawDesc = nil
	file_ret_code_proto_goTypes = nil
	file_ret_code_proto_depIdxs = nil
}

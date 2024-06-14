// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.5.1
// source: inner/gm.proto

package pbi

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

// 发送邮件 - toUid==1表示全局邮件
type GmSendMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail *MailTemplate `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
}

func (x *GmSendMail) Reset() {
	*x = GmSendMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inner_gm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GmSendMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GmSendMail) ProtoMessage() {}

func (x *GmSendMail) ProtoReflect() protoreflect.Message {
	mi := &file_inner_gm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GmSendMail.ProtoReflect.Descriptor instead.
func (*GmSendMail) Descriptor() ([]byte, []int) {
	return file_inner_gm_proto_rawDescGZIP(), []int{0}
}

func (x *GmSendMail) GetMail() *MailTemplate {
	if x != nil {
		return x.Mail
	}
	return nil
}

var File_inner_gm_proto protoreflect.FileDescriptor

var file_inner_gm_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x67, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x62, 0x69, 0x1a, 0x12, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0a, 0x47, 0x6d, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x69, 0x2e, 0x4d, 0x61, 0x69, 0x6c,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x1a,
	0x5a, 0x18, 0x67, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_inner_gm_proto_rawDescOnce sync.Once
	file_inner_gm_proto_rawDescData = file_inner_gm_proto_rawDesc
)

func file_inner_gm_proto_rawDescGZIP() []byte {
	file_inner_gm_proto_rawDescOnce.Do(func() {
		file_inner_gm_proto_rawDescData = protoimpl.X.CompressGZIP(file_inner_gm_proto_rawDescData)
	})
	return file_inner_gm_proto_rawDescData
}

var file_inner_gm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_inner_gm_proto_goTypes = []interface{}{
	(*GmSendMail)(nil),   // 0: pbi.GmSendMail
	(*MailTemplate)(nil), // 1: pbi.MailTemplate
}
var file_inner_gm_proto_depIdxs = []int32{
	1, // 0: pbi.GmSendMail.mail:type_name -> pbi.MailTemplate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inner_gm_proto_init() }
func file_inner_gm_proto_init() {
	if File_inner_gm_proto != nil {
		return
	}
	file_inner_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_inner_gm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GmSendMail); i {
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
			RawDescriptor: file_inner_gm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inner_gm_proto_goTypes,
		DependencyIndexes: file_inner_gm_proto_depIdxs,
		MessageInfos:      file_inner_gm_proto_msgTypes,
	}.Build()
	File_inner_gm_proto = out.File
	file_inner_gm_proto_rawDesc = nil
	file_inner_gm_proto_goTypes = nil
	file_inner_gm_proto_depIdxs = nil
}

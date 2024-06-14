// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.5.1
// source: ret/common.proto

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code   `protobuf:"varint,1,opt,name=code,proto3,enum=ret.Code" json:"code,omitempty"` //错误码
	Des  string `protobuf:"bytes,2,opt,name=des,proto3" json:"des,omitempty"`                  //描述
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_ret_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_ret_common_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Ok
}

func (x *Error) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

// 模板id和数量
type TItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid int32 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Num int64 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *TItem) Reset() {
	*x = TItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TItem) ProtoMessage() {}

func (x *TItem) ProtoReflect() protoreflect.Message {
	mi := &file_ret_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TItem.ProtoReflect.Descriptor instead.
func (*TItem) Descriptor() ([]byte, []int) {
	return file_ret_common_proto_rawDescGZIP(), []int{1}
}

func (x *TItem) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *TItem) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

// 唯一id和数量
type UItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *UItem) Reset() {
	*x = UItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UItem) ProtoMessage() {}

func (x *UItem) ProtoReflect() protoreflect.Message {
	mi := &file_ret_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UItem.ProtoReflect.Descriptor instead.
func (*UItem) Descriptor() ([]byte, []int) {
	return file_ret_common_proto_rawDescGZIP(), []int{2}
}

func (x *UItem) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UItem) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

// 背包道具
type AItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Tid   int32  `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`
	Num   int64  `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	Ctime int64  `protobuf:"varint,4,opt,name=ctime,proto3" json:"ctime,omitempty"` //获得时间
}

func (x *AItem) Reset() {
	*x = AItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AItem) ProtoMessage() {}

func (x *AItem) ProtoReflect() protoreflect.Message {
	mi := &file_ret_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AItem.ProtoReflect.Descriptor instead.
func (*AItem) Descriptor() ([]byte, []int) {
	return file_ret_common_proto_rawDescGZIP(), []int{3}
}

func (x *AItem) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AItem) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *AItem) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *AItem) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

// 英雄
type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	HeroId uint64 `protobuf:"varint,2,opt,name=heroId,proto3" json:"heroId,omitempty"`
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_ret_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_ret_common_proto_rawDescGZIP(), []int{4}
}

func (x *Hero) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Hero) GetHeroId() uint64 {
	if x != nil {
		return x.HeroId
	}
	return 0
}

var File_ret_common_proto protoreflect.FileDescriptor

var file_ret_common_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x65, 0x74, 0x1a, 0x0e, 0x72, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1d, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09,
	0x2e, 0x72, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x65,
	0x73, 0x22, 0x2b, 0x0a, 0x05, 0x54, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2b,
	0x0a, 0x05, 0x55, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x53, 0x0a, 0x05, 0x41,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x30, 0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x72, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x72, 0x6f,
	0x49, 0x64, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ret_common_proto_rawDescOnce sync.Once
	file_ret_common_proto_rawDescData = file_ret_common_proto_rawDesc
)

func file_ret_common_proto_rawDescGZIP() []byte {
	file_ret_common_proto_rawDescOnce.Do(func() {
		file_ret_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_ret_common_proto_rawDescData)
	})
	return file_ret_common_proto_rawDescData
}

var file_ret_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ret_common_proto_goTypes = []interface{}{
	(*Error)(nil), // 0: ret.Error
	(*TItem)(nil), // 1: ret.TItem
	(*UItem)(nil), // 2: ret.UItem
	(*AItem)(nil), // 3: ret.AItem
	(*Hero)(nil),  // 4: ret.Hero
	(Code)(0),     // 5: ret.Code
}
var file_ret_common_proto_depIdxs = []int32{
	5, // 0: ret.Error.code:type_name -> ret.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ret_common_proto_init() }
func file_ret_common_proto_init() {
	if File_ret_common_proto != nil {
		return
	}
	file_ret_code_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ret_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_ret_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TItem); i {
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
		file_ret_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UItem); i {
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
		file_ret_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AItem); i {
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
		file_ret_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero); i {
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
			RawDescriptor: file_ret_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ret_common_proto_goTypes,
		DependencyIndexes: file_ret_common_proto_depIdxs,
		MessageInfos:      file_ret_common_proto_msgTypes,
	}.Build()
	File_ret_common_proto = out.File
	file_ret_common_proto_rawDesc = nil
	file_ret_common_proto_goTypes = nil
	file_ret_common_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.5.1
// source: ret/player.proto

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

// 心跳
type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{0}
}

// 获取用户数据
type GetUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserInfo) Reset() {
	*x = GetUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfo) ProtoMessage() {}

func (x *GetUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfo.ProtoReflect.Descriptor instead.
func (*GetUserInfo) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{1}
}

// 获取邮件
type GetMails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMails) Reset() {
	*x = GetMails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMails) ProtoMessage() {}

func (x *GetMails) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMails.ProtoReflect.Descriptor instead.
func (*GetMails) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{2}
}

type Heartbeat_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Heartbeat_Request) Reset() {
	*x = Heartbeat_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat_Request) ProtoMessage() {}

func (x *Heartbeat_Request) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat_Request.ProtoReflect.Descriptor instead.
func (*Heartbeat_Request) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{0, 0}
}

type Heartbeat_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Heartbeat_Reply) Reset() {
	*x = Heartbeat_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat_Reply) ProtoMessage() {}

func (x *Heartbeat_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat_Reply.ProtoReflect.Descriptor instead.
func (*Heartbeat_Reply) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{0, 1}
}

type GetUserInfo_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserInfo_Request) Reset() {
	*x = GetUserInfo_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfo_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfo_Request) ProtoMessage() {}

func (x *GetUserInfo_Request) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfo_Request.ProtoReflect.Descriptor instead.
func (*GetUserInfo_Request) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{1, 0}
}

type GetUserInfo_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *GetUserInfo_Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"` //自己的玩家信息
	Bag   *GetUserInfo_Bag  `protobuf:"bytes,2,opt,name=bag,proto3" json:"bag,omitempty"`
	Mails []uint64          `protobuf:"varint,3,rep,packed,name=Mails,proto3" json:"Mails,omitempty"` //邮件通常比较多,所以单独处理
}

func (x *GetUserInfo_Reply) Reset() {
	*x = GetUserInfo_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfo_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfo_Reply) ProtoMessage() {}

func (x *GetUserInfo_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfo_Reply.ProtoReflect.Descriptor instead.
func (*GetUserInfo_Reply) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GetUserInfo_Reply) GetBase() *GetUserInfo_Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *GetUserInfo_Reply) GetBag() *GetUserInfo_Bag {
	if x != nil {
		return x.Bag
	}
	return nil
}

func (x *GetUserInfo_Reply) GetMails() []uint64 {
	if x != nil {
		return x.Mails
	}
	return nil
}

type GetUserInfo_Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`  //uid
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` //名字
}

func (x *GetUserInfo_Base) Reset() {
	*x = GetUserInfo_Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfo_Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfo_Base) ProtoMessage() {}

func (x *GetUserInfo_Base) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfo_Base.ProtoReflect.Descriptor instead.
func (*GetUserInfo_Base) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{1, 2}
}

func (x *GetUserInfo_Base) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetUserInfo_Base) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 可根据情况拆分-单独请求此数据
type GetUserInfo_Bag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//通用资产包
	Assets map[uint64]*BItem `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //道具/皮肤/...
	//有特殊属性的单独成包
	Heroes map[uint64]*Hero `protobuf:"bytes,2,rep,name=heroes,proto3" json:"heroes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //英雄
}

func (x *GetUserInfo_Bag) Reset() {
	*x = GetUserInfo_Bag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfo_Bag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfo_Bag) ProtoMessage() {}

func (x *GetUserInfo_Bag) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfo_Bag.ProtoReflect.Descriptor instead.
func (*GetUserInfo_Bag) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{1, 3}
}

func (x *GetUserInfo_Bag) GetAssets() map[uint64]*BItem {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *GetUserInfo_Bag) GetHeroes() map[uint64]*Hero {
	if x != nil {
		return x.Heroes
	}
	return nil
}

type GetMails_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uids []uint64 `protobuf:"varint,1,rep,packed,name=uids,proto3" json:"uids,omitempty"`
}

func (x *GetMails_Request) Reset() {
	*x = GetMails_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMails_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMails_Request) ProtoMessage() {}

func (x *GetMails_Request) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMails_Request.ProtoReflect.Descriptor instead.
func (*GetMails_Request) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetMails_Request) GetUids() []uint64 {
	if x != nil {
		return x.Uids
	}
	return nil
}

type GetMails_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mails []*GetMails_Mail `protobuf:"bytes,1,rep,name=mails,proto3" json:"mails,omitempty"`
}

func (x *GetMails_Reply) Reset() {
	*x = GetMails_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMails_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMails_Reply) ProtoMessage() {}

func (x *GetMails_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMails_Reply.ProtoReflect.Descriptor instead.
func (*GetMails_Reply) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{2, 1}
}

func (x *GetMails_Reply) GetMails() []*GetMails_Mail {
	if x != nil {
		return x.Mails
	}
	return nil
}

// 模块
type GetMails_Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         uint64            `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                                                                                              //唯一id
	Tid         uint64            `protobuf:"varint,2,opt,name=tid,proto3" json:"tid,omitempty"`                                                                                              //模板编号，决定以下参数如何使用
	ToUid       uint64            `protobuf:"varint,3,opt,name=toUid,proto3" json:"toUid,omitempty"`                                                                                          //是否需要？
	FromUid     uint64            `protobuf:"varint,4,opt,name=fromUid,proto3" json:"fromUid,omitempty"`                                                                                      //发送方唯一id
	FromName    string            `protobuf:"bytes,5,opt,name=fromName,proto3" json:"fromName,omitempty"`                                                                                     //发送方名字
	Title       string            `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`                                                                                           //标题
	Content     string            `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`                                                                                       //内容
	Rewards     []*Item           `protobuf:"bytes,8,rep,name=rewards,proto3" json:"rewards,omitempty"`                                                                                       //奖励
	Params      map[string]string `protobuf:"bytes,9,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //参数
	CreateTime  int64             `protobuf:"varint,10,opt,name=createTime,proto3" json:"createTime,omitempty"`                                                                               //创建时间
	ReadTime    bool              `protobuf:"varint,11,opt,name=readTime,proto3" json:"readTime,omitempty"`                                                                                   //存在则表示-已读
	GetTime     bool              `protobuf:"varint,12,opt,name=getTime,proto3" json:"getTime,omitempty"`                                                                                     //存在则表示-已领取
	DeletedTime bool              `protobuf:"varint,13,opt,name=deletedTime,proto3" json:"deletedTime,omitempty"`                                                                             //存在则表示-已删除
}

func (x *GetMails_Mail) Reset() {
	*x = GetMails_Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ret_player_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMails_Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMails_Mail) ProtoMessage() {}

func (x *GetMails_Mail) ProtoReflect() protoreflect.Message {
	mi := &file_ret_player_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMails_Mail.ProtoReflect.Descriptor instead.
func (*GetMails_Mail) Descriptor() ([]byte, []int) {
	return file_ret_player_proto_rawDescGZIP(), []int{2, 2}
}

func (x *GetMails_Mail) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetMails_Mail) GetTid() uint64 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *GetMails_Mail) GetToUid() uint64 {
	if x != nil {
		return x.ToUid
	}
	return 0
}

func (x *GetMails_Mail) GetFromUid() uint64 {
	if x != nil {
		return x.FromUid
	}
	return 0
}

func (x *GetMails_Mail) GetFromName() string {
	if x != nil {
		return x.FromName
	}
	return ""
}

func (x *GetMails_Mail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMails_Mail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetMails_Mail) GetRewards() []*Item {
	if x != nil {
		return x.Rewards
	}
	return nil
}

func (x *GetMails_Mail) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GetMails_Mail) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GetMails_Mail) GetReadTime() bool {
	if x != nil {
		return x.ReadTime
	}
	return false
}

func (x *GetMails_Mail) GetGetTime() bool {
	if x != nil {
		return x.GetTime
	}
	return false
}

func (x *GetMails_Mail) GetDeletedTime() bool {
	if x != nil {
		return x.DeletedTime
	}
	return false
}

var File_ret_player_proto protoreflect.FileDescriptor

var file_ret_player_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x65, 0x74, 0x1a, 0x10, 0x72, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x09, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x07, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xc1, 0x03, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x70, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72,
	0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x62, 0x61, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x67, 0x52, 0x03, 0x62, 0x61,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x05, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x2c, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x86, 0x02, 0x0a, 0x03, 0x42, 0x61, 0x67, 0x12, 0x38, 0x0a,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x72, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x42, 0x61, 0x67, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x67, 0x2e, 0x48, 0x65,
	0x72, 0x6f, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x65,
	0x73, 0x1a, 0x45, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x65, 0x74, 0x2e, 0x42, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x44, 0x0a, 0x0b, 0x48, 0x65, 0x72, 0x6f,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x65, 0x74, 0x2e, 0x48,
	0x65, 0x72, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x95,
	0x04, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x1d, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x75, 0x69, 0x64, 0x73, 0x1a, 0x31, 0x0a, 0x05, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0xb6, 0x03,
	0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x55, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x65, 0x74, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x67, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x72, 0x61, 0x69, 0x6e, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ret_player_proto_rawDescOnce sync.Once
	file_ret_player_proto_rawDescData = file_ret_player_proto_rawDesc
)

func file_ret_player_proto_rawDescGZIP() []byte {
	file_ret_player_proto_rawDescOnce.Do(func() {
		file_ret_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_ret_player_proto_rawDescData)
	})
	return file_ret_player_proto_rawDescData
}

var file_ret_player_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_ret_player_proto_goTypes = []interface{}{
	(*Heartbeat)(nil),           // 0: ret.Heartbeat
	(*GetUserInfo)(nil),         // 1: ret.GetUserInfo
	(*GetMails)(nil),            // 2: ret.GetMails
	(*Heartbeat_Request)(nil),   // 3: ret.Heartbeat.Request
	(*Heartbeat_Reply)(nil),     // 4: ret.Heartbeat.Reply
	(*GetUserInfo_Request)(nil), // 5: ret.GetUserInfo.Request
	(*GetUserInfo_Reply)(nil),   // 6: ret.GetUserInfo.Reply
	(*GetUserInfo_Base)(nil),    // 7: ret.GetUserInfo.Base
	(*GetUserInfo_Bag)(nil),     // 8: ret.GetUserInfo.Bag
	nil,                         // 9: ret.GetUserInfo.Bag.AssetsEntry
	nil,                         // 10: ret.GetUserInfo.Bag.HeroesEntry
	(*GetMails_Request)(nil),    // 11: ret.GetMails.Request
	(*GetMails_Reply)(nil),      // 12: ret.GetMails.Reply
	(*GetMails_Mail)(nil),       // 13: ret.GetMails.Mail
	nil,                         // 14: ret.GetMails.Mail.ParamsEntry
	(*BItem)(nil),               // 15: ret.BItem
	(*Hero)(nil),                // 16: ret.Hero
	(*Item)(nil),                // 17: ret.Item
}
var file_ret_player_proto_depIdxs = []int32{
	7,  // 0: ret.GetUserInfo.Reply.base:type_name -> ret.GetUserInfo.Base
	8,  // 1: ret.GetUserInfo.Reply.bag:type_name -> ret.GetUserInfo.Bag
	9,  // 2: ret.GetUserInfo.Bag.assets:type_name -> ret.GetUserInfo.Bag.AssetsEntry
	10, // 3: ret.GetUserInfo.Bag.heroes:type_name -> ret.GetUserInfo.Bag.HeroesEntry
	15, // 4: ret.GetUserInfo.Bag.AssetsEntry.value:type_name -> ret.BItem
	16, // 5: ret.GetUserInfo.Bag.HeroesEntry.value:type_name -> ret.Hero
	13, // 6: ret.GetMails.Reply.mails:type_name -> ret.GetMails.Mail
	17, // 7: ret.GetMails.Mail.rewards:type_name -> ret.Item
	14, // 8: ret.GetMails.Mail.params:type_name -> ret.GetMails.Mail.ParamsEntry
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_ret_player_proto_init() }
func file_ret_player_proto_init() {
	if File_ret_player_proto != nil {
		return
	}
	file_ret_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ret_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
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
		file_ret_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfo); i {
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
		file_ret_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMails); i {
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
		file_ret_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat_Request); i {
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
		file_ret_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat_Reply); i {
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
		file_ret_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfo_Request); i {
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
		file_ret_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfo_Reply); i {
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
		file_ret_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfo_Base); i {
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
		file_ret_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfo_Bag); i {
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
		file_ret_player_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMails_Request); i {
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
		file_ret_player_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMails_Reply); i {
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
		file_ret_player_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMails_Mail); i {
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
			RawDescriptor: file_ret_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ret_player_proto_goTypes,
		DependencyIndexes: file_ret_player_proto_depIdxs,
		MessageInfos:      file_ret_player_proto_msgTypes,
	}.Build()
	File_ret_player_proto = out.File
	file_ret_player_proto_rawDesc = nil
	file_ret_player_proto_goTypes = nil
	file_ret_player_proto_depIdxs = nil
}

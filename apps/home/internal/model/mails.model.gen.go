// Code generated by https://github.com/chenxyzl/gen_tools; DO NOT EDIT.
// gen_tools version: 1.0.3
// generate time: 2024-06-18 17:54:06
// src code version:
// src code commit time :
package model

import (
	"encoding/json"
	"fmt"
	"github.com/chenxyzl/gsgen/gsmodel"
)

func (s *Mail) GetUid() uint64 {
	return s.uid
}
func (s *Mail) SetUid(v uint64) {
	s.uid = v
	s.UpdateDirty(1 << 0)
}
func (s *Mail) GetTid() uint64 {
	return s.tid
}
func (s *Mail) SetTid(v uint64) {
	s.tid = v
	s.UpdateDirty(1 << 1)
}
func (s *Mail) GetTouid() uint64 {
	return s.toUid
}
func (s *Mail) SetTouid(v uint64) {
	s.toUid = v
	s.UpdateDirty(1 << 2)
}
func (s *Mail) GetFromuid() uint64 {
	return s.fromUid
}
func (s *Mail) SetFromuid(v uint64) {
	s.fromUid = v
	s.UpdateDirty(1 << 3)
}
func (s *Mail) GetFromname() string {
	return s.fromName
}
func (s *Mail) SetFromname(v string) {
	s.fromName = v
	s.UpdateDirty(1 << 4)
}
func (s *Mail) GetTitle() string {
	return s.title
}
func (s *Mail) SetTitle(v string) {
	s.title = v
	s.UpdateDirty(1 << 5)
}
func (s *Mail) GetContent() string {
	return s.content
}
func (s *Mail) SetContent(v string) {
	s.content = v
	s.UpdateDirty(1 << 6)
}
func (s *Mail) GetRewards() *gsmodel.DList[*AItem] {
	return s.rewards
}
func (s *Mail) SetRewards(v *gsmodel.DList[*AItem]) {
	if v != nil {
		v.SetParent(7, s.UpdateDirty)
	}
	s.rewards = v
	s.UpdateDirty(1 << 7)
}
func (s *Mail) GetParams() *gsmodel.DMap[string, string] {
	return s.params
}
func (s *Mail) SetParams(v *gsmodel.DMap[string, string]) {
	if v != nil {
		v.SetParent(8, s.UpdateDirty)
	}
	s.params = v
	s.UpdateDirty(1 << 8)
}
func (s *Mail) GetCreatetime() int64 {
	return s.createTime
}
func (s *Mail) SetCreatetime(v int64) {
	s.createTime = v
	s.UpdateDirty(1 << 9)
}
func (s *Mail) GetReadtime() bool {
	return s.readTime
}
func (s *Mail) SetReadtime(v bool) {
	s.readTime = v
	s.UpdateDirty(1 << 10)
}
func (s *Mail) GetGettime() bool {
	return s.getTime
}
func (s *Mail) SetGettime(v bool) {
	s.getTime = v
	s.UpdateDirty(1 << 11)
}
func (s *Mail) GetDeletedtime() bool {
	return s.deletedTime
}
func (s *Mail) SetDeletedtime(v bool) {
	s.deletedTime = v
	s.UpdateDirty(1 << 12)
}
func (s *Mail) String() string {
	doc := struct {
		Uid         uint64                        `bson:"_id"`
		Tid         uint64                        `bson:"tid"`
		Touid       uint64                        `bson:"to_uid"`
		Fromuid     uint64                        `bson:"from_uid"`
		Fromname    string                        `bson:"from_name"`
		Title       string                        `bson:"title"`
		Content     string                        `bson:"content"`
		Rewards     *gsmodel.DList[*AItem]        `bson:"rewards"`
		Params      *gsmodel.DMap[string, string] `bson:"params"`
		Createtime  int64                         `bson:"create_time"`
		Readtime    bool                          `bson:"read_time"`
		Gettime     bool                          `bson:"get_time"`
		Deletedtime bool                          `bson:"deleted_time"`
	}{s.uid, s.tid, s.toUid, s.fromUid, s.fromName, s.title, s.content, s.rewards, s.params, s.createTime, s.readTime, s.getTime, s.deletedTime}
	return fmt.Sprintf("%v", &doc)
}
func (s *Mail) MarshalJSON() ([]byte, error) {
	doc := struct {
		Uid         uint64                        `bson:"_id"`
		Tid         uint64                        `bson:"tid"`
		Touid       uint64                        `bson:"to_uid"`
		Fromuid     uint64                        `bson:"from_uid"`
		Fromname    string                        `bson:"from_name"`
		Title       string                        `bson:"title"`
		Content     string                        `bson:"content"`
		Rewards     *gsmodel.DList[*AItem]        `bson:"rewards"`
		Params      *gsmodel.DMap[string, string] `bson:"params"`
		Createtime  int64                         `bson:"create_time"`
		Readtime    bool                          `bson:"read_time"`
		Gettime     bool                          `bson:"get_time"`
		Deletedtime bool                          `bson:"deleted_time"`
	}{s.uid, s.tid, s.toUid, s.fromUid, s.fromName, s.title, s.content, s.rewards, s.params, s.createTime, s.readTime, s.getTime, s.deletedTime}
	return json.Marshal(doc)
}
func (s *Mail) UnmarshalJSON(data []byte) error {
	doc := struct {
		Uid         uint64                        `bson:"_id"`
		Tid         uint64                        `bson:"tid"`
		Touid       uint64                        `bson:"to_uid"`
		Fromuid     uint64                        `bson:"from_uid"`
		Fromname    string                        `bson:"from_name"`
		Title       string                        `bson:"title"`
		Content     string                        `bson:"content"`
		Rewards     *gsmodel.DList[*AItem]        `bson:"rewards"`
		Params      *gsmodel.DMap[string, string] `bson:"params"`
		Createtime  int64                         `bson:"create_time"`
		Readtime    bool                          `bson:"read_time"`
		Gettime     bool                          `bson:"get_time"`
		Deletedtime bool                          `bson:"deleted_time"`
	}{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}
	s.SetUid(doc.Uid)
	s.SetTid(doc.Tid)
	s.SetTouid(doc.Touid)
	s.SetFromuid(doc.Fromuid)
	s.SetFromname(doc.Fromname)
	s.SetTitle(doc.Title)
	s.SetContent(doc.Content)
	s.SetRewards(doc.Rewards)
	s.SetParams(doc.Params)
	s.SetCreatetime(doc.Createtime)
	s.SetReadtime(doc.Readtime)
	s.SetGettime(doc.Gettime)
	s.SetDeletedtime(doc.Deletedtime)
	return nil
}
func (s *Mail) Clone() (*Mail, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	ret := Mail{}
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
func (s *Mails) GetMailids() *gsmodel.DList[uint64] {
	return s.mailIds
}
func (s *Mails) SetMailids(v *gsmodel.DList[uint64]) {
	if v != nil {
		v.SetParent(0, s.UpdateDirty)
	}
	s.mailIds = v
	s.UpdateDirty(1 << 0)
}
func (s *Mails) GetLastglobalmailtime() uint64 {
	return s.lastGlobalMailTime
}
func (s *Mails) SetLastglobalmailtime(v uint64) {
	s.lastGlobalMailTime = v
	s.UpdateDirty(1 << 1)
}
func (s *Mails) String() string {
	doc := struct {
		Mailids            *gsmodel.DList[uint64] `bson:"mail_ids"`
		Lastglobalmailtime uint64                 `bson:"last_global_mail_time"`
	}{s.mailIds, s.lastGlobalMailTime}
	return fmt.Sprintf("%v", &doc)
}
func (s *Mails) MarshalJSON() ([]byte, error) {
	doc := struct {
		Mailids            *gsmodel.DList[uint64] `bson:"mail_ids"`
		Lastglobalmailtime uint64                 `bson:"last_global_mail_time"`
	}{s.mailIds, s.lastGlobalMailTime}
	return json.Marshal(doc)
}
func (s *Mails) UnmarshalJSON(data []byte) error {
	doc := struct {
		Mailids            *gsmodel.DList[uint64] `bson:"mail_ids"`
		Lastglobalmailtime uint64                 `bson:"last_global_mail_time"`
	}{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}
	s.SetMailids(doc.Mailids)
	s.SetLastglobalmailtime(doc.Lastglobalmailtime)
	return nil
}
func (s *Mails) Clone() (*Mails, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	ret := Mails{}
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

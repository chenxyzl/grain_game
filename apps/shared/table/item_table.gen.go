// Code generated by https://github.com/chenxyzl/gsgen; DO NOT EDIT.
// gen_tools version: 1.1.9
// generate time: 2024-07-03 16:09:42
package table

import (
	"encoding/json"
	"fmt"
	"time"
)

func (s *ItemElem) GetTid() int32 {
	return s.tid
}
func (s *ItemElem) setTid(v int32) {
	s.tid = v
}
func (s *ItemElem) GetRare() int {
	return s.rare
}
func (s *ItemElem) setRare(v int) {
	s.rare = v
}
func (s *ItemElem) GetMaxNum() int64 {
	return s.maxNum
}
func (s *ItemElem) setMaxNum(v int64) {
	s.maxNum = v
}
func (s *ItemElem) GetStackNum() int64 {
	return s.stackNum
}
func (s *ItemElem) setStackNum(v int64) {
	s.stackNum = v
}
func (s *ItemElem) GetLimitTimeType() int32 {
	return s.limitTimeType
}
func (s *ItemElem) setLimitTimeType(v int32) {
	s.limitTimeType = v
}
func (s *ItemElem) GetLimitTime1() int32 {
	return s.limitTime1
}
func (s *ItemElem) setLimitTime1(v int32) {
	s.limitTime1 = v
}
func (s *ItemElem) GetLimitTime2() string {
	return s.limitTime2
}
func (s *ItemElem) setLimitTime2(v string) {
	s.limitTime2 = v
}
func (s *ItemElem) GetLimitTime2Format() *time.Time {
	return s.limitTime2Format
}
func (s *ItemElem) setLimitTime2Format(v *time.Time) {
	s.limitTime2Format = v
}
func (s *ItemElem) GetEffectId() int32 {
	return s.effectId
}
func (s *ItemElem) setEffectId(v int32) {
	s.effectId = v
}
func (s *ItemElem) String() string {
	doc := struct {
		Tid              int32      `json:"id"`
		Rare             int        `json:"rare"`
		MaxNum           int64      `json:"limitNum"`
		StackNum         int64      `json:"limitStack"`
		LimitTimeType    int32      `json:"time_type"`
		LimitTime1       int32      `json:"limit_type_time1"`
		LimitTime2       string     `json:"limit_type_time2"`
		LimitTime2Format *time.Time `json:"-"`
		EffectId         int32      `json:"effect_id"`
	}{s.tid, s.rare, s.maxNum, s.stackNum, s.limitTimeType, s.limitTime1, s.limitTime2, s.limitTime2Format, s.effectId}
	return fmt.Sprintf("%v", &doc)
}
func (s *ItemElem) MarshalJSON() ([]byte, error) {
	doc := struct {
		Tid              int32      `json:"id"`
		Rare             int        `json:"rare"`
		MaxNum           int64      `json:"limitNum"`
		StackNum         int64      `json:"limitStack"`
		LimitTimeType    int32      `json:"time_type"`
		LimitTime1       int32      `json:"limit_type_time1"`
		LimitTime2       string     `json:"limit_type_time2"`
		LimitTime2Format *time.Time `json:"-"`
		EffectId         int32      `json:"effect_id"`
	}{s.tid, s.rare, s.maxNum, s.stackNum, s.limitTimeType, s.limitTime1, s.limitTime2, s.limitTime2Format, s.effectId}
	return json.Marshal(doc)
}
func (s *ItemElem) UnmarshalJSON(data []byte) error {
	doc := struct {
		Tid              int32      `json:"id"`
		Rare             int        `json:"rare"`
		MaxNum           int64      `json:"limitNum"`
		StackNum         int64      `json:"limitStack"`
		LimitTimeType    int32      `json:"time_type"`
		LimitTime1       int32      `json:"limit_type_time1"`
		LimitTime2       string     `json:"limit_type_time2"`
		LimitTime2Format *time.Time `json:"-"`
		EffectId         int32      `json:"effect_id"`
	}{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}
	s.setTid(doc.Tid)
	s.setRare(doc.Rare)
	s.setMaxNum(doc.MaxNum)
	s.setStackNum(doc.StackNum)
	s.setLimitTimeType(doc.LimitTimeType)
	s.setLimitTime1(doc.LimitTime1)
	s.setLimitTime2(doc.LimitTime2)
	s.setLimitTime2Format(doc.LimitTime2Format)
	s.setEffectId(doc.EffectId)
	return nil
}
func (s *ItemElem) Clone() (*ItemElem, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	ret := ItemElem{}
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
func (s *itemTable) String() string {
	doc := struct {
	}{}
	return fmt.Sprintf("%v", &doc)
}
func (s *itemTable) MarshalJSON() ([]byte, error) {
	doc := struct {
	}{}
	return json.Marshal(doc)
}
func (s *itemTable) UnmarshalJSON(data []byte) error {
	doc := struct {
	}{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return err
	}
	return nil
}
func (s *itemTable) Clone() (*itemTable, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	ret := itemTable{}
	err = json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

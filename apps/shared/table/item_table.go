package table

import (
	"time"
)

type ItemElem struct {
	tid              int32      `json:"id"`             //模板id--id段决定了道具类型
	rare             int        `json:"rare"`           //品质/稀有度等级
	maxNum           int64      `json:"maxNum"`         //最大数量
	stackNum         int64      `json:"maxStack"`       //单个格子堆叠数量
	limitTimeType    int32      `json:"timeLimitType"`  //0不限时,1指定时长过期,2指定时间点过期
	limitTime1       int32      `json:"timeLimitType1"` //时长过期,单位秒
	limitTime2       string     `json:"timeLimitType2"` //时间点过期 eg:"2024-1-1 5:59:59"
	limitTime2Format *time.Time `json:"-"`
	effectId         int32      `json:"effect_id"` //关联到不同的效果表
	//name:名字 icon:图标 sort:排序
}

type itemTable struct {
	rowLoader[int32, *ItemElem] `in:"-"`
}

func (x *itemTable) afterLoad() error {
	var err error
	x.Range(func(k int32, v *ItemElem) bool {
		var tim time.Time
		tim, err = time.Parse(time.DateTime, v.limitTime2)
		if err != nil {
			return false
		}
		v.setLimitTime2Format(&tim)
		return true
	})
	return err

}

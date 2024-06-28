package table

import "time"

type Item struct {
	tid           int32     `json:"id"`               //模板id--id段决定了道具类型
	rare          int       `json:"rare"`             //品质/稀有度等级
	maxNum        int64     `json:"limitNum"`         //最大数量
	stackNum      int64     `json:"limitStack"`       //单个格子堆叠数量
	limitTimeType int32     `json:"time_type"`        //0不限时,1指定时长过期,2指定时间点过期
	limitTime1    int32     `json:"limit_type_time1"` //指定时长过期,单位秒
	limitTime2    time.Time `json:"limit_type_time2"` //指定时间过期 eg:"2024-1-1 5:59:59"
	effectId      int32     `json:"effect_id"`        //关联到不同的效果表
	//name:名字 icon:图标 sort:排序
}

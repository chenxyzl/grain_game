package model

// TItem 模板id-数量
type TItem struct {
	tid int32 `bson:"tid"`
	num int32 `bson:"num"`
}

// UItem 唯一id-数量
type UItem struct {
	uid uint64 `bson:"uid"`
	num int32  `bson:"num"`
}

// AItem 背包道具数据
type AItem struct {
	uid             uint64 `bson:"uid"`               //唯一id
	tid             int32  `bson:"tid"`               //模板id
	num             int64  `bson:"num"`               //数量
	receiveTimeUnix int64  `bson:"receive_time_unix"` //时间 如果0表示不限时
}

package model

// TItem 模板id对应的数据
type TItem struct {
	tid int32
	num int32
}

// UItem 唯一id对应的数量
type UItem struct {
	uid uint64
	num int32
}

// AItem 背包道具数据
type AItem struct {
	tid           int32
	num           int64
	uid           uint64 //唯一id
	limitTimeUnix int64  //时间 如果0表示不限时
}

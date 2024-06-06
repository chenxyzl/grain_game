package model

type Item struct {
	tid           int64
	num           int32
	uid           uint64 //唯一id
	limitTimeUnix int64  //时间 如果0表示不限时
}

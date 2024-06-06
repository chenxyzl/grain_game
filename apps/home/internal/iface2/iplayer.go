package iface2

import (
	"grain_game/apps/shared1/iface1"
)

type IPlayer interface {
	iface1.IEntity
	//RegisterOnTick 注册tick回调
	RegisterOnTick(fl ...OnTickFun)
	//RegisterOnCrossDay 注册跨天回调
	RegisterOnCrossDay(fl ...OnCrossDay)
	//RegisterOnCrossWeek 注册跨周回调
	RegisterOnCrossWeek(fl ...OnCrossWeek)
	//RegisterOnline 注册在线回调
	RegisterOnline(fl ...OnlineFun)
	//RegisterOffline 注册离线回调
	RegisterOffline(fl ...OfflineFun)
}

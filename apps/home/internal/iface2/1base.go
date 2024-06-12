package iface2

import (
	"grain_game/apps/shared1/iface1"
)

type IPlayerModule interface {
	iface1.IModule
	OnInit()
	OnStarted()
	OnPreStop()
	OnTickFun(nowUnix int64)
	OnCrossDay(natural bool, daysUnix ...int64)
	OnCrossWeek(natural bool, daysUnix ...int64)
	OnlineFun(nowUnix int64)
	OfflineFun(nowUnix int64)
}

type BasePlayerModule struct{ iface1.BaseModule }

func (b *BasePlayerModule) OnInit() {}

func (b *BasePlayerModule) OnStarted() {}

func (b *BasePlayerModule) OnPreStop() {}

func (b *BasePlayerModule) OnTickFun(nowUnix int64) {}

func (b *BasePlayerModule) OnCrossDay(natural bool, daysUnix ...int64) {}

func (b *BasePlayerModule) OnCrossWeek(natural bool, daysUnix ...int64) {}

func (b *BasePlayerModule) OnlineFun(nowUnix int64) {}

func (b *BasePlayerModule) OfflineFun(nowUnix int64) {}

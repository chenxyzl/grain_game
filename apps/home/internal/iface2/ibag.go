package iface2

import (
	"grain_game/apps/home/internal/model"
	"grain_game/apps/shared1/iface1"
)

type IBag interface {
	iface1.IModule
	Add(items ...*model.Item)
	Remove(items ...*model.Item)
	GetCount(tid int64) int64
	GetCountByUid(uid int64) int64
}

package iface2

import (
	"grain_game/apps/home/internal/model"
	"grain_game/apps/shared1/iface1"
	"grain_game/proto/gen/ret"
)

type IBag interface {
	iface1.IModule
	//Add 添加物品
	Add(items ...*model.AItem)
	//RemoveM 移除物品,不够数量则会panic(ret.Error)
	RemoveM(items ...*model.AItem)
	//RemoveE 移除物品, 不够数量则会返回错误码 ret.Code
	RemoveE(items ...*model.AItem) ret.Code
	//GetCountByTid 返回同uid的数量
	GetCountByTid(tid int32) uint64
	//GetCountByUid 返回唯一id对应的数量
	GetCountByUid(uid uint64) uint64
}

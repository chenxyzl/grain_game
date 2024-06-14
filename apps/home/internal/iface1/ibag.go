package iface1

import (
	"grain_game/apps/home/internal/model"
	"grain_game/proto/gen/ret"
)

type IBag interface {
	IPlayerModule
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

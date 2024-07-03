package iface1

import (
	"grain_game/apps/shared/common_model"
	"grain_game/proto/gen/ret"
)

type IBag interface {
	IPlayerModule
	//Add 添加物品
	Add(items ...*common_model.Item)
	//RemoveM 移除物品,不够数量则会panic(ret.Error)
	RemoveM(items ...*common_model.Item)
	//RemoveE 移除物品, 不够数量则会返回错误码 ret.Code
	RemoveE(items ...*common_model.Item) ret.Code
	//GetCountByTid 返回同uid的数量
	GetCountByTid(tid int32) uint64
	//GetCountByUid 返回唯一id对应的数量
	GetCountByUid(uid uint64) uint64
}

package impl

import (
	"grain_game/apps/home/internal/iface2"
	"grain_game/apps/home/internal/model"
	"grain_game/apps/shared1/iface1"
	"grain_game/proto/gen/ret"
)

var _ iface2.IBag = (*Bag)(nil)

type Bag struct {
	iface1.BaseModule
}

func (b *Bag) Add(items ...*model.AItem) {
	//TODO implement me
	panic("implement me")
}

// RemoveM 删除物品--必须
func (b *Bag) RemoveM(items ...*model.AItem) {
	//TODO implement me
	panic("implement me")
}

// RemoveE 删除物品--返回是否ok
func (b *Bag) RemoveE(items ...*model.AItem) ret.Code {
	//TODO implement me
	panic("implement me")
}

func (b *Bag) GetCount(tid int64) int64 {
	//TODO implement me
	panic("implement me")
}

func (b *Bag) GetCountByUid(uid int64) int64 {
	//TODO implement me
	panic("implement me")
}

func (b *Bag) add(item *model.AItem) {
	if item == nil {
		return
	}
	bag := iface1.Get[iface2.IBag](b.GetEntity())
	if bag == nil {

	}
}
func (b *Bag) remove(item *model.AItem) {
	if item == nil {
		return
	}
}

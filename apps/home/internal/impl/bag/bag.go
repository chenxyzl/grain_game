package bag

import (
	"grain_game/apps/home/internal/iface1"
	"grain_game/apps/home/internal/model"
	"grain_game/proto/gen/ret"
)

var _ iface1.IBag = (*Bag)(nil)

type Bag struct {
	iface1.BasePlayerModule
}

func (m *Bag) Add(items ...*model.AItem) {}

func (m *Bag) RemoveM(items ...*model.AItem) {}

func (m *Bag) RemoveE(items ...*model.AItem) ret.Code {
	return ret.Code_Ok
}

func (m *Bag) GetCountByTid(tid int32) uint64 {
	return 0
}

func (m *Bag) GetCountByUid(uid uint64) uint64 {
	//TODO implement me
	panic("implement me")
}

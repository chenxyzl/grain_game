package bag

import (
	"grain_game/apps/home/internal/iface1"
	"grain_game/apps/shared/common_model"
	"grain_game/apps/shared/table"
	"grain_game/proto/gen/ret"
)

var _ iface1.IBag = (*Bag)(nil)

type Bag struct {
	iface1.BasePlayerModule
}

func (m *Bag) Add(items ...*common_model.Item) {}

func (m *Bag) RemoveM(items ...*common_model.Item) {}

func (m *Bag) RemoveE(items ...*common_model.Item) ret.Code {
	return ret.Code_ok
}

func (m *Bag) GetCountByTid(tid int32) uint64 {
	return 0
}

func (m *Bag) GetCountByUid(uid uint64) uint64 {
	_ = table.Get().ConstTable
	table.Get().ItemTable.GetById(1).GetLimitTime1()
	table.Get().ItemTable.GetById(1).GetLimitTime2Format()
	table.Get().ItemTable.Range(func(k int32, v *table.ItemElem) bool {
		return true
	})
	//TODO implement me
	panic("implement me")
}

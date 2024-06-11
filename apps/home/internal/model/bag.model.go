package model

var _ IModel = (*Bag)(nil)

type Bag struct {
	items map[uint64]*AItem
}

func (m *Bag) Name() string { return "bag" }

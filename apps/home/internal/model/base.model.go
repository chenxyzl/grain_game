package model

var _ IModel = (*Base)(nil)

type Base struct {
	uid  uint64
	name string
}

func (m *Base) Name() string { return "bag" }

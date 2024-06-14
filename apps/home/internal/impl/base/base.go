package base

import "grain_game/apps/home/internal/iface1"

var _ iface1.IBase = (*Base)(nil)

type Base struct {
	iface1.BasePlayerModule
}

func (m *Base) SetName(name string) {
	//TODO implement me
	panic("implement me")
}

func (m *Base) GetName() string {
	//TODO implement me
	panic("implement me")
}

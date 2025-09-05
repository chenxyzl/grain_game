package base

import (
	"grain_game/apps/home/internal/iface"
)

var _ iface.IBase = (*Base)(nil)

type Base struct {
	iface.BasePlayerModule
}

func (m *Base) SetName(name string) {
	//TODO implement me
	panic("implement me")
}

func (m *Base) GetName() string {
	//TODO implement me
	panic("implement me")
}

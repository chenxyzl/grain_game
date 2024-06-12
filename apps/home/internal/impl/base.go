package impl

import "grain_game/apps/home/internal/iface2"

var _ iface2.IBase = (*Base)(nil)

type Base struct {
	iface2.BasePlayerModule
}

func (m *Base) OnInit() {

}

func (m *Base) SetName(name string) {
	//TODO implement me
	panic("implement me")
}

func (m *Base) GetName() string {
	//TODO implement me
	panic("implement me")
}

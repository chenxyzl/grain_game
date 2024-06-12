package iface2

type IBase interface {
	IPlayerModule
	SetName(name string)
	GetName() string
}

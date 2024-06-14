package iface1

type IBase interface {
	IPlayerModule
	SetName(name string)
	GetName() string
}

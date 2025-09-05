package iface

type IBase interface {
	IPlayerModule
	SetName(name string)
	GetName() string
}

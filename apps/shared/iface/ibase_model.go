package iface

type IModel interface {
	Name() string
	Load() error
	Save() error
	Delete() error
	IsDirty() bool
	CleanDirty()
}

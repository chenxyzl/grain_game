package iface

type IModule interface {
	init(player IEntity)
	GetEntity() IEntity
}

type BaseModule struct {
	entity IEntity
}

func (m *BaseModule) init(entity IEntity) {
	m.entity = entity
}

func (m *BaseModule) GetEntity() IEntity {
	return m.entity
}

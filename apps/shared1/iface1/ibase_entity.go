package iface1

import (
	"github.com/chenxyzl/grain/actor"
	"grain_game/apps/shared1/helper1"
	"reflect"
)

type IEntity interface {
	actor.IActor
	//RegisterModule 注册model
	//注:重复注册相同名字的model会panic
	//期望RegisterModule[T model](); 可惜go不支持
	RegisterModule(module IModule)
	//GetModule 获取Module
	//期望GetModule[T model]()T; 可惜go不支持
	getModule(module IModule) IModule
}

type BaseEntity struct {
	actor.BaseActor
	modules map[string]IModule
}

func NewBaseEntity() *BaseEntity {
	return &BaseEntity{modules: make(map[string]IModule)}
}

// actor life
func (e *BaseEntity) Started()                  {}
func (e *BaseEntity) PreStop()                  {}
func (e *BaseEntity) Receive(ctx actor.Context) {}

func (e *BaseEntity) RegisterModule(module IModule) {
	name := reflect.TypeOf(module).Elem().Name()
	if _, ok := e.modules[name]; ok {
		panic("duplicate module name: " + name)
	}
	e.modules[name] = module
	module.init(e)
}

func (e *BaseEntity) getModule(module IModule) IModule {
	name := reflect.TypeOf(module).Elem().Name()
	return e.modules[name]
}

func Get[MOD IModule](entity IEntity) MOD {
	var module MOD
	return entity.getModule(module).(MOD)
}

func GetMust[MOD IModule](entity IEntity) MOD {
	var module MOD
	module = entity.getModule(module).(MOD)
	v := helper1.NotNull(module, "not found module")
	return v
}

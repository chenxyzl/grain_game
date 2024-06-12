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
	getModule(name string) IModule
}

type BaseEntity struct {
	actor.BaseActor
	modules map[string]IModule
}

func NewBaseEntity() *BaseEntity {
	return &BaseEntity{modules: make(map[string]IModule)}
}

/*---------------------------------------actor life---------------------------------------*/

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

func (e *BaseEntity) getModule(name string) IModule {
	return e.modules[name]
}

func Get[T IModule](entity IEntity) T {
	var modType = reflect.TypeOf(new(T)).Elem() //todo 考虑用组件名字来获取组件,减少反射
	name := modType.Name()[1:]
	mod := entity.getModule(name)
	var zero T
	if mod == nil {
		return zero
	}
	return mod.(T)
}

func GetMust[T IModule](entity IEntity) T {
	com := Get[T](entity)
	var t IModule = com
	return helper1.NotNull(t, "module not found").(T)
}

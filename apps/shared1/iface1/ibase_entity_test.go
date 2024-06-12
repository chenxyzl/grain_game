package iface1

import (
	"fmt"
	"testing"
)

type ITestModule interface {
	IModule
}
type TestModule struct {
	BaseModule
}

type TestEntity struct {
	*BaseEntity
}

func TestGetMust(t *testing.T) {
	entity := &TestEntity{BaseEntity: NewBaseEntity()}
	//entity.RegisterModule(&TestModule{})
	v := GetMust[ITestModule](entity)
	fmt.Println(v)
}

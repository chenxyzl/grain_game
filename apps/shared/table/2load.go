package table

import (
	"fmt"
	"reflect"
	"sync/atomic"
)

var table atomic.Value //配置数据

func Reload(path string) error {
	v := &Table{}
	m := reflect.ValueOf(v).Elem()
	//先加载iLoader
	for i := 0; i < m.NumField(); i++ {
		file := m.Type().Field(i).Tag.Get("file")
		if file != "" {
			cv := reflect.New(m.Field(i).Type().Elem())
			m.Field(i).Set(cv)
			c := cv.Interface()
			v, ok := c.(iLoader)
			if !ok {
				return fmt.Errorf("file=%s, table does not implement iLoader interface", file)
			}
			err := v.load(path + file)
			if err != nil {
				return fmt.Errorf("file=%s, iLoaderfailed, err:%s", file, err)
			}
		}
	}
	//再加载finalLoader
	for i := 0; i < m.NumField(); i++ {
		file := m.Type().Field(i).Tag.Get("file")
		if file == "" {
			cv := reflect.New(m.Field(i).Type().Elem())
			m.Field(i).Set(cv)
			c := cv.Interface()
			v, ok := c.(finalLoader)
			if !ok {
				return fmt.Errorf("file=%s, table does not implement finalLoader interface", file)
			}
			err := v.load()
			if err != nil {
				return fmt.Errorf("file=%s, finalLoader failed, err:%s", file, err)
			}
		}
	}
	table.Store(v)
	return nil
}

func Get() *Table {
	return table.Load().(*Table)
}

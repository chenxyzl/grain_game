package helper1

import (
	"grain_game/proto/gen/ret"
	"reflect"
	"strings"
)

func NotNull[T comparable](v T, des ...string) T {
	var t T
	if v == t {
		ref := reflect.TypeOf(v)
		if ref.Kind() == reflect.Ptr {
			ref = ref.Elem()
		}
		des = append(des, ref.Name()+": is null")
		panic(ret.NewError(ret.Code_IsNil, strings.Join(des, ",")))
	}
	return v
}

func MustOk(code ret.Code, des ...string) {
	if code != ret.Code_Ok {
		panic(ret.NewError(code, strings.Join(des, ",")))
	}
}

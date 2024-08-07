package utils

import (
	"grain_game/proto/gen/ret"
	"strings"
)

func NotNull[T comparable](v T, des ...string) T {
	var t T
	if v == t {
		des = append(des, "must not null")
		panic(ret.NewError(ret.Code_isNil, strings.Join(des, ",")))
	}
	return v
}

func MustTrue(b bool, des ...string) {
	if !b {
		des = append(des, "must true")
		panic(ret.NewError(ret.Code_isFalse, strings.Join(des, ",")))
	}
}

func MustOk(code ret.Code, des ...string) {
	if code != ret.Code_ok {
		des = append(des, "must ret.Ok")
		panic(ret.NewError(code, strings.Join(des, ",")))
	}
}

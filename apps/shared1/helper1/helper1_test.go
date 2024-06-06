package helper1

import (
	"fmt"
	"testing"
)

type testCase struct{}

func TestNotNull(t *testing.T) {
	v := Todo()
	fmt.Println(v == nil)

	var v2 interface{}
	fmt.Println(v2 == nil)

	var res *err
	fmt.Println(res == nil)

	var v1 *testCase
	NotNull(0)
	NotNull(v1)
}

type err struct {
	Code int64
	Msg  string
}

func Todo() interface{} {
	var res *err
	return res
}

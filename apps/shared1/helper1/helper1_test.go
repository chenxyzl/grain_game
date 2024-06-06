package helper1

import (
	"fmt"
	"testing"
	"time"
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

func fp() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover")
		}
	}()
	i := 0
	one_sec := time.NewTicker(1 * time.Second)
	for i < 100 {
		i++
		select {
		case r, ok := <-one_sec.C:
			fmt.Println(r, ok)
			if i == 10 {
				panic("panic")
			}
		}
	}
}

func TestPanic(t *testing.T) {
	fp()
}

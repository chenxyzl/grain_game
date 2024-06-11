package model

import (
	"fmt"
	"testing"
)

func TestModel(t *testing.T) {
	{
		a := Bag{}
		name := a.Name()
		fmt.Println(name)
	}
	{
		a := Base{}
		name := a.Name()
		fmt.Println(name)
	}
}

package table

import (
	"github.com/chenxyzl/gsgen/gsmodel"
)

type ConstElem struct {
	testInt         int32                       `json:"testInt"`
	testStr         int32                       `json:"testString"`
	testType        *Item                       `json:"testType"`
	testIntSlice    *gsmodel.AList[int32]       `json:"testIntSlice"`
	testStrSlice    *gsmodel.AList[int32]       `json:"testStringSlice"`
	testMap         *gsmodel.AMap[int32, int32] `json:"testMap"`
	testMapWithType *gsmodel.AMap[int32, *Item] `json:"testMapWithType"`
}

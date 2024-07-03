package table

import "github.com/chenxyzl/gsgen/gsmodel"

type ConstElem struct {
	testInt         int32                       `json:"test"`
	testStr         int32                       `json:"test_str"`
	testIntSlice    *gsmodel.AList[int32]       `json:"test_int_slice"`
	testStrSlice    *gsmodel.AList[int32]       `json:"test_string_slice"`
	testMap         *gsmodel.AMap[int32, int32] `json:"test_map"`
	testMapWithType *gsmodel.AMap[int32, *Item] `json:"test_map_with_type"`
}

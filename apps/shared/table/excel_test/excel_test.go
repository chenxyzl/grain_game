package excel_test

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestReadExcel(t *testing.T) {
	f, err := excelize.OpenFile("")
	if err != nil {
		t.Fatal(err)
	}
	for _, sheetName := range f.GetSheetList() {
		fmt.Println(sheetName)
	}

}

func TestMarshalMap(t *testing.T) {
	var a = "[1,2,3]"
	var b []int
	err := json.Unmarshal([]byte(a), &b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(b)
}
func TestMarshalList(t *testing.T) {
	type T struct {
		A int32  `json:"a"`
		B string `json:"b"`
	}
	var a = &T{A: 1, B: "2"}
	by, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	var s = string(by)
	err = json.Unmarshal([]byte(s), a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

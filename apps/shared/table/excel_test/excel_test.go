package excel_test

import (
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

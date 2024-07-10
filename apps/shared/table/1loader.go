package table

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
)

// iLoader 加载接口的实现
type iLoader interface {
	load(file string) error
	afterLoad() error
	check() error
}

// finalLoader iLoader之后的实现,一般用于在iLoad加载完成后组装数据重新构成表
type finalLoader interface {
	load() error
	afterLoad() error
	check() error
}

func readExcel(excelFile string, sheetName string, rang func(rowData [][]string)) error {
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		return errors.Join(fmt.Errorf("read file err, excel:%v", excelFile), err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("close file err, excel:%v", excelFile)
		}
	}()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return errors.Join(fmt.Errorf("get rows err, excel:%v", excelFile), err)
	}

	//colLinkTable

	rang(rows)
	return nil
}

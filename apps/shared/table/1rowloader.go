package table

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

type rowLoader[k comparable, T any] struct {
	list []T
	data map[k]T
}

// load  unmarshal from csv
func (l *rowLoader[K, T]) load(file string) error {
	var err error
	var excelFile = file
	var sheetName = "Sheet1"
	var fileAndSheetName = strings.Split(file, ";")
	if len(fileAndSheetName) > 1 {
		excelFile = fileAndSheetName[0]
		sheetName = fileAndSheetName[1]
	}

	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		return errors.Join(fmt.Errorf("read file err, excel:%v", file), err)
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

	for _, row := range rows {
		//todo parse row
	}
	return err
}

// afterLoad overwrite for data
func (l *rowLoader[K, T]) afterLoad() error {
	return nil
}

// check overwrite for check
func (l *rowLoader[K, T]) check() error {
	return nil
}

// GetById return data by id
func (l *rowLoader[K, T]) GetById(id K) T {
	return l.data[id]
}

// Range for range
func (l *rowLoader[K, T]) Range(f func(k K, v T) bool) {
	for k, v := range l.data {
		f(k, v)
	}
}

package table

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

type rowLoader[k comparable, T any] struct {
	list        []T
	data        map[k]T
	idxes       map[string]int
	defaultVals map[string]string
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

	for rowIdx, row := range rows {
		_ = rowIdx
		switch rowIdx {
		case 0:
			l.parseHeader(row)
		case 1: //name, ignore
		case 2: //default
			l.parseDefault(row)
		case 3: //link
			l.parseLink(row)
		case 4: //param type, ignore
		case 5: //export type, ignore
		default:
			l.parseData(row)
		}
		for _, val := range row {
			fmt.Println(val)
			//
		}
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

func (l *rowLoader[K, T]) parseHeader(row []string) {
	l.idxes = make(map[string]int)
	var t T
	tt := reflect.TypeOf(t)
Next:
	for i := range tt.NumField() {
		field := tt.Field(i)
		tag := field.Tag.Get("json")
		for idx, name := range row {
			if name == tag {
				l.idxes[tag] = idx
				break Next
			}
		}
		panic(fmt.Errorf("not found index, jsosn tag name: %s", tag))
	}
}
func (l *rowLoader[K, T]) parseDefault(row []string) {
	l.defaultVals = make(map[string]string)
	for name, idx := range l.idxes {
		l.defaultVals[name] = row[idx]
	}
}
func (l *rowLoader[K, T]) parseLink(row []string) {
	//先解析link?
}
func (l *rowLoader[K, T]) parseData(row []string) {

}
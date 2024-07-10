package table

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

const subTag = "-"
const nestTag = "+"
const aliseTag = "="

type rowLoader[k comparable, T any] struct {
	//数据相关
	list []T     //数据 list
	data map[k]T //数据的key/val
	//解析相关~
	fileFs      *excelize.File            //文件句柄 用于同表的lin和别名
	excelName   string                    //excel名字
	sheetName   string                    //sheet名字
	idxes       map[string]int            //名字对应的索引
	defaultVals map[int]string            //默认值
	aliseName   map[int]map[string]string //别名
	link        map[int][][]string        //子表
}

// load  unmarshal from csv
func (l *rowLoader[K, T]) load(file string) error {
	var err error
	l.excelName = file
	l.sheetName = "Sheet1"
	var fileAndSheetName = strings.Split(file, ";")
	if len(fileAndSheetName) > 1 {
		l.excelName = fileAndSheetName[0]
		l.sheetName = fileAndSheetName[1]
	}

	f, err := excelize.OpenFile(l.excelName)
	if err != nil {
		return errors.Join(fmt.Errorf("read file err, excel:%v", file), err)
	}
	l.fileFs = f
	defer func() {
		l.fileFs = nil
		if err = f.Close(); err != nil {
			fmt.Printf("close file err, excel:%v", l.excelName)
		}
	}()
	rows, err := f.GetRows(l.sheetName)
	if err != nil {
		return errors.Join(fmt.Errorf("get rows err, excel:%v|shellName:%v", l.excelName, l.sheetName), err)
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
		for colIdx, name := range row {
			if name == tag {
				l.idxes[tag] = colIdx
				break Next
			}
		}
		panic(fmt.Errorf("not found index, jsosn tag name: %s", tag))
	}
}
func (l *rowLoader[K, T]) parseDefault(row []string) {
	l.defaultVals = make(map[int]string)
	for _, idx := range l.idxes {
		l.defaultVals[idx] = row[idx]
	}
}
func (l *rowLoader[K, T]) parseLink(row []string) {
	l.link = make(map[int][][]string)
	l.aliseName = make(map[int]map[string]string)
	for _, idx := range l.idxes {
		linkVal := row[idx]
		if linkVal == "" {
			continue
		}
		//
		if subTables := strings.Split(linkVal, nestTag); len(subTables) == 2 { //parse nest table
			val, err := l.readSubExcelSheet(subTables[0], subTables[1])
			if err != nil {
				panic(err)
			}
			l.link[idx] = val
			//read other table and sheet
		} else if nameAlise := strings.Split(linkVal, aliseTag); len(nameAlise) == 2 { //parse alise
			rows, err := l.readSubExcelSheet(nameAlise[0], nameAlise[1])
			if err != nil {
				panic(err)
			}
			l.aliseName[idx] = make(map[string]string)
			for i, aliseRow := range rows {
				if i == 0 {
					//第一行默认title
					continue
				}
				//
				l.aliseName[idx][aliseRow[0]] = nameAlise[1]
			}
		} else {
			panic("link must \"|\" in table|sheet for sub table or \".\" in table.sheet for name alise")
		}
	}
}
func (l *rowLoader[K, T]) readSubExcelSheet(excelName, sheetName string) ([][]string, error) {
	var subFileFs *excelize.File
	//default excelName
	if excelName == "" {
		excelName = l.excelName
	}

	if excelName == l.excelName {
		subFileFs = l.fileFs
	} else {
		f, err := excelize.OpenFile(l.excelName)
		if err != nil {
			return nil, errors.Join(fmt.Errorf("read sub excel/sheet err, excel:%v|sheet:%v", excelName, sheetName), err)
		}
		subFileFs = f
	}
	defer func() {
		if excelName != l.excelName {
			if err := subFileFs.Close(); err != nil {
				fmt.Printf("close  sub excel/sheet err, excel:%v|sheet:%v", excelName, sheetName)
			}
		}
	}()
	rows, err := subFileFs.GetRows(l.sheetName)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("get sub excel/sheet rows err, excel:%v|sheet:%v", excelName, sheetName), err)
	}
	return rows, nil
}

func (l *rowLoader[K, T]) parseData(row []string) {

	for name, colIdx := range l.idxes {
		val := row[colIdx]
		//get val or default
		if val == "" {
			val = l.defaultVals[colIdx]
		}

		//check alise
		if aliseMap, indexAliseOK := l.aliseName[colIdx]; indexAliseOK {
			if aliseVal, aliseOk := aliseMap[val]; aliseOk {
				val = aliseVal
			}
			continue
		}

		//todo get field type
		//todo check link

		//parse data
	}
}

package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
	"strings"
)

const subTag = "-"
const nestTag = "+"
const aliseTag = "="

type tRowData = []string
type tTableData = []tRowData

type rowLoader[k comparable, T any] struct {
	//数据相关
	list []T     //数据 list
	data map[k]T //数据的key/val
	//解析相关~
	fileFs         *excelize.File            //文件句柄 用于同表的lin和别名
	excelName      string                    //excel名字
	sheetName      string                    //sheet名字
	colIdxes       map[string]int            //名字对应的索引 name->colIdx
	colDefaultVals map[int]string            //默认值 colIdx->defaultVal
	colAliseName   map[int]map[string]string //别名 colIdx-alise-realVal
	colLinkTable   map[int]map[string]string //子表 colIdx-id-[]tRowData
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
		case 3: //colLinkTable
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
	//check
	if len(row) == 0 {
		panic(errors.New("row is empty"))
	}
	if row[0] != "id" {
		panic(errors.New("row[0] must named id, now is " + row[0]))
	}

	l.colIdxes = make(map[string]int)
	var t T
	tt := reflect.TypeOf(t)
Next:
	for i := range tt.NumField() {
		field := tt.Field(i)
		tag := field.Tag.Get("json")
		for colIdx, name := range row {
			if name == tag {
				l.colIdxes[tag] = colIdx
				break Next
			}
		}
		panic(fmt.Errorf("not found index, jsosn tag name: %s", tag))
	}
}
func (l *rowLoader[K, T]) parseDefault(row []string) {
	l.colDefaultVals = make(map[int]string)
	for _, idx := range l.colIdxes {
		l.colDefaultVals[idx] = row[idx]
	}
}
func (l *rowLoader[K, T]) parseLink(row []string) {
	l.colLinkTable = make(map[int]map[string]string)
	l.colAliseName = make(map[int]map[string]string)
	var t T
	var refT = reflect.TypeOf(t)
	for fieldName, idx := range l.colIdxes {
		linkVal := row[idx]
		if linkVal == "" {
			continue
		}
		field, ok := refT.FieldByName(fieldName)
		if !ok {
			panic(errors.New("field not exist, field name:" + fieldName))
		}
		//
		if subTables := strings.Split(linkVal, nestTag); len(subTables) == 2 { //parse nest table
			//todo 获取field对应的类型，如果是AMap/AList则~可以
			subRows, err := l.readSubExcelSheet(subTables[0], subTables[1])
			if err != nil {
				panic(err)
			}
			if field.Type.Kind() != reflect.Ptr {
				panic(errors.New("refT must be ptr"))
			}

			if l.colLinkTable[idx] == nil {
				l.colLinkTable[idx] = make(map[string]string)
			}
			if field.Type.Elem().Name() == "AList" {
				for _, subRow := range subRows {
					if l.colLinkTable[idx][subRow[0]] == "" {
						l.colLinkTable[idx][subRow[0]] = "[" + subRow[1] + "]"
					} else {
						//todo insert
					}
				}
			} else if field.Type.Elem().Name() == "AMap" {
				for _, subRow := range subRows {
					if l.colLinkTable[idx][subRow[0]] == "" {
						l.colLinkTable[idx][subRow[0]] = "{" + subRow[1] + ":" + subRow[2] + "}"
					} else {
						//todo insert
					}
				}
			} else {
				panic(errors.New("refT must be gsmodel.AList or gsmodel.AMap"))
			}
		} else if nameAlise := strings.Split(linkVal, aliseTag); len(nameAlise) == 2 { //parse alise
			rows, err := l.readSubExcelSheet(nameAlise[0], nameAlise[1])
			if err != nil {
				panic(err)
			}
			if l.colAliseName[idx] == nil {
				l.colAliseName[idx] = make(map[string]string)
			}
			for i, aliseRow := range rows {
				if i == 0 {
					//第一行默认title
					continue
				}
				//
				l.colAliseName[idx][aliseRow[0]] = nameAlise[1]
			}
		} else {
			panic("colLinkTable must \"|\" in table|sheet for sub table or \".\" in table.sheet for name alise")
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
	var t T
	value := reflect.ValueOf(t)
	//get id
	for fieldName, colIdx := range l.colIdxes {
		field := value.FieldByName(fieldName)
		if !field.IsValid() {
			panic(errors.New("field not exist, field fieldName:" + fieldName))
		}
		val := row[colIdx]
		//get val or default
		if val == "" {
			val = l.colDefaultVals[colIdx]
		}
		//check alise
		if aliseMap, indexAliseOK := l.colAliseName[colIdx]; indexAliseOK {
			if aliseVal, aliseOk := aliseMap[val]; aliseOk {
				val = aliseVal
			}
		}
		//check link
		if linkData, indexLinkOk := l.colLinkTable[colIdx]; indexLinkOk {
			if subRowData, ok := linkData[row[0]]; ok {
				val = subRowData
			}
		}
		//parse data
		switch field.Kind() {
		case reflect.String:
			field.SetString(val)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var v int64
			v, err := strconv.ParseInt(val, 0, field.Type().Bits())
			if err != nil {
				panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), err))
			}
			field.SetInt(v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var v uint64
			v, err := strconv.ParseUint(val, 0, field.Type().Bits())
			if err != nil {
				panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), err))
			}
			field.SetUint(v)
		case reflect.Float32, reflect.Float64:
			var v float64
			v, err := strconv.ParseFloat(val, field.Type().Bits())
			if err != nil {
				panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), err))
			}
			field.SetFloat(v)
		case reflect.Bool:
			var v bool
			v, err := strconv.ParseBool(val)
			if err != nil {
				panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), err))
			}
			field.SetBool(v)
		case reflect.Ptr:
			err := json.Unmarshal([]byte(val), field.Addr().Interface())
			if err != nil {
				panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), err))
			}
		case reflect.Struct:
			panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), errors.New("must use pointer to struct")))
		case reflect.Slice:
			panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), errors.New("must replace with gsmodel.AList")))
		case reflect.Map:
			panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), errors.New("must replace with gsmodel.AMap")))
		default:
			panic(errors.Join(fmt.Errorf("file:%s,sheet:%s,field:%s", l.excelName, l.sheetName, fieldName), errors.New("unknown type")))
		}
	}
}

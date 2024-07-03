package table

type Table struct {
	ConstTable *colLoader[*ConstElem]       `json:"constTable"`
	ItemTable  *itemTable                   `json:"itemTable"`
	MailTable  *rowLoader[int32, *MailElem] `json:"mailTable"`
}

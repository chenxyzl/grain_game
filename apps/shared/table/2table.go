package table

type Table struct {
	ConstTable *colLoader[*ConstElem]       `file:"const.xlsx"`
	ItemTable  *itemTable                   `file:"item.xlsx;item"`
	MailTable  *rowLoader[int32, *MailElem] `file:"mail.xlsx"`
}

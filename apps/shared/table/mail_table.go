package table

import "github.com/chenxyzl/gsgen/gsmodel"

type MailElem struct {
	id int32 `json:"id"`
}

type MailTable struct {
	d *gsmodel.AList[*MailElem]
	m *gsmodel.AMap[int32, *MailElem]
}

package iface1

import (
	pbi "grain_game/proto/gen/inner"
	"grain_game/proto/gen/ret"
)

type IMails interface {
	IPlayerModule
	SendMail(template *pbi.MailTemplate)
	ReadMail(mailId uint64) ret.Code
	DeleteMail(mailId uint64) ret.Code
	GetMailReward(mailId uint64) ret.Code
	GetMails(mailIds []uint64) []*ret.GetMails_Mail
}

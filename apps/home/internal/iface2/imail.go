package iface2

import "grain_game/proto/gen/ret"

type IMail interface {
	IPlayerModule
	SendMail()
	ReadMail(mailId uint64)
	DeleteMail(mailId uint64)
	GetMailReward(mailId uint64) ret.Code
	GetMails(mailIds []uint64)
}

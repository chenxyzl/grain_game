package iface2

import "grain_game/proto/gen/ret"

type IMails interface {
	IPlayerModule
	SendMail()
	ReadMail(mailId uint64) ret.Code
	DeleteMail(mailId uint64) ret.Code
	GetMailReward(mailId uint64) ret.Code
	GetMails(mailIds []uint64)
}

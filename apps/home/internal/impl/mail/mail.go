package mail

import (
	"grain_game/apps/home/internal/iface1"
	pbi "grain_game/proto/gen/inner"
	"grain_game/proto/gen/ret"
)

var _ iface1.IMails = (*Mails)(nil)

type Mails struct {
	iface1.BasePlayerModule
}

func (m *Mails) SendMail(template *pbi.MailTemplate) {
	//TODO implement me
	panic("implement me")
}

func (m *Mails) ReadMail(mailId uint64) ret.Code {
	//TODO implement me
	panic("implement me")
}

func (m *Mails) DeleteMail(mailId uint64) ret.Code {
	//TODO implement me
	panic("implement me")
}

func (m *Mails) GetMailReward(mailId uint64) ret.Code {
	//TODO implement me
	panic("implement me")
}

func (m *Mails) GetMails(mailIds []uint64) []*ret.GetMails_Mail {
	//TODO implement me
	panic("implement me")
}

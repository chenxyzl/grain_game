package internal

import (
	pbi "grain_game/proto/gen/inner"
	"grain_game/proto/gen/ret"
	"net"
	"time"

	"github.com/chenxyzl/grain"
)

var _ grain.IActor = (*session)(nil)

type session struct {
	grain.BaseActor
	wss             grain.ActorRef
	conn            net.Conn
	handler         func(ctx grain.Context)
	idleTickStopper grain.CancelScheduleFunc
}

func newSession(wss grain.ActorRef, conn net.Conn) *session {
	return &session{wss: wss, conn: conn}
}

func (p *session) Started() {
	p.Logger().Info("session started")
	p.handler = p.UnAuth
	p.resetIdleCheck()
	//todo watch home change
}

func (p *session) PreStop() {
	p.Logger().Info("session stopped")
}

func (p *session) Receive(ctx grain.Context) {
	switch ctx.Message().(type) {
	case *pbi.Tick30_Notify:
		p.idleCheckSuccess(ctx)
	default:
		//
		p.resetIdleCheck()
		//
		if p.handler != nil {
			p.handler(ctx)
		} else {
			p.Logger().Warn("session handler is nil, may session be closed")
		}
	}
}

func (p *session) UnAuth(ctx grain.Context) {
	switch msg := ctx.Message().(type) {
	case *ret.ReqPack:
		msg.GetRpcId()
	default:

	}
	//todo if auth success, change receiver
	p.handler = p.Authed
}

func (p *session) Authed(ctx grain.Context) {

}

func (p *session) resetIdleCheck() {
	if p.idleTickStopper != nil {
		p.idleTickStopper()
	}
	p.idleTickStopper = p.ScheduleSelfOnce(time.Second*30, &pbi.Tick30_Notify{})
}

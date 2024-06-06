package internal

import (
	"github.com/chenxyzl/grain/actor"
	pbi "grain_game/proto/gen/inner"
	"grain_game/proto/gen/ret"
	"net"
	"time"
)

var _ actor.IActor = (*session)(nil)

type session struct {
	actor.BaseActor
	wss             *actor.ActorRef
	conn            net.Conn
	handler         func(ctx actor.Context)
	idleTickStopper actor.CancelScheduleFunc
}

func newSession(wss *actor.ActorRef, conn net.Conn) *session {
	return &session{wss: wss, conn: conn}
}

func (p *session) Started() {
	p.Logger().Info("session started")
	p.handler = p.UnAuth
	p.resetIdleCheck()
}

func (p *session) PreStop() {
	p.Logger().Info("session stopped")
}

func (p *session) Receive(ctx actor.Context) {
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

func (p *session) UnAuth(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *ret.ReqPack:
		msg.GetRpcId()
	default:

	}
	//todo if auth success, change receiver
	p.handler = p.Authed
}

func (p *session) Authed(ctx actor.Context) {

}

func (p *session) resetIdleCheck() {
	if p.idleTickStopper != nil {
		p.idleTickStopper()
	}
	p.idleTickStopper = p.ScheduleSelfOnce(time.Second*30, &pbi.Tick30_Notify{})
}

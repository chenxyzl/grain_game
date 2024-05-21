package internal

import (
	"github.com/chenxyzl/grain/actor"
	pbo "grain_game/proto/gen/outer"
	"net"
)

var _ actor.IActor = (*session)(nil)

type session struct {
	actor.BaseActor
	wss     *actor.ActorRef
	conn    net.Conn
	handler func(ctx actor.Context)
}

func newSession(wss *actor.ActorRef, conn net.Conn) *session {
	return &session{wss: wss, conn: conn}
}

func (p *session) Started() {
	p.Logger().Info("session started")
	//todo parse outer rpc
	//todo parse inner rpc
	//todo start schedule ?
	p.handler = p.UnAuth
}

func (p *session) PreStop() {
	p.Logger().Info("session stopped")
}

func (p *session) Receive(ctx actor.Context) {
	if p.handler != nil {
		p.handler(ctx)
	} else {
		p.Logger().Warn("session handler is nil, may session be closed")
	}
}

func (p *session) UnAuth(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *pbo.ReqPack:
		msg.GetRpcId()
	default:

	}
	//todo if auth success, change receiver
	p.handler = p.Authed
}

func (p *session) Authed(ctx actor.Context) {

}

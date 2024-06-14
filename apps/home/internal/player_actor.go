package internal

import (
	"github.com/chenxyzl/grain/actor"
	"github.com/chenxyzl/grain/utils/helper"
	"grain_game/apps/common1"
	"grain_game/apps/home/internal/iface2"
	"grain_game/apps/shared1/iface1"
	pbi "grain_game/proto/gen/inner"
	"grain_game/proto/gen/ret"
	"time"
)

var _ actor.IActor = (*Player)(nil)

type Player struct {
	*iface1.BaseEntity
	modules    map[string]iface2.IPlayerModule
	modulesSl  []iface2.IPlayerModule //for range
	cancelTick actor.CancelScheduleFunc
}

func NewPlayer() *Player {
	return &Player{BaseEntity: iface1.NewBaseEntity(), modules: make(map[string]iface2.IPlayerModule)}
}

func (p *Player) Started() {
	//todo load db?
	//
	for _, f := range p.modulesSl {
		f.OnInit()
	}
	for _, f := range p.modulesSl {
		f.OnStarted()
	}
	//
	p.cancelTick = p.ScheduleSelfRepeated(time.Second, time.Second, &pbi.Tick{})
}

func (p *Player) PreStop() {
	if p.cancelTick != nil {
		p.cancelTick()
		p.cancelTick = nil
	}
	for _, f := range p.modulesSl {
		f.OnPreStop()
	}
}

func (p *Player) Receive(ctx actor.Context) {
	defer helper.Recover(func(e any, trace string) {
		if err, ok := e.(*ret.Error); ok {
			//todo send err code to client
			p.Logger().Warn("receive catch err code", "code", err.Code, "des", err.Des)
		} else {
			p.Logger().Error("receive unexpect panic", "err", e)
		}
	})
	switch _ := ctx.Message().(type) {
	case *pbi.Tick:
		p.onTick()
	default:
		if ctx.Sender().GetKind() == common1.SessionKind {
			//todo 外部rpc分发
		} else {
			//todo 内部rpc分发
		}

	}
}

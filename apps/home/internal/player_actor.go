package internal

import (
	"github.com/chenxyzl/grain/actor"
	"grain_game/apps/_common"
	"grain_game/apps/home/internal/iface1"
	"grain_game/apps/shared/iface"
	"grain_game/apps/shared/utils"
	pbi "grain_game/proto/gen/inner"
	"grain_game/proto/gen/ret"
	"time"
)

var _ actor.IActor = (*Player)(nil)

type Player struct {
	*iface.BaseEntity
	modules    map[string]iface1.IPlayerModule
	modulesSl  []iface1.IPlayerModule //for range
	cancelTick actor.CancelScheduleFunc
}

func NewPlayer() *Player {
	return &Player{BaseEntity: iface.NewBaseEntity(), modules: make(map[string]iface1.IPlayerModule)}
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
	defer utils.Recover(func(e any, trace string) {
		if err, ok := e.(*ret.Error); ok {
			//todo send err code to client
			p.Logger().Warn("receive catch err code", "code", err.Code, "des", err.Des)
		} else {
			p.Logger().Error("receive unexpect panic", "err", e)
		}
	})
	switch typ := ctx.Message().(type) {
	case *pbi.Tick:
		p.onTick()
	default:
		_ = typ
		if ctx.Sender().GetKind() == common.SessionKind {
			//todo 外部rpc分发
		} else {
			//todo 内部rpc分发
		}

	}
}

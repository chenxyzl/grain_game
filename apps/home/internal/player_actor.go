package internal

import (
	"grain_game/apps/home/internal/iface"
	"grain_game/apps/shared/common"
	"grain_game/apps/shared/giface"
	"grain_game/apps/shared/utils"
	pbi "grain_game/proto/gen/inner"
	"grain_game/proto/gen/ret"
	"time"

	"github.com/chenxyzl/grain"
)

var _ grain.IActor = (*Player)(nil)

type Player struct {
	*giface.BaseEntity
	modules    map[string]iface.IPlayerModule
	modulesSl  []iface.IPlayerModule //for range
	cancelTick grain.CancelScheduleFunc
}

func NewPlayer() *Player {
	return &Player{BaseEntity: giface.NewBaseEntity(), modules: make(map[string]iface.IPlayerModule)}
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

func (p *Player) Receive(ctx grain.Context) {
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

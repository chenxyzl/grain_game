package internal

import (
	"fmt"
	"github.com/chenxyzl/grain/actor"
	"grain_game/apps/home/internal/iface2"
	"grain_game/apps/shared1/helper1"
	"grain_game/apps/shared1/iface1"
	pbi "grain_game/proto/gen/inner"
	"grain_game/proto/gen/ret"
	"time"
)

var _ actor.IActor = (*Player)(nil)

type Player struct {
	*iface1.BaseEntity
	onInitSl      []iface2.OnInit
	onStartedSl   []iface2.OnStarted
	onPreStopSl   []iface2.OnPreStop
	onTickSl      []iface2.OnTickFun
	onCrossDaySl  []iface2.OnCrossDay
	onCrossWeekSl []iface2.OnCrossWeek
	onlineSl      []iface2.OnlineFun
	offlineSl     []iface2.OfflineFun
	cancelTick    actor.CancelScheduleFunc
}

func NewPlayer() *Player {
	return &Player{BaseEntity: iface1.NewBaseEntity()}
}

func (p *Player) Started() {
	//todo load db?
	//
	for idx, f := range p.onInitSl {
		if cod := f(); cod != ret.Code_Ok {
			helper1.MustOk(cod, "onInitSl.failed", fmt.Sprintf("idx:%v", idx))
		}
	}
	for idx, f := range p.onStartedSl {
		if cod := f(); cod != ret.Code_Ok {
			helper1.MustOk(cod, "onStartedSl.failed", fmt.Sprintf("idx:%v", idx))
		}
	}
	p.cancelTick = p.ScheduleSelfRepeated(time.Second, time.Second, &pbi.Tick{})
}

func (p *Player) PreStop() {
	if p.cancelTick != nil {
		p.cancelTick()
	}
	for idx, f := range p.onPreStopSl {
		if cod := f(); cod != ret.Code_Ok {
			helper1.MustOk(cod, "onPreStopSl.failed", fmt.Sprintf("idx:%v", idx))
		}
	}
}

func (p *Player) Receive(ctx actor.Context) {
	//todo recover, send to client?
	switch _ := ctx.Message().(type) {
	case *pbi.Tick:
		p.onTick()
	}
}

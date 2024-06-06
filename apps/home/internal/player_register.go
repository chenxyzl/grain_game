package internal

import (
	"grain_game/apps/home/internal/iface2"
)

var _ iface2.IPlayer = (*Player)(nil)

func (p *Player) RegisterOnTick(fl ...iface2.OnTickFun) {
	p.onTickSl = append(p.onTickSl, fl...)
}

func (p *Player) RegisterOnCrossDay(fl ...iface2.OnCrossDay) {
	p.onCrossDaySl = append(p.onCrossDaySl, fl...)
}

func (p *Player) RegisterOnCrossWeek(fl ...iface2.OnCrossWeek) {
	p.onCrossWeekSl = append(p.onCrossWeekSl, fl...)
}

func (p *Player) RegisterOnline(fl ...iface2.OnlineFun) {
	p.onlineSl = append(p.onlineSl, fl...)
}

func (p *Player) RegisterOffline(fl ...iface2.OfflineFun) {
	p.offlineSl = append(p.offlineSl, fl...)
}

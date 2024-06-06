package internal

import (
	"github.com/chenxyzl/grain/utils/helper"
	"grain_game/proto/gen/ret"
	"time"
)

type iPlayerModuleLife interface {
	OnInit() ret.Code
	OnStarted() ret.Code
	OnPreStop() ret.Code
	OnTickFun(nowUnix int64)
	OnCrossDay(natural bool, daysUnix ...int64)
	OnCrossWeek(natural bool, daysUnix ...int64)
	OnlineFun(nowUnix int64)
	OfflineFun(nowUnix int64)
}

func (p *Player) onTick() {
	defer helper.Recover(func(e any, trace string) {
		if err, ok := e.(*ret.Error); ok {
			p.Logger().Warn("tick catch err code", "code", err.Code, "des", err.Des)
		} else {
			//todo 应该踢下线，避免错误扩大
			p.Logger().Error("tick unexpect panic", "err", e)
		}
	})
	nowUnix := time.Now().Unix()
	for _, f := range p.modulesSl {
		f.OnTickFun(nowUnix)
	}
	//todo check cross day?
	//todo check cross week?
}

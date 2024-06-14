package internal

import (
	"grain_game/apps/shared/utils"
	"grain_game/proto/gen/ret"
	"time"
)

func (p *Player) onTick() {
	defer utils.Recover(func(e any, trace string) {
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

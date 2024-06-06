package internal

import "github.com/chenxyzl/grain/actor"

func (p *session) idleCheckSuccess(ctx actor.Context) {
	if p.conn != nil {
		p.conn.Close() // 链接断开时,actor会在外层被停止,这里不用管
		p.conn = nil
		p.Logger().Info("connection idle too long, closed connection")
	}
}

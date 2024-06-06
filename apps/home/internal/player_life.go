package internal

import "time"

func (p *Player) onTick() {
	//todo recover?
	nowUnix := time.Now().Unix()
	for _, f := range p.onTickSl {
		f(nowUnix)
	}
	//todo check cross day?
	//todo check cross week?
}

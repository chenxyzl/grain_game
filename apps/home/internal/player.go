package internal

import "github.com/chenxyzl/grain/actor"

var _ actor.IActor = (*Player)(nil)

type Player struct{ actor.BaseActor }

func (p *Player) Started() {
	//todo parse outer rpc
	//todo parse inner rpc
	//todo start schedule ?
}

func (p *Player) PreStop() {

}

func (p *Player) Receive(ctx actor.Context) {
	//TODO implement me
	panic("implement me")
}

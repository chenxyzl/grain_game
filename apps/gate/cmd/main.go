package main

import (
	"github.com/chenxyzl/grain/actor"
	"grain_game/apps/common"
	"grain_game/apps/gate/internal"
)

func main() {
	actor.InitLog("./home.log")
	//config
	config := actor.NewConfig("hello_cluster", "0.0.1", []string{"127.0.0.1:2379"})
	//system
	system := actor.NewSystem[*actor.ProviderEtcd](config)
	//start
	system.Logger().Warn("system starting")
	system.Start()
	system.Logger().Warn("system started successfully")
	//
	ws := system.Spawn(func() actor.IActor { return internal.NewWebsocketServer("/ws", "127.0.0.1", "32100") }, actor.WithOptsKindName(common.WsServerKind))
	if ws == nil {
		system.Logger().Error("failed to spawn websocket server")
		panic("failed to spawn websocket server")
	}
	//wait ctrl+c
	system.WaitStopSignal()
	//
	system.Logger().Warn("system stopped successfully")
}

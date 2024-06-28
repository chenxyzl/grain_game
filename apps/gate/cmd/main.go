package main

import (
	"github.com/chenxyzl/grain/actor"
	"grain_game/apps/gate/internal"
	"grain_game/apps/shared/common"
	"grain_game/apps/shared/config"
	"grain_game/apps/shared/runner"
)

func main() {
	runner.Run(func() {
		//cConfig
		cConfig := actor.NewConfig(config.Get().GetApp(), config.Get().GetVersion(), config.Get().GetEtcd().ToList())
		//system
		system := actor.NewSystem[*actor.ProviderEtcd](cConfig)
		//start
		system.Logger().Warn("system starting")
		system.Start()
		system.Logger().Warn("system started successfully")
		//
		ws := system.Spawn(func() actor.IActor {
			return internal.NewWebsocketServer(config.Get().GetGate().GetWspath(), "127.0.0.1", "0")
		}, actor.WithOptsKindName(common.WsServerKind))
		if ws == nil {
			system.Logger().Error("failed to spawn websocket server")
			panic("failed to spawn websocket server")
		}
		//wait ctrl+c
		system.WaitStopSignal()
		//
		system.Logger().Warn("system stopped successfully")
	})
}

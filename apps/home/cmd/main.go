package main

import (
	"github.com/chenxyzl/grain/actor"
	"grain_game/apps/home/internal"
	"grain_game/apps/shared/common"
	"grain_game/apps/shared/config"
	"grain_game/apps/shared/runner"
)

func main() {
	runner.Run(func() {
		//cConfig
		cConfig := actor.NewConfig(config.Get().GetApp(), config.Get().GetVersion(), config.Get().GetEtcd().ToList(),
			actor.WithConfigKind(common.PlayerKind, func() actor.IActor { return internal.NewPlayer() }))
		//system
		system := actor.NewSystem[*actor.ProviderEtcd](cConfig)
		//start
		system.Logger().Warn("system starting")
		system.Start()
		system.Logger().Warn("system started successfully")
		//wait ctrl+c
		system.WaitStopSignal()
		//
		system.Logger().Warn("system stopped successfully")
	})
}

package main

import (
	"grain_game/apps/home/internal"
	"grain_game/apps/shared/common"
	"grain_game/apps/shared/config"
	"grain_game/apps/shared/runner"
	pbi "grain_game/proto/gen/inner"

	"github.com/chenxyzl/grain"
)

func main() {
	runner.Run(func() {
		//system
		system := grain.NewSystem(config.Get().GetApp(), config.Get().GetVersion(), config.Get().GetEtcd().ToList(),
			grain.WithConfigKind(common.PlayerKind, func() grain.IActor { return internal.NewPlayer() }))
		//start
		system.Logger().Warn("system starting")
		system.Start()
		system.Logger().Warn("system started successfully")
		//
		system.PublishGlobal(&pbi.HomeOnline_Notify{})
		//wait ctrl+c
		system.WaitStopSignal(nil, nil)
		//
		system.Logger().Warn("system stopped successfully")
	})
}

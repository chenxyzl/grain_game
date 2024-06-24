package main

import (
	"fmt"
	"github.com/chenxyzl/grain/actor"
	"github.com/gookit/config/v2"
	"grain_game/apps/_common"
	"grain_game/apps/_runner"
	"grain_game/apps/home/internal"
	"grain_game/apps/shared/utils"
	"os"
)

func main() {
	runner.Run(func() {
		actor.InitLog(fmt.Sprintf("./%v.%v.log", utils.GetExecName(), os.Getegid()))
		//cConfig
		cConfig := actor.NewConfig(config.String("app"), config.String("version"), config.Strings("etcd"),
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

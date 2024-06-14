package main

import (
	"github.com/chenxyzl/grain/actor"
	"grain_game/apps/common"
	"grain_game/apps/home/internal"
)

func main() {
	actor.InitLog("./home.log")
	//config
	config := actor.NewConfig("hello_cluster", "0.0.1", []string{"127.0.0.1:2379"},
		actor.WithConfigKind(common.PlayerKind, func() actor.IActor { return internal.NewPlayer() }))
	//system
	system := actor.NewSystem[*actor.ProviderEtcd](config)
	//start
	system.Logger().Warn("system starting")
	system.Start()
	system.Logger().Warn("system started successfully")
	//wait ctrl+c
	system.WaitStopSignal()
	//
	system.Logger().Warn("system stopped successfully")
}

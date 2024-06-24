package runner

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"github.com/spf13/cobra"
	"grain_game"
	"grain_game/apps/shared/utils"
	"os"
)

func Run(f func()) {
	//pid
	pidPath := utils.CreatePid()
	defer utils.RemovePid(pidPath)
	//command
	rootCmd := &cobra.Command{
		Use:   "run",
		Short: "grain game",
		Run: func(cmd *cobra.Command, args []string) {
			//parse dir
			configs, err := cmd.Flags().GetStringSlice("configs")
			if err != nil {
				panic(fmt.Sprintf("param parse err: configs, err:%v", err))
			}
			if len(configs) == 0 {
				panic(fmt.Sprintf("param parse err: configs, len:0"))
			}
			config.WithOptions(config.ParseEnv)
			config.AddDriver(toml.Driver)
			err = config.LoadFiles(configs...)
			if err != nil {
				panic(err)
			}
			f()
		},
	}
	//增加默认命令
	rootCmd.Flags().StringSliceP("configs", "c", nil, "配置,按顺序后覆盖前")
	//添加命令
	rootCmd.AddCommand(grain_game.VersionCmd())
	//执行
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

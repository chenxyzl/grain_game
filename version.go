package grain_game

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

// 编译时注入
var (
	BuildTime  = ""
	CommitID   = ""
	CommitTime = ""
	Branch     = "dev"
	GoVer      = runtime.Version()
)

func String() string {
	return "版本信息:" +
		"\nBuild Time:" + BuildTime +
		"\nBranch:" + Branch +
		"\nCommit ID:" + CommitID +
		"\nCommit Time:" + CommitTime +
		"\nGoVer:" + GoVer
}

// VersionCmd 版本信息
func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "exec file version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version:%s\n", String())
		},
	}

	cmd.Flags().String("name", "World", "Name of the person to greet")

	return cmd
}

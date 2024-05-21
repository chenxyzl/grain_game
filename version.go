package grain_game

import (
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

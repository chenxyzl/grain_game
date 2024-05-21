package helper2

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

// CreatePid 创建pid文件
func CreatePid(logger *slog.Logger) string {
	// create pid file
	arg0, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	absExecFile, err := filepath.Abs(arg0)
	if err != nil {
		panic(err)
	}
	execDir, execFile := filepath.Split(absExecFile)
	pid := execDir + execFile + ".pid"
	//
	err = os.WriteFile(pid, []byte(fmt.Sprintf("%d", os.Getpid())), 0644)
	if err != nil {
		panic(err)
	}
	logger.Info("create pid success", "pid", pid)
	return pid
}

// RemovePid 删除pid文件
func RemovePid(pid string, logger *slog.Logger) {
	if pid != "" {
		err := os.Remove(pid)
		if err != nil {
			logger.Error("remove pid failed", "pid", pid, "err", err)
		} else {
			logger.Info("remove pid success", "pid", pid)
		}
	}
}

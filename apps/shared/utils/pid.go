package utils

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

// CreatePid 创建pid文件
// @return pid文件全路径
// @return 进程名字
// @return 进程id
func CreatePid(logger ...*slog.Logger) string {
	var _logger = slog.Default()
	if len(logger) > 0 {
		_logger = logger[0]
	}
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
	_logger.Info("create pid success", "pid", pid)
	return pid
}

// RemovePid 删除pid文件
func RemovePid(pid string, logger ...*slog.Logger) {
	var _logger = slog.Default()
	if len(logger) > 0 {
		_logger = logger[0]
	}

	if pid != "" {
		err := os.Remove(pid)
		if err != nil {
			_logger.Error("remove pid failed", "pid", pid, "err", err)
		} else {
			_logger.Info("remove pid success", "pid", pid)
		}
	}
}

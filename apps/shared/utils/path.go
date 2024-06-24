package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

// GetExecName 获取当前可执行文件名
func GetExecName() string {
	arg0, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	absExecFile, err := filepath.Abs(arg0)
	if err != nil {
		panic(err)
	}
	_, execFile := filepath.Split(absExecFile)
	return execFile
}

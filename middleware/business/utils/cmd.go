package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

// ExecCommand 执行脚本命令
func ExecCommand(command string) (string, error) {
	if runtime.GOOS == "windows" {
		return "", fmt.Errorf("windows not supported")
	}
	cmd := exec.Command("/bin/sh", command)

	// 捕获标准输出和标准错误
	var stdout, stderr bytes.Buffer

	// 标准输出
	cmd.Stdout = &stdout

	// 标准错误
	cmd.Stderr = &stderr

	var err error
	// 执行命令
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	if stderr.String() != "" {
		return "", errors.New(stderr.String())
	}
	return stdout.String(), nil
}

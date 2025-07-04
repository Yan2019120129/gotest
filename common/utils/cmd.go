package utils

import (
	"bytes"
	"errors"
	"os/exec"
)

type cmdType int

// ExecCommand 执行脚本命令
func ExecCommand(command string) (string, error) {
	return Exec(exec.Command("/bin/sh", "-c", command))
}

// ExecShell 执行脚本
func ExecShell(shellPath string) (string, error) {
	return Exec(exec.Command(shellPath))
}

func Exec(cmd *exec.Cmd) (string, error) {
	// 捕获标准输出和标准错误
	var stdout, stderr bytes.Buffer

	// 标准输出
	cmd.Stdout = &stdout

	// 标准错误
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	if stderr.String() != "" {
		return "", errors.New(stderr.String())
	}
	return stdout.String(), nil
}

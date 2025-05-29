package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

type cmdType int

const (
	typeCmd cmdType = iota
	typeSh
)

// ExecCommand 执行脚本命令
func ExecCommand(command string) (string, error) {
	return Exec(command, typeCmd)
}

// ExecShell 执行脚本
func ExecShell(shellPath string) (string, error) {
	return Exec(shellPath, typeSh)
}

func Exec(v string, t cmdType) (string, error) {
	if runtime.GOOS == "windows" {
		return "", fmt.Errorf("windows not supported")
	}
	var cmd *exec.Cmd

	switch t {
	case typeSh:
		cmd = exec.Command(v)
	case typeCmd:
		cmd = exec.Command("/bin/sh", "-c", v)
	}

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

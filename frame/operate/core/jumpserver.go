package core

import (
	"operate/conf"
	"os"
	"os/exec"
)

// JumpServer 关联JumpServer
func JumpServer() error {
	cmd := exec.Command("ssh", "-tt", conf.Conf.Base.JumpServer.User+"@jumpserver.onething.net", "-p", conf.Conf.Base.JumpServer.Port)

	// 连接当前终端的输入输出（实现交互）
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

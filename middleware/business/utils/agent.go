package utils

import "business/enum"

// Agent 实例
type Agent struct {
}

// NewAgent 创建Agent 实例
func NewAgent() *Agent {
	return &Agent{}
}

// Start 启动Agent
func (a *Agent) Start() (string, error) {
	return ExecShell(enum.PathStartScriptsFile)
}

// Stop 关闭Agent
func (a *Agent) Stop() (string, error) {
	return ExecShell(enum.PathStopScriptsFile)
}

// Reboot 重启Agent
func (a *Agent) Reboot() (string, error) {
	v, err := a.Stop()
	if err != nil {
		return v, err
	}
	return a.Start()
}

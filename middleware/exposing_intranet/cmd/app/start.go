package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

// StartCmd 启动命令
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start service",
	Long:  "start intranet penetration service",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 启动内网穿透服务
		if len(args) == 0 {
			return fmt.Errorf("")
		}
	},
}

// init 初始化启动命令
func init() {
	StartCmd.AddCommand()
}

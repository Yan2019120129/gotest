package app

import "github.com/spf13/cobra"

// StopCmd 关闭服务
var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop server",
	Long:  "stop intranet penetration service",
	PreRun: func(cmd *cobra.Command, args []string) {
		// 关闭内网穿透服务
	},
}

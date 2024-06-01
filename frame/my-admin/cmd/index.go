package cmd

import (
	"github.com/spf13/cobra"
	"my-admin/app"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "my-admin",
	Short:        "my-admin",
	SilenceUsage: true, // 默认使用
	Run: func(cmd *cobra.Command, args []string) {
		app.InitServer()
	},
}

// Execute 执行命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

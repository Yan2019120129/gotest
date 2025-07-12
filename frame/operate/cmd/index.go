package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "operate",
	Short: "个人使用的操作命令，添加一些常用的容能",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute 执行命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

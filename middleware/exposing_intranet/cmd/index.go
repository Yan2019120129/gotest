package cmd

import (
	"exposing_intranet/cmd/app"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = cobra.Command{
	Use:   "it",
	Short: "it",
	Long:  "intranet",
}

// 初始化项目命令
func init() {
	rootCmd.AddCommand(app.StartCmd)
	rootCmd.AddCommand(app.StopCmd)
}

// Execute 初始化命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

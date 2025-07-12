package cmd

import (
	"github.com/spf13/cobra"
)

// 上传文件命令
var up = &cobra.Command{
	Use:   "hello",
	Short: "打印问候信息",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(up)
}

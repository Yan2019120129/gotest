package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"operate/core"
)

// 上传文件命令
var jumpCmd = &cobra.Command{
	Use:   "j",
	Short: "jump",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.JumpServer(); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(jumpCmd)
}

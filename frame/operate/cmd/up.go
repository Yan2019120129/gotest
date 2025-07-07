package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"operate/core"
	"os"
)

var (
	fileName string
)

// 上传文件命令
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Upload",
	PreRun: func(cmd *cobra.Command, args []string) {
		if fileName == "" {
			fmt.Println("❌ 参数 -f (--f) 是必须的，请指定上传文件路径")
			cmd.Help()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.Up(fileName); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	upCmd.Flags().StringVarP(&fileName, "f", "f", "", "指定上传文件路径")
	rootCmd.AddCommand(upCmd)
}

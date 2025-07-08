package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"operate/core"
)

var (
	n string
)

// 上传文件命令
var historyCmd = &cobra.Command{
	Use:   "h",
	Short: "History",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.History(n); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	historyCmd.Flags().StringVarP(&n, "n", "n", "5", "指定上传文件路径")
	rootCmd.AddCommand(historyCmd)
}

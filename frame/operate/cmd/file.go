package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"operate/core"
)

var (
	path       string = ""
	targetPath string = ""
	numb       int64  = 0
)

// 上传文件命令
var fileCmd = &cobra.Command{
	Use:   "o",
	Short: "outfile",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.OutFile(targetPath, path, numb); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	fileCmd.Flags().StringVarP(&path, "p", "p", "./", "输出指定文件或目录中的文件内容")
	fileCmd.Flags().StringVarP(&targetPath, "t", "t", "", "将文本内容写入到指定目录中")
	fileCmd.Flags().Int64VarP(&numb, "n", "n", 0, "输出指定行数的文本")
	rootCmd.AddCommand(fileCmd)
}

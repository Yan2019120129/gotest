package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"operate/core"
	"strconv"
)

var (
	filePath  string
	lineStr   string
	outPath   string
	recursive bool // 新增递归标志
)

func init() {
	// 文件路径
	printCmd.Flags().StringVarP(&filePath, "file", "f", "", "文件路径 (必填)")
	_ = printCmd.MarkFlagRequired("file")

	// 打印多少行
	printCmd.Flags().StringVarP(&lineStr, "lines", "l", "", "打印前 N 行 (默认打印全部)")

	// 输出到哪个文件
	printCmd.Flags().StringVarP(&outPath, "out", "o", "", "将打印内容写入指定文件 (可选)")

	// 新增递归标志
	printCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "递归打印目录下所有文件")

	rootCmd.AddCommand(printCmd)
}

// printCmd 打印文件内容
var printCmd = &cobra.Command{
	Use:   "p",
	Short: "Print file content",
	Run: func(cmd *cobra.Command, args []string) {
		var limit int
		if lineStr == "" {
			limit = -1
		} else {
			n, err := strconv.Atoi(lineStr)
			if err != nil || n < 0 {
				fmt.Println("行数参数不正确，应为正整数")
				return
			}
			limit = n
		}

		if err := core.PrintFileContent(filePath, outPath, limit, recursive); err != nil {
			fmt.Println("错误:", err.Error())
		}
	},
}

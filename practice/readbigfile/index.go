package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	path := "/home/yan/Documents/file/gofile/gotest/middleware/business/logs/minion.log"
	err := ReadFile(path, func(val string) {
		fmt.Println(val)
	})
	fmt.Println(err)
}

// ReadFile 读取文件方法
func ReadFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		lineStr := strings.TrimSpace(string(line))
		handle(lineStr)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

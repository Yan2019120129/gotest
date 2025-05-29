package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// IsExistFile 判断文件路径是否存在
func IsExistFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// MkdirAll 创建全部路径
func MkdirAll(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// MkdirFile 创建目录和文件
func MkdirFile(path string) {
	index := strings.LastIndex(path, "/")
	dir := path[:index]
	if !IsExistFile(path) {
		MkdirAll(dir)
	}
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	file.Close()
}

func readLastNLines(filename string, n int) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 获取文件大小
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := stat.Size()

	lines := make([]string, 0, n)
	lineCount := 0

	// 从文件末尾开始读取
	for offset := int64(1); offset <= size; offset++ {
		// 定位到倒数offset字节处
		_, err := file.Seek(-offset, io.SeekEnd)
		if err != nil {
			break
		}

		// 读取一个字节
		b := make([]byte, 1)
		_, err = file.Read(b)
		if err != nil {
			break
		}

		// 遇到换行符
		if b[0] == '\n' {
			// 读取整行
			_, err := file.Seek(-offset+1, io.SeekEnd)
			if err != nil {
				break
			}

			reader := bufio.NewReader(file)
			line, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				break
			}

			// 添加到结果集（逆序添加）
			lines = append(lines, line)
			lineCount++

			// 达到所需行数
			if lineCount >= n {
				break
			}
		}
	}

	// 处理文件行数不足的情况
	if lineCount < n {
		_, _ = file.Seek(0, io.SeekStart)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	// 反转结果集（因为是从后往前读取）
	reversed := make([]string, len(lines))
	for i, j := 0, len(lines)-1; i < len(lines); i, j = i+1, j-1 {
		reversed[i] = lines[j]
	}

	return reversed, nil
}

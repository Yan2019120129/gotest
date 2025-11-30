package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func PrintFileContent(filePath string, outPath string, limit int, recursive bool) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("无法访问路径: %v", err)
	}

	if info.IsDir() {
		return readDir(filePath, outPath, limit, recursive)
	}

	return readSingleFile(filePath, outPath, limit)
}

// readDir 支持递归读取目录
func readDir(dirPath string, outPath string, limit int, recursive bool) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("无法读取目录: %v", err)
	}

	for _, entry := range entries {
		full := filepath.Join(dirPath, entry.Name())
		if entry.IsDir() {
			if recursive {
				// 递归目录
				if err := readDir(full, outPath, limit, recursive); err != nil {
					fmt.Printf("读取目录失败 %s: %v\n", full, err)
				}
			}
			continue
		}

		if outPath == "" {
			fmt.Printf("\n========== 文件：%s ==========\n", full)
		}

		if err := readSingleFile(full, outPath, limit); err != nil {
			fmt.Printf("读取文件失败 %s: %v\n", full, err)
		}
	}
	return nil
}

// readSingleFile 读取单个文件内容（支持 limit 和 outPath）
func readSingleFile(filePath string, outPath string, limit int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 输出文件对象
	var outFile *os.File
	writeOnly := false // 是否只写文件，不打印

	if outPath != "" {
		writeOnly = true

		// outPath 是目录 → 自动拼接文件名
		if stat, err := os.Stat(outPath); err == nil && stat.IsDir() {
			outPath = filepath.Join(outPath, filepath.Base(filePath))
		}

		outFile, err = os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("无法创建输出文件: %v", err)
		}
		defer outFile.Close()
	}

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		// 如果没有 outPath，则打印到控制台
		if !writeOnly {
			fmt.Println(line)
		}

		// 如果需要写文件
		if outFile != nil {
			outFile.WriteString(line + "\n")
		}

		// 行数限制
		if limit > 0 {
			count++
			if count >= limit {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取文件时发生错误: %v", err)
	}

	return nil
}

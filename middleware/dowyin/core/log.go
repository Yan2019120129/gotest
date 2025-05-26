package core

import (
	"log"
	"os"
	"path/filepath"
)

func LogInit(logDir, logFile string) {
	// 1. 定义日志文件路径（支持自定义路径）
	fullPath := filepath.Join(logDir, logFile)
	// 2. 检查目录是否存在，不存在则创建
	if err := ensureDirExists(logDir); err != nil {
		log.Fatalf("创建日志目录失败: %v", err)
	}

	// 3. 检查文件是否存在，不存在则创建（带权限控制）
	file, err := ensureLogFileExists(fullPath)
	if err != nil {
		log.Fatalf("创建日志文件失败: %v", err)
	}
	defer file.Close()

	// 4. 配置日志输出到文件
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds) // 设置日志格式
}

// ensureDirExists 检查目录是否存在，不存在则创建
func ensureDirExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，创建（递归创建父目录）
		return os.MkdirAll(dirPath, 0755) // 0755 权限：rwxr-xr-x
	}
	return nil
}

// ensureLogFileExists 检查文件是否存在，不存在则创建
func ensureLogFileExists(filePath string) (*os.File, error) {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 文件不存在，创建并设置权限
		return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640) // 0640 权限：rw-r-----
	}
	// 文件已存在，直接打开
	return os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0640)
}

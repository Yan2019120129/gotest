package core

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const maxFileSize = 5 * 1024 * 1024 // 5MB

// OutFile 输出文件内容到控制台或 targetPath
func OutFile(targetPath, path string, n int64) error {
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat failed: %w", err)
	}

	var outputWriter io.Writer = os.Stdout
	var targetFile *os.File

	if targetPath != "" {
		// 打开或创建 targetPath 文件（以追加方式写入）
		targetFile, err = os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("failed to open target file: %w", err)
		}
		defer targetFile.Close()
		outputWriter = targetFile
	}

	if fi.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("read dir failed: %w", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			filePath := filepath.Join(path, entry.Name())
			if _, err := fmt.Fprintf(outputWriter, ">>> [%s] <<<\n", filePath); err != nil {
				return fmt.Errorf("write header failed: %w", err)
			}

			err := processFile(filePath, n, outputWriter)
			if err != nil {
				return fmt.Errorf("error in file %s: %w", filePath, err)
			}
		}
	} else {
		err := processFile(path, n, outputWriter)
		if err != nil {
			return err
		}
	}

	return nil
}

// processFile 执行单个文件的判断与读取，并写入 outputWriter
func processFile(filePath string, n int64, outputWriter io.Writer) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	if info.Size() > maxFileSize {
		return fmt.Errorf("file %s is larger than 5MB", filePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	lines, err := readLines(f, n)
	if err != nil {
		return fmt.Errorf("failed to read lines: %w", err)
	}

	for _, line := range lines {
		if _, err := fmt.Fprintln(outputWriter, line); err != nil {
			return fmt.Errorf("write output failed: %w", err)
		}
	}

	return nil
}

// readLines 根据 n 读取文件中的行
func readLines(reader io.Reader, n int64) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	var allLines []string

	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %w", err)
	}

	total := int64(len(allLines))

	switch {
	case n == 0:
		return allLines, nil
	case n > 0:
		if n > total {
			n = total
		}
		return allLines[:n], nil
	case n < 0:
		n = -n
		if n > total {
			n = total
		}
		return allLines[total-n:], nil
	default:
		return nil, errors.New("invalid n value")
	}
}

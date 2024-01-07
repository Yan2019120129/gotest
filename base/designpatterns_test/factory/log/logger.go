package log

import "fmt"

// Logger 用于创建logger日志
type Logger struct{}

// CreateLog 创建logger日志
func (log *Logger) CreateLog() string {
	fmt.Println("create logger")
	return "logger"
}

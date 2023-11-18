package log

import "fmt"

// Log4j 用于创建log4j日志
type Log4j struct{}

// CreateLog 创建log4j日志
func (log *Log4j) CreateLog() string {
	fmt.Println("create log4j")
	return "log4j"
}

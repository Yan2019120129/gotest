package lumberjack_test

import (
	"github.com/natefinch/lumberjack/v3"
	"log"
)

// NewLumberjack 创建实例
func NewLumberjack() {
	options := &lumberjack.Options{
		MaxAge:     10,
		MaxBackups: 5,
		LocalTime:  false,
		Compress:   true,
	}
	logger, err := lumberjack.NewRoller("app.log", 100, options)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logger)

	log.Println("This is a log entry.")

	logger.Close()
}

package lumberjack_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/natefinch/lumberjack/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
	"time"
)

// NewLumberjack 创建实例
func NewLumberjack() {
	options := &lumberjack.Options{
		MaxAge:     10,
		MaxBackups: 5,
		LocalTime:  true,
		Compress:   false,
	}
	logger, err := lumberjack.NewRoller("app.log", 100, options)
	defer logger.Close()
	if err != nil {
		panic(err)
	}
	log.SetOutput(logger)
	log.Println("This is a log entry.")

	for {
		log.Println("This is a log entry.")
	}

}

// NewLumberjackAndZap 创建实例
func NewLumberjackAndZap() {
	encoderConfig := zapcore.EncoderConfig{}
	var core []zapcore.Core
	core = append(core, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zap.InfoLevel))
	core = append(core, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), getLogWriter("app.log", 1024*1024, 5, 10, true, false), zap.InfoLevel))
	// 自定义一个zap不存在的值，为了实现不打印错误信息堆栈
	var level zapcore.Level = 10
	Logger := zap.New(zapcore.NewTee(
		core...,
	),
		zap.AddStacktrace(level),
		zap.Development(),
		zap.AddCaller(),
	)
	for {
		Logger.Info("test", zap.String("Name", gofakeit.Name()))
	}
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(filename string, maxSize int64, maxBackups int, maxAge time.Duration, IsLocalTime, IsCompress bool) zapcore.WriteSyncer {
	//IsPathExist(filename)
	// filename 文件名
	// maxSize 文件大小
	// Options 文件配置
	roller, err := lumberjack.NewRoller(filename, maxSize, &lumberjack.Options{
		MaxAge:     maxAge,      // 保留旧文件的最大天数
		MaxBackups: maxBackups,  // 保留旧文件的最大个数
		LocalTime:  IsLocalTime, // 是否输出本地时间
		Compress:   IsCompress,  // 是否压缩/归档旧文件
	})
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(roller)
}

// IsPathExist 判断日志文件是否存在，如果不存在则创建路径
func IsPathExist(path string) {
	// 获取日志文件目录
	index := strings.LastIndex(path, "/")
	path = path[:index]

	// 使用 os.Stat 检查目录是否存在
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 目录不存在，使用 os.MkdirAll 创建目录
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Directory created successfully:", path)
	} else if err != nil {
		// 发生其他错误
		fmt.Println("Error checking directory:", err)
		return
	}
}

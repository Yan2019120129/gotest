package logs

import (
	"bandwidth_summary/utils"
	"log"

	"github.com/natefinch/lumberjack"
)

func InitLog(path string, maxSize, maxBackups, maxAge int, compress bool) {
	if !utils.IsExistFile(path) {
		utils.MkdirAll(path)
	}

	// 配置日志轮转
	log.SetOutput(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,    // MB
		MaxBackups: maxBackups, // 保留的旧日志数量
		MaxAge:     maxAge,     // 保留天数
		Compress:   compress,   // 压缩旧日志
	})

	// 设置日志格式
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

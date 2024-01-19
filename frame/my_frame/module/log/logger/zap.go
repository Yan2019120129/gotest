package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"time"
)

var Logger *zap.Logger

// init 初始化zap日志
func init() {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = customColorEncodeLevel
	//config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	//// 调整编码器默认配置
	//config.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	//	encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	//}
	//var err error
	//Logger, err = config.Build()
	//if err != nil {
	//	panic(err)
	//}

	config := zap.Config{
		EncoderConfig: zapcore.EncoderConfig{
			EncodeLevel: customColorEncodeLevel,
			EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
			},
		},
		Level: zap.NewAtomicLevelAt(zap.DebugLevel),
	}
	l, err := config.Build()
	if err != nil {
		panic(err)
	}
	Logger = l
	defer Logger.Sync()
}

// 自定义颜色编码器
func customColorEncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l {
	case zapcore.DebugLevel:
		enc.AppendString("\x1b[35mDEBUG\x1b[0m") // 紫色
	case zapcore.InfoLevel:
		enc.AppendString("\x1b[32mINFO\x1b[0m") // 绿色
	case zapcore.WarnLevel:
		enc.AppendString("\x1b[33mWARN\x1b[0m") // 黄色
	case zapcore.ErrorLevel:
		enc.AppendString("\x1b[31mERROR\x1b[0m") // 红色
	case zapcore.DPanicLevel:
		enc.AppendString("\x1b[31mDPANIC\x1b[0m") // 红色
	case zapcore.PanicLevel:
		enc.AppendString("\x1b[31mPANIC\x1b[0m") // 红色
	case zapcore.FatalLevel:
		enc.AppendString("\x1b[31mFATAL\x1b[0m") // 红色
	}
}

var Instance = &zapLogger{}

// ZapLogger zap日志配置
type zapLogger struct {
	LogLevel logger.LogLevel
}

// LogMode 配置日志模式
func (l *zapLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

// Info 配置info日志
func (l *zapLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	fmt.Println("ctx:", ctx)
	fmt.Println("msg:", msg)
	fmt.Println("data:", data)
	//Logger.Debug(msg, zap.Reflect("", ""))
}

// Warn 配置Warn日志
func (l *zapLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	fmt.Println("ctx:", ctx)
	fmt.Println("msg:", msg)
	fmt.Println("data:", data)
}

// Error 配置Error日志
func (l *zapLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	fmt.Println("ctx:", ctx)
	fmt.Println("msg:", msg)
	fmt.Println("data:", data)
}

// Trace 配置Trace日志
func (l *zapLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	fmt.Println("ctx:", ctx)
	fmt.Println("begin:", begin)
	fmt.Println(l.LogLevel)
	sql, rows := fc()
	fmt.Println("sql:", sql)
	fmt.Println("rows:", rows)
}

package logs

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

var Logger *zap.Logger

// init 初始化zap日志
func init() {
	config := zap.NewDevelopmentConfig()

	// 是否禁用信息文件位置
	//config.DisableCaller = true

	// 开启开发者模式
	config.Development = true

	// 设置对应的日志级别
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	// 给日志级别添加颜色
	config.EncoderConfig.EncodeLevel = customColorEncodeLevel

	// 禁用堆栈，错误和警告不提示上下文关联方法的信息
	config.EncoderConfig.StacktraceKey = ""

	// 添加调用位置信息添加全路径显示
	//config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	// 调整编码器默认配置
	config.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
		//encoder.AppendString(time.Format("2006/01/02 15:04:05.000"))
	}

	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	Logger = l
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

var ErrRecordNotFound = errors.New("record not found")

// ZapLogger zap日志配置
type zapLogger struct {
	logger.Writer
	logger.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode 配置日志模式
func (l *zapLogger) LogMode(level logger.LogLevel) logger.Interface {
	fmt.Println("level:", level)
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info 配置info日志
func (l *zapLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn 配置Warn日志
func (l *zapLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error 配置Error日志
func (l *zapLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace 配置Trace日志
func (l *zapLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			Logger.Error("Gorm", zap.String("err", err.Error()), zap.String("sql", sql))
			//Logger.Error(fmt.Sprintf("%v %v", color.SetRed(err), sql))
		} else {
			Logger.Error("Gorm", zap.String("err", err.Error()), zap.String("sql", sql))
			//Logger.Error(fmt.Sprintf("%v %v", color.SetRed(err), sql))
		}

	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			Logger.Warn("Gorm", zap.String("err", slowLog), zap.String("sql", sql))
		} else {
			Logger.Warn("Gorm", zap.String("err", slowLog), zap.String("sql", sql))
		}

	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			Logger.Info("Gorm", zap.String("sql", sql))
			//Logger.Info(fmt.Sprintf("%v", sql))
		} else {
			Logger.Info("Gorm", zap.String("sql", sql))
			//Logger.Info(fmt.Sprintf("%v", sql))
		}
	}
}

package logs

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"gotest/frame/my_frame/module/logs"
	"time"
)

var LeverGorm = map[string]logger.LogLevel{
	"Debug":  logger.Silent,
	"Info":   logger.Info,
	"Warn":   logger.Warn,
	"Error":  logger.Error,
	"DPanic": logger.Error,
	"Panic":  logger.Error,
	"Fatal":  logger.Error,
}

var Instance = &zapLogger{}

// ZapLogger zap日志配置
type zapLogger struct {
	logger.Writer
	logger.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode 配置日志模式
func (l *zapLogger) LogMode(level logger.LogLevel) logger.Interface {
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
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			logs.Logger.WithOptions(zap.WithCaller(false)).Named(utils.FileWithLineNum()).Error("Gorm", zap.Error(err), zap.String("sql", sql))
		} else {
			logs.Logger.WithOptions(zap.WithCaller(false)).Named(utils.FileWithLineNum()).Error("Gorm", zap.Error(err), zap.String("sql", sql))
		}

	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			logs.Logger.WithOptions(zap.WithCaller(false)).Named(utils.FileWithLineNum()).Warn("Gorm", zap.Error(errors.New(slowLog)), zap.String("sql", sql))
		} else {
			logs.Logger.WithOptions(zap.WithCaller(false)).Named(utils.FileWithLineNum()).Warn("Gorm", zap.Error(errors.New(slowLog)), zap.String("sql", sql))
		}

	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			logs.Logger.WithOptions(zap.WithCaller(false)).Named(utils.FileWithLineNum()).Info("Gorm", zap.String("sql", sql))
		} else {
			logs.Logger.WithOptions(zap.WithCaller(false)).Named(utils.FileWithLineNum()).Info("Gorm", zap.String("sql", sql))
		}
	}
}

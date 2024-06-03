package database

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
	"my-frame/module/gorm/database/mysql"
	"my-frame/module/gorm/database/postgresql"
	"my-frame/module/logs"

	"sync"
	"time"
)

var lever = map[string]logger.LogLevel{
	"Debug":  logger.Silent,
	"Info":   logger.Info,
	"Warn":   logger.Warn,
	"Error":  logger.Error,
	"DPanic": logger.Error,
	"Panic":  logger.Error,
	"Fatal":  logger.Error,
}

// 定义once 保证初始化只执行一次
var _once sync.Once

// DB 定义全局数据库对象
var DB *gorm.DB

// 定义公共数据库配置
var databaseOpen gorm.Dialector

// 获取配置文件数据
var cfg = configs.GetGorm()

// 初始化数据库连接，保证只执行一次
func init() {
	if DB == nil {
		_once.Do(func() {
			InitOpen()
			var err error
			if DB, err = gorm.Open(databaseOpen, &gorm.Config{
				NamingStrategy: schema.NamingStrategy{ // 命名策略
					SingularTable: cfg.SingularTable, // 单表去复数s
				},
				QueryFields: cfg.QueryFields, // 是否全字段映射
				//Logger: logger.Default.LogMode(logger.Info), // 日志级别
				Logger: Instance.LogMode(lever[configs.GetZap().Level]), // 日志级别
			}); err != nil {
				logs.Logger.Error("gorm", zap.String("method", "init"), zap.Error(err))
			}
			logs.Logger.Info("gorm", zap.String("method", "init"), zap.String("instance", fmt.Sprintf("%p", DB)))
		})
	} else {
		logs.Logger.Info("gorm", zap.String("method", "init"), zap.String("instance existed", fmt.Sprintf("%p", DB)))
	}
}

// InitOpen 初始化配置要连接的数据库
func InitOpen() {
	// 选用数据库
	switch cfg.UseDatabase {
	case configs.DatabaseTypePostgresql:
		// 初始化Postgresql数据库
		databaseOpen = postgresql.GetOpen()

	case configs.DatabaseTypeMysql:
		// 初始化mysql数据库
		databaseOpen = mysql.GetOpen()
	}
	return
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
			logs.Logger.WithOptions(zap.WithCaller(false)).Named(utils.FileWithLineNum()).Info("Gorm", zap.Error(err), zap.String("sql", sql))
		} else {
			logs.Logger.WithOptions(zap.WithCaller(false)).Named(utils.FileWithLineNum()).Info("Gorm", zap.Error(err), zap.String("sql", sql))
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

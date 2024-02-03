package logs

import (
	"github.com/natefinch/lumberjack/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gotest/common/config"
	"os"
	"time"
)

var Logger *zap.Logger

const (
	LogMsgFiber    = "fiber"
	LogMsgGorm     = "gorm"
	LogMsgApp      = "app"
	LogMsgOkx      = "okx"
	LogModeCustom  = "custom"
	LogModeProduct = "product"
	LogModeDevel   = "devel"
)

var Lever = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// 时间格式
var formatTime = map[string]func(time time.Time, encoder zapcore.PrimitiveArrayEncoder){
	"long":        longTimeEncoder,
	"short":       shortTimeEncoder,
	"iso8601":     zapcore.ISO8601TimeEncoder,
	"rfc3339":     zapcore.RFC3339TimeEncoder,
	"rfc3339nano": zapcore.RFC3339NanoTimeEncoder,
}

var fileLength = map[string]func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder){
	"full":  zapcore.FullCallerEncoder,
	"short": zapcore.ShortCallerEncoder,
}

// init 初始化zap日志
func init() {
	cfg := config.GetZap()
	encoderConfig := zapcore.EncoderConfig{}
	switch cfg.Mode {

	case LogModeCustom: // 自定义环境
		encoderConfig = customConfig(cfg)

	case LogModeProduct: // 适合在生产环境中使用
		encoderConfig = zap.NewProductionEncoderConfig()

	default: // 适合开发环境使用
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	Logger = zap.New(zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), Lever[cfg.Level]),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), getLogWriter("app.log", 10*1024*1024, 10, 10, true, true), Lever[cfg.Level]),
	),
		zap.AddStacktrace(zap.FatalLevel),
		zap.Development(),
		zap.AddCaller(),
	)
}

func customConfig(cfg *config.ZapConfig) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",   //	输入信息的key名
		LevelKey:       "level", //	输出日志级别的key名
		TimeKey:        "time",  //	输出时间的key名
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,     //	每行的分隔符。"\\n"
		EncodeLevel:    customColorEncodeLevel,        //	将日志级别字符串转化为小写
		EncodeTime:     formatTime[cfg.FormatTime],    //	输出的时间格式
		EncodeDuration: zapcore.StringDurationEncoder, //	执行消耗的时间转化成浮点型的秒
		EncodeCaller:   fileLength[cfg.FileLength],    //	以包/文件:行号 格式化调用堆栈
		EncodeName:     zapcore.FullNameEncoder,       //	可选值。
	}
}

// 自定义时间格式
func longTimeEncoder(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
}

// 自定义时间格式
func shortTimeEncoder(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(time.Format("[" + "15:04:05.000" + "]"))
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
	default:
		enc.AppendString("\x1b[37mDEFAULT\x1b[0m")

	}
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(filename string, maxSize int64, maxBackups int, maxAge time.Duration, IsLocalTime, IsCompress bool) zapcore.WriteSyncer {
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

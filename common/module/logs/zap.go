package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gotest/common/config"
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
	"Full":  zapcore.FullCallerEncoder,
	"Short": zapcore.ShortCallerEncoder,
}

// init 初始化zap日志
func init() {
	cfg := config.GetZap()
	config := zap.Config{}
	switch cfg.Mode {

	case LogModeCustom: // 自定义环境
		config = customConfig(cfg)

	case LogModeProduct: // 适合在生产环境中使用
		config = zap.NewProductionConfig()
		SetConfig(config, cfg)

	default: // 适合开发环境使用
		config = zap.NewDevelopmentConfig()
		SetConfig(config, cfg)
	}

	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	Logger = l
}

// SetConfig  设置配置
func SetConfig(zapConf zap.Config, cfg *config.ZapConfig) {
	// 日志文件输出为位置
	zapConf.OutputPaths = cfg.OutPath
	// 日志内部错误位置
	zapConf.ErrorOutputPaths = cfg.ErrOutPath
	// 日志等级
	zapConf.Level = zap.NewAtomicLevelAt(Lever[cfg.Level])
	// 编码方式，支持json, console
	zapConf.Encoding = cfg.Encoding
	// 输出的时间格式
	zapConf.EncoderConfig.EncodeTime = formatTime[cfg.FormatTime]
	// 文件路径格式，绝对路径和相对路径
	zapConf.EncoderConfig.EncodeCaller = fileLength[cfg.FileLength]
}

func customConfig(cfg *config.ZapConfig) zap.Config {
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(Lever[cfg.Level]), //	日志级别
		Development:       true,                                   //	是否是开发环境。如果是开发模式，对DPanicLevel进行堆栈跟踪
		DisableCaller:     false,                                  //	不显示调用函数的文件名称和行号。默认情况下，所有日志都显示。
		DisableStacktrace: true,                                   //	是否禁用堆栈跟踪捕获。默认对Warn级别以上和生产error级别以上的进行堆栈跟踪。
		Sampling:          nil,                                    //	抽样策略。设置为nil禁用采样。
		Encoding:          cfg.Encoding,                           //	编码方式，支持json, console
		EncoderConfig: zapcore.EncoderConfig{
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
		},
		OutputPaths:      cfg.OutPath,              //	可以配置多个输出路径，路径可以是文件路径和stdout（标准输出）
		ErrorOutputPaths: cfg.ErrOutPath,           //	错误输出路径（日志内部错误）
		InitialFields:    map[string]interface{}{}, //	每条日志中都会输出这些值
	}
	return config
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

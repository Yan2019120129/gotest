package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	config2 "gotest/common/config"
	"time"
)

var Logger *zap.Logger

var lever = map[string]zap.AtomicLevel{
	"Debug":  zap.NewAtomicLevelAt(zapcore.DebugLevel),
	"Info":   zap.NewAtomicLevelAt(zapcore.InfoLevel),
	"Warn":   zap.NewAtomicLevelAt(zapcore.WarnLevel),
	"Error":  zap.NewAtomicLevelAt(zapcore.ErrorLevel),
	"DPanic": zap.NewAtomicLevelAt(zapcore.DPanicLevel),
	"Panic":  zap.NewAtomicLevelAt(zapcore.PanicLevel),
	"Fatal":  zap.NewAtomicLevelAt(zapcore.FatalLevel),
}

var formatTime = map[string]func(time time.Time, encoder zapcore.PrimitiveArrayEncoder){
	"CustomOne":   customTimeEncoder,
	"CustomTwo":   customTimeEncoderTwo,
	"ISO8601":     zapcore.ISO8601TimeEncoder,
	"RFC3339":     zapcore.RFC3339TimeEncoder,
	"RFC3339Nano": zapcore.RFC3339NanoTimeEncoder,
	"Layout":      zapcore.TimeEncoderOfLayout(""),
}

var fileLength = map[string]func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder){
	"Full":  zapcore.FullCallerEncoder,
	"Short": zapcore.ShortCallerEncoder,
}

// init 初始化zap日志
func init() {
	cfg := config2.GetZap()
	config := zap.Config{}
	switch cfg.Mode {

	case "custom": // 自定义环境
		config = customConfig()
		config.Level = lever[cfg.Level]

	case "product": // 适合在生产环境中使用
		config = zap.NewProductionConfig()

	default: // 适合开发环境使用
		config = zap.NewDevelopmentConfig()
	}

	l, err := config.Build()
	if err != nil {
		panic(err)
	}

	Logger = l
}

func customConfig() zap.Config {
	cfg := config2.GetZap()
	//config := zap.NewDevelopmentConfig()
	config := zap.Config{
		Level:             lever[cfg.Level], //	日志级别
		Development:       true,             //	是否是开发环境。如果是开发模式，对DPanicLevel进行堆栈跟踪
		DisableCaller:     false,            //	不显示调用函数的文件名称和行号。默认情况下，所有日志都显示。
		DisableStacktrace: true,             //	是否禁用堆栈跟踪捕获。默认对Warn级别以上和生产error级别以上的进行堆栈跟踪。
		Sampling:          nil,              //	抽样策略。设置为nil禁用采样。
		Encoding:          cfg.Encoding,     //	编码方式，支持json, console
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
		OutputPaths:      []string{"stderr"},       //	可以配置多个输出路径，路径可以是文件路径和stdout（标准输出）
		ErrorOutputPaths: []string{"stderr"},       //	错误输出路径（日志内部错误）
		InitialFields:    map[string]interface{}{}, //	每条日志中都会输出这些值
	}
	return config
}

// 自定义时间格式
func customTimeEncoder(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
}

// 自定义时间格式
func customTimeEncoderTwo(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
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

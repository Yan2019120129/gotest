package main

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 配置日志输出格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		//EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 使用颜色编码的级别输出
		EncodeTime:     customTimeEncoder, // 自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
	}

	// 配置输出到控制台
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleDebugging, zapcore.DebugLevel)

	// 创建 Logger
	logger := zap.New(consoleCore, zap.AddCaller())

	// 打印示例日志
	logger.Info("处理信息", zap.String("tickers", `{"arg":{"channel":"tickers","instId":"1INCH-EUR"},"data":[{"instType":"SPOT","instId":"1INCH-EUR","ast":"0.4364","lastSz":"760.192755","askPx":"0.4398","askSz":"454.855302","bidPx":"0.4389","bidSz":"582.114873","open24h":"0.4438","high24h":"0.4491","low24h":"0.4258","sodUtc0":"0.4424","sodUtc8":"0.4361","volCcy24h":"2987.8446122025","vol24h":"6829.22717","ts":"1704215484272"}]}`))

	// 关闭 Logger
	defer logger.Sync()
}

// 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}

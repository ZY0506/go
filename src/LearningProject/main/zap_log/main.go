package main

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger      *zap.Logger
	sugarlogger *zap.SugaredLogger
)

func main() {
	initLogger()
	defer logger.Sync()
	simpleHttpGet("https://www.baidu.com")
	simpleHttpGet("https://www.sogo.com")
}

func initLogger() {
	wirteSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, wirteSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
	sugarlogger = logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,    // 日志文件最大M
		MaxBackups: 5,     // 最多保留5个日志文件
		MaxAge:     30,    // 最多保留30天
		Compress:   false, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error fetching URL...",
			zap.String("URL", url),
			zap.Error(err))
	} else {
		logger.Info("Successfully fetched URL...",
			zap.String("URL", url),
			zap.Int("StatusCode", resp.StatusCode))
		resp.Body.Close()
	}
}

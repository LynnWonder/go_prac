package config

import (
	"fmt"
	"github.com/LynnWonder/gin_prac/biz/common"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

// 初始化 Logger 并配置好写日志至文件中
func initLogger() {
	// log file
	if err := os.MkdirAll(AppConfig.Log.LogDir, 0755); err != nil {
		panic(fmt.Errorf("create log dir: %s error", AppConfig.Log.LogDir))
	}
	logFileName := filepath.Join(AppConfig.Log.LogDir, AppConfig.Log.LogName)
	fmt.Printf("log to file: %s\n", logFileName)

	// logger
	logLevel, err := zapcore.ParseLevel(AppConfig.Log.Level)
	if err != nil {
		panic(fmt.Errorf("parse log level error, format %s", AppConfig.Log.Level))
	}
	// zap 不支持日志文件归档，因此结合 lumberjack 一起使用
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(getLogWriter(logFileName), os.Stdout)
	common.Logger = zap.New(zapcore.NewCore(getEncoder(), multiWriteSyncer, logLevel), zap.AddCaller())
	zap.ReplaceGlobals(common.Logger)
	common.SugaredLogger = common.Logger.Sugar()

}

func getEncoder() zapcore.Encoder {

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
			zapcore.ShortCallerEncoder(caller, encoder)
			encoder.AppendString("|")
		},
		EncodeLevel: func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
			zapcore.CapitalColorLevelEncoder(level, encoder)
			encoder.AppendString("|")
		},
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006/01/02 15:04:05.000 |"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// 编码器设置
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    128,
		MaxBackups: 7,
		MaxAge:     7,
	}
	// 打印到控制台和文件
	return zapcore.AddSync(lumberJackLogger)
}

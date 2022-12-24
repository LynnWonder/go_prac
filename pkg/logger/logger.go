package logger

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/LynnWonder/gin_prac/pkg/common"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	log "gorm.io/gorm/logger"
)

// Zap是非常快的、结构化的，分日志级别的Go日志库。
var logger *zap.Logger
var callerSkipLogger *zap.Logger

func init() {
	logger = NewZapLogger(common.AppConfig.Log.Level, common.AppConfig.Log.LogName)
	callerSkipLogger = logger.WithOptions(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
}

func Logger() *zap.Logger {
	return logger
}

func SkipLogger() *zap.Logger {
	return callerSkipLogger
}

func NewZapLogger(levelString, filename string) (logger *zap.Logger) {
	var level zapcore.Level
	if err := level.Set(levelString); err != nil {
		fmt.Printf("Failed to set logger level %v", common.AppConfig.Log.Level)
		panic(err)
	}
	core := getEncoder(level, filename)
	return zap.New(core, zap.AddCaller())
}

func getEncoder(level zapcore.Level, filename string) (core zapcore.Core) {
	fileWriter := &lumberjack.Logger{
		MaxSize:    common.AppConfig.Log.MaxSize,
		MaxBackups: common.AppConfig.Log.MaxBackups,
		MaxAge:     common.AppConfig.Log.MaxAge,
		Compress:   common.AppConfig.Log.Compress,
	}
	if _, err := os.Stat(common.AppConfig.Log.LogDir); os.IsNotExist(err) {
		_ = os.MkdirAll(common.AppConfig.Log.LogDir, 0755)
	}

	fileWriter.Filename = path.Join(common.AppConfig.Log.LogDir, filename)
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	if common.AppConfig.Log.Encoder == "json" {
		return zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(fileWriter), level)
	}
	return zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(fileWriter), level)
}

func Debug(msg string, fields ...zap.Field) {
	callerSkipLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	callerSkipLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	callerSkipLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	callerSkipLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	callerSkipLogger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	callerSkipLogger.Panic(msg, fields...)
}

type GormLogger struct {
	logger *zap.Logger
}

func GetGormLogger() *GormLogger {
	return &GormLogger{
		logger: NewZapLogger(common.AppConfig.Log.Level, "db"),
	}
}

func (g *GormLogger) LogMode(level log.LogLevel) log.Interface {
	return g
}

func (g *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	g.logger.Info(fmt.Sprintf(s, i...))
}

func (g *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	g.logger.Warn(fmt.Sprintf(s, i...))
}

func (g *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	g.logger.Error(fmt.Sprintf(s, i...))
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	duration := float64(elapsed.Nanoseconds()) / 1e6
	sql, rows := fc()
	if rows == -1 {
		g.logger.Info(fmt.Sprintf("[Trace] duration %v, sql %v", duration, sql))
	} else {
		g.logger.Info(fmt.Sprintf("[Trace] duration %v, sql %v rows %v", duration, sql, rows))

	}
}

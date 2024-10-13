package main

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

type LogrusAdapter struct {
	logger *logrus.Logger
}

func (l *LogrusAdapter) Info(message string) {
	l.logger.Info(message)
}

func (l *LogrusAdapter) Error(message string) {
	l.logger.Error(message)
}

func NewLogrusAdapter() *LogrusAdapter {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	return &LogrusAdapter{
		logger: logger,
	}
}

type ZapAdapter struct {
	logger *zap.Logger
}

func (z *ZapAdapter) Info(message string) {
	z.logger.Info(message)
}

func (z *ZapAdapter) Error(message string) {
	z.logger.Error(message)
}

func NewZapAdapter() *ZapAdapter {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)

	return &ZapAdapter{
		logger: logger,
	}
}

func main() {
	var logger Logger = NewLogrusAdapter()

	logger.Info("Logrus adapter log info")
	logger.Error("Logrus adapter log error")

	logger = NewZapAdapter()
	logger.Info("Zap adapter log info")
	logger.Error("Zap adapter log error")
}

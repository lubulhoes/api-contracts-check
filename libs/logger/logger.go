package logger

import (
	"log"
	"log/slog"
	"os"
)

var GlobalLogger *logger

type Logger interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
}

var loggerLevel = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

type logger struct {
	logger *slog.Logger
	level  slog.Level
}

func NewLogger(level string) {
	levelVar := &slog.LevelVar{}
	levelVar.Set(loggerLevel[level])

	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: levelVar})

	GlobalLogger = &logger{
		logger: slog.New(handler),
		level:  loggerLevel[level],
	}
}

func (l logger) Info(msg string, args ...any) {
	l.Info(msg, args...)
}

func (l logger) Warn(msg string, args ...any) {
	l.Warn(msg, args...)
}

func (l logger) Error(msg string, args ...any) {
	l.Error(msg, args...)
}

func (l logger) Debug(msg string, args ...any) {
	l.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	if GlobalLogger != nil {
		GlobalLogger.Info(msg, args...)
	} else {
		slog.Info(msg, args...)
	}
}

func Warn(msg string, args ...any) {
	if GlobalLogger != nil {
		GlobalLogger.Warn(msg, args...)
	} else {
		slog.Warn(msg, args...)
	}
}

func Error(msg string, args ...any) {
	if GlobalLogger != nil {
		GlobalLogger.Error(msg, args...)
	} else {
		slog.Error(msg, args...)
	}
}

func Debug(msg string, args ...any) {
	if GlobalLogger != nil {
		GlobalLogger.Debug(msg, args...)
	} else {
		slog.Debug(msg, args...)
	}
}

func Fatal(args ...any) {
	log.Fatal(args...)
}

package logger

import (
	"log/slog"
	"os"
)

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

func NewLogger(level string) Logger {
	levelVar := &slog.LevelVar{}
	levelVar.Set(loggerLevel[level])

	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: levelVar})
	return &logger{logger: slog.New(handler)}
}

func (l logger) Info(msg string, args ...any) {}

func (l logger) Warn(msg string, args ...any) {}

func (l logger) Error(msg string, args ...any) {}

func (l logger) Debug(msg string, args ...any) {}

package logger

import (
	"io"
	"log"
	"log/slog"
)

type SlogLogger struct {
	log *slog.Logger
}

func NewSlogLogger() Log {
	sl := slog.New(slog.NewJSONHandler(io.Discard, nil))

	return &SlogLogger{log: sl}
}

func (zl *SlogLogger) Debug(msg string, keysAndValues ...interface{}) {
	zl.log.Debug(msg, keysAndValues...)
}

func (zl *SlogLogger) Info(msg string, keysAndValues ...interface{}) {
	zl.log.Info(msg, keysAndValues...)
}

func (zl *SlogLogger) Warn(msg string, keysAndValues ...interface{}) {
	zl.log.Warn(msg, keysAndValues...)
}

func (zl *SlogLogger) Error(msg string, keysAndValues ...interface{}) {
	zl.log.Error(msg, keysAndValues...)
}

func (zl *SlogLogger) Fatal(msg string, keysAndValues ...interface{}) {
	zl.log.Error(msg, keysAndValues...)
	log.Fatal(msg)
}

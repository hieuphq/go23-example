package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	log zerolog.Logger
}

func NewZeroLogger() Log {
	zl := zerolog.New(io.Discard).With().Timestamp().Logger()
	return &ZeroLogger{log: zl}
}

func (zl *ZeroLogger) Debug(msg string, keysAndValues ...interface{}) {
	zl.log.Debug().Fields(keysAndValues).Msg(msg)
}

func (zl *ZeroLogger) Info(msg string, keysAndValues ...interface{}) {
	zl.log.Info().Fields(keysAndValues).Msg(msg)
}

func (zl *ZeroLogger) Warn(msg string, keysAndValues ...interface{}) {
	zl.log.Warn().Fields(keysAndValues).Msg(msg)
}

func (zl *ZeroLogger) Error(msg string, keysAndValues ...interface{}) {
	zl.log.Error().Fields(keysAndValues).Msg(msg)
}

func (zl *ZeroLogger) Fatal(msg string, keysAndValues ...interface{}) {
	zl.log.Fatal().Fields(keysAndValues).Msg(msg)
}

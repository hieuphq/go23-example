package logger

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	log *zap.SugaredLogger
}

// A Syncer is a spy for the Sync portion of zapcore.WriteSyncer.
type Syncer struct {
	err    error
	called bool
}

// SetError sets the error that the Sync method will return.
func (s *Syncer) SetError(err error) {
	s.err = err
}

// Sync records that it was called, then returns the user-supplied error (if
// any).
func (s *Syncer) Sync() error {
	s.called = true
	return s.err
}

// Called reports whether the Sync method was called.
func (s *Syncer) Called() bool {
	return s.called
}

// A Discarder sends all writes to io.Discard.
type Discarder struct{ Syncer }

// Write implements io.Writer.
func (d *Discarder) Write(b []byte) (int, error) {
	return io.Discard.Write(b)
}

func NewZapLogger() Log {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	logger := zap.New(zapcore.NewCore(
		enc,
		&Discarder{},
		zapcore.DebugLevel,
	))

	sugar := logger.Sugar()
	return &ZapLogger{log: sugar}
}

func (zl *ZapLogger) Debug(msg string, keysAndValues ...interface{}) {
	zl.log.Debugw(msg, keysAndValues...)
}

func (zl *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	zl.log.Infow(msg, keysAndValues...)
}

func (zl *ZapLogger) Warn(msg string, keysAndValues ...interface{}) {
	zl.log.Warnw(msg, keysAndValues...)
}

func (zl *ZapLogger) Error(msg string, keysAndValues ...interface{}) {
	zl.log.Errorw(msg, keysAndValues...)
}

func (zl *ZapLogger) Fatal(msg string, keysAndValues ...interface{}) {
	zl.log.Fatalw(msg, keysAndValues...)
}

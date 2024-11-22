package log

import (
	"go.uber.org/zap/zapcore"
	"io"
)

type Logger interface {
	With(args ...any) Logger
	WithError(err error) Logger

	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)

	Debugf(template string, args ...any)
	Infof(template string, args ...any)
	Warnf(template string, args ...any)
	Errorf(template string, args ...any)
	Fatalf(template string, args ...any)

	NewNameLogger(componentName string) Logger

	SetLoggerDb(w io.Writer)
}

type Level zapcore.Level

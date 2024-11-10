package configuration

import (
	"gateway/pkg/log"
	"gateway/pkg/log/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logLevel zapcore.Level, filename string) log.Logger {
	switch logLevel {
	case zapcore.DebugLevel:
		return zap.NewDebugLogger(logLevel, filename)
	default:
		return zap.NewLogger(logLevel, filename)
	}
}

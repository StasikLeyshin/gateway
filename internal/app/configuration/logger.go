package configuration

import (
	"gateway/pkg/log"
	"gateway/pkg/log/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(loggerName string, logLevel zapcore.Level, filename string) log.Logger {
	return zap.NewLogger(loggerName, logLevel, filename)
}

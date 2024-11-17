package zap

import (
	"fmt"
	"gateway/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

const errKey = "error"

type LoggerImpl struct {
	logger   *zap.SugaredLogger
	logLevel zapcore.Level
}

func NewDebugLogger(logLevel zapcore.Level, filename string) log.Logger {
	//encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "datetime",
		LevelKey:    "level",
		NameKey:     "logger",
		CallerKey:   "caller",
		FunctionKey: "func",
		MessageKey:  "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	//encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderCfg.EncodeCaller = zapcore.FullCallerEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
		zapcore.NewCore(fileEncoder, file, logLevel),
	)

	logger := zap.New(
		core,
		//zap.Fields(zap.String("name", loggerName)),
		zap.WithCaller(true),
		zap.AddCallerSkip(1),
		//zap.AddStacktrace(logLevel),
	).Sugar()

	return &LoggerImpl{
		logger:   logger,
		logLevel: logLevel,
	}
}

func NewLogger(logLevel zapcore.Level, filename string) log.Logger {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "datetime",
		LevelKey:    "level",
		NameKey:     "logger",
		CallerKey:   "caller",
		FunctionKey: "func",
		MessageKey:  "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
		zapcore.NewCore(fileEncoder, file, logLevel),
	)

	logger := zap.New(
		core,
		//zap.Fields(zap.String("name", loggerName)),
		zap.WithCaller(true),
		zap.AddCallerSkip(1),
		//zap.AddStacktrace(logLevel),
	).Sugar()

	return &LoggerImpl{
		logger:   logger,
		logLevel: logLevel,
	}
}

func (log *LoggerImpl) SetLoggerDb(w io.Writer) {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "datetime",
		LevelKey:    "level",
		NameKey:     "logger",
		CallerKey:   "caller",
		FunctionKey: "func",
		MessageKey:  "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(w), log.logLevel),
		log.logger.Desugar().Core(),
	)

	logger := zap.New(
		core,
		//zap.Fields(zap.String("name", loggerName)),
		zap.WithCaller(true),
		zap.AddCallerSkip(1),
		//zap.AddStacktrace(logLevel),
	).Sugar()

	log.logger = logger

	fmt.Println(log)
}

func (log *LoggerImpl) NewNameLogger(componentName string) log.Logger {
	log.logger = log.logger.Named(componentName)
	return log
}

func (log *LoggerImpl) With(args ...any) log.Logger {
	return &LoggerImpl{
		logger: log.logger.With(args...),
	}
}

func (log *LoggerImpl) WithError(err error) log.Logger {
	return log.With(errKey, err)
}

func (log *LoggerImpl) Debug(args ...any) {
	log.logger.Debug(args...)
}

func (log *LoggerImpl) Info(args ...any) {
	log.logger.Info(args...)
}

func (log *LoggerImpl) Warn(args ...any) {
	log.logger.Warn(args...)
}

func (log *LoggerImpl) Error(args ...any) {
	log.logger.Error(args...)
}

func (log *LoggerImpl) Fatal(args ...any) {
	log.logger.Fatal(args...)
}

func (log *LoggerImpl) Debugf(template string, args ...any) {
	log.logger.Debugf(template, args...)
}

func (log *LoggerImpl) Infof(template string, args ...any) {
	log.logger.Infof(template, args...)
}

func (log *LoggerImpl) Warnf(template string, args ...any) {
	log.logger.Warnf(template, args...)
}

func (log *LoggerImpl) Errorf(template string, args ...any) {
	log.logger.Errorf(template, args...)
}

func (log *LoggerImpl) Fatalf(template string, args ...any) {
	log.logger.Fatalf(template, args...)
}

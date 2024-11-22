package zap

import (
	"gateway/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

const (
	errKey  = "error"
	nameKey = "logger"
)

type (
	Config struct {
		logLevel zapcore.Level
		writer   io.Writer
	}

	Logger struct {
		logger   *zap.SugaredLogger
		logLevel zapcore.Level
	}

	LoggerImpl struct {
		loggerImpl *Logger
		name       string
	}
)

func NewDebugLogger(logLevel zapcore.Level, filename string) log.Logger {
	//encoderCfg := zap.NewProductionEncoderConfig()
	encoderConsoleCfg := zapcore.EncoderConfig{
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

	encoderFileCfg := zapcore.EncoderConfig{
		TimeKey:     "datetime",
		LevelKey:    "level",
		NameKey:     "logger",
		CallerKey:   "caller",
		FunctionKey: "func",
		MessageKey:  "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConsoleCfg)
	fileEncoder := zapcore.NewJSONEncoder(encoderFileCfg)

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
		loggerImpl: &Logger{
			logger:   logger,
			logLevel: logLevel,
		},
	}
}

func NewLogger(logLevel zapcore.Level, filename string) log.Logger {
	encoderConsoleCfg := zapcore.EncoderConfig{
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

	encoderFileCfg := zapcore.EncoderConfig{
		TimeKey:     "datetime",
		LevelKey:    "level",
		NameKey:     "logger",
		CallerKey:   "caller",
		FunctionKey: "func",
		MessageKey:  "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConsoleCfg)
	fileEncoder := zapcore.NewJSONEncoder(encoderFileCfg)

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
		loggerImpl: &Logger{
			logger:   logger,
			logLevel: logLevel,
		},
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
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(w), log.loggerImpl.logLevel),
		log.loggerImpl.logger.Desugar().Core(),
	)

	logger := zap.New(
		core,
		//zap.Fields(zap.String("name", loggerName)),
		zap.WithCaller(true),
		zap.AddCallerSkip(1),
		//zap.AddStacktrace(logLevel),
	).Sugar()

	log.loggerImpl.logger = logger
}

func (log *LoggerImpl) NewNameLogger(componentName string) log.Logger {

	return &LoggerImpl{loggerImpl: log.loggerImpl,
		name: componentName,
	}
}

func (log *LoggerImpl) WithName(args ...any) *LoggerImpl {
	return &LoggerImpl{
		loggerImpl: &Logger{
			logger:   log.loggerImpl.logger.With(args...),
			logLevel: log.loggerImpl.logLevel,
		},
	}
}

func (log *LoggerImpl) With(args ...any) log.Logger {
	return &LoggerImpl{
		loggerImpl: &Logger{
			logger:   log.loggerImpl.logger.With(args...),
			logLevel: log.loggerImpl.logLevel,
		},
	}
}

func (log *LoggerImpl) WithError(err error) log.Logger {
	return log.With(errKey, err)
}

func (log *LoggerImpl) Debug(args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Debug(args...)
}

func (log *LoggerImpl) Info(args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Info(args...)
}

func (log *LoggerImpl) Warn(args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Warn(args...)
}

func (log *LoggerImpl) Error(args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Error(args...)
}

func (log *LoggerImpl) Fatal(args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Fatal(args...)
}

func (log *LoggerImpl) Debugf(template string, args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Debugf(template, args...)
}

func (log *LoggerImpl) Infof(template string, args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Infof(template, args...)
}

func (log *LoggerImpl) Warnf(template string, args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Warnf(template, args...)
}

func (log *LoggerImpl) Errorf(template string, args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Errorf(template, args...)
}

func (log *LoggerImpl) Fatalf(template string, args ...any) {
	log.WithName(nameKey, log.name).loggerImpl.logger.Fatalf(template, args...)
}

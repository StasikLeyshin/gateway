package app

import (
	"context"
	"gateway/internal/app/configuration"
	"gateway/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const (
	configFolder = "config"
	logFolder    = "log"
)

type Application struct {
	serviceProvider *serviceProvider
	logger          log.Logger
	configPath      string
	loggerPath      string

	levelLogger zapcore.Level
}

func NewApp(ctx context.Context) (*Application, error) {
	app := &Application{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *Application) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initENV,
		a.initLogger,
		a.initServiceProvider,
		a.initConfig,
		a.initGRPCServer,
		a.initComponents,
		a.addComponents,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Application) initENV(_ context.Context) error {
	var deployFolder string

	configFileName := "config.local.yaml"

	levelLogger := zap.DebugLevel
	fileNameLogger := "app.log"

	if IsDeploy := os.Getenv("LOGGER_LEVEL"); IsDeploy != "" {
		levelLogger = zap.InfoLevel
	}

	if path := os.Getenv("DEPLOY_PATH"); path != "" {
		configFileName = "config.yaml"
		deployFolder = path

		// Нужно для того, чтобы успела скопироваться папка с конфигами yaml в volume k8s
		time.Sleep(time.Second * 1)
	}

	// Путь до файла с конфигурацией проекта
	a.configPath = filepath.Join(deployFolder, configFolder, configFileName)

	// Путь до файла с логами
	a.loggerPath = filepath.Join(deployFolder, logFolder, fileNameLogger)

	a.levelLogger = levelLogger

	return nil
}

func (a *Application) initLogger(_ context.Context) error {
	// Создаём логгер
	//a.logger = zapImpl.NewLogger()
	a.logger = configuration.NewLogger(a.levelLogger, a.loggerPath)

	return nil
}

func (a *Application) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.logger.NewNameLogger(ComponentNameServiceProvider.String()))
	return nil
}

func (a *Application) initConfig(_ context.Context) error {
	err := a.serviceProvider.initConfig(a.configPath)
	if err != nil {
		return err
	}
	return nil
}

func (a *Application) initGRPCServer(_ context.Context) error {
	a.serviceProvider.initGRPCServer()

	return nil
}

func (a *Application) initComponents(_ context.Context) error {
	a.serviceProvider.initComponents()

	return nil
}

func (a *Application) addComponents(_ context.Context) error {
	err := a.serviceProvider.addComponents()
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) initDbLogger(_ context.Context) error {
	a.logger.SetLoggerDb(a.serviceProvider.LogService())

	return nil
}

func (a *Application) initPostConfigure(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDbLogger,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Application) Run(ctx context.Context) {
	componentsCtx, componentsStopCtx := signal.NotifyContext(ctx, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer componentsStopCtx()

	a.serviceProvider.InjectCallbacks()

	errs := a.serviceProvider.components.Configure(ctx, a.serviceProvider.config)
	if len(errs) > 0 {
		a.logger.Error("Error: ", errs)
	}

	err := a.initPostConfigure(ctx)
	if err != nil {
		a.logger.Error("Error: ", err)
	}

	errs = a.serviceProvider.components.Start(ctx)
	if len(errs) > 0 {
		a.logger.Error("Error: ", errs)
	}

	<-componentsCtx.Done()

	a.serviceProvider.components.Stop(ctx)
}

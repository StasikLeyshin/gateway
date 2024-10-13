package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

type Application struct {
	serviceProvider *serviceProvider
	logger          *logrus.Logger
	configPath      string
}

type component interface {
	Start() error
	Stop(ctx context.Context) error
}

func NewApp(ctx context.Context, configPath string, logger *logrus.Logger) (*Application, error) {
	app := &Application{
		configPath: configPath,
		logger:     logger,
	}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *Application) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
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

func (a *Application) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.logger)
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

func (a *Application) Run1(ctx context.Context) {
	componentsCtx, componentsStopCtx := signal.NotifyContext(ctx, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer componentsStopCtx()

	a.serviceProvider.InjectCallbacks()

	err := a.serviceProvider.components.Configure(ctx, a.serviceProvider.config)
	if err != nil {
		a.logger.Error("Error: ", err)
	}

	a.serviceProvider.components.Start(ctx)

	<-componentsCtx.Done()

	a.serviceProvider.components.Stop(ctx)
}

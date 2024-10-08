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
	//components      []component
	configPath string
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

	a.serviceProvider.components.Start(ctx)

	<-componentsCtx.Done()

	a.serviceProvider.components.Stop(ctx)
}

//func (a *App) Run(ctx context.Context) {
//	componentsCtx, componentsStopCtx := signal.NotifyContext(ctx, syscall.SIGHUP,
//		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
//	defer componentsStopCtx()
//
//	for _, comp := range a.components {
//		err := comp.Start()
//		if err != nil {
//			a.logger.Printf("error when starting the component %v", err)
//		}
//	}
//
//	<-componentsCtx.Done()
//
//	for _, comp := range a.components {
//		err := comp.Stop(ctx)
//		if err != nil {
//			a.logger.Printf("error when stopping the component %v", err)
//		}
//	}
//
//}

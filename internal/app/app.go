package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

type App struct {
	serviceProvider *serviceProvider
	logger          *logrus.Logger
	components      []component
	configPath      string
}

type component interface {
	Start() error
	Stop(ctx context.Context) error
}

func NewApp(ctx context.Context, configPath string, logger *logrus.Logger) (*App, error) {
	app := &App{
		configPath: configPath,
		logger:     logger,
	}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initConfig,
		a.initGRPCServer,
		a.initComponents,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.logger)
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := a.serviceProvider.initConfig(a.configPath)
	if err != nil {
		return err
	}
	return nil
}

//func NewApp(logger *logrus.Logger, components ...component) *App {
//	return &App{
//		logger:     logger,
//		components: components,
//	}
//}

//func Initialization(service *service.GlobalService) map[any]any {
//	registrations := make(map[any]any)
//
//	registrations[roleServiceGRPC.NewImplementationRoleService(roleService.NewRoleService(service))] = desc.RegisterRoleServiceServer
//	return registrations
//}

//func (a *App) initModule(_ context.Context) {
//	a.serviceProvider.GlobalService()
//	return
//}

func (a *App) initGRPCServer(_ context.Context) error {
	a.serviceProvider.initGRPCServer()

	return nil
}

func (a *App) initComponents(_ context.Context) error {
	a.components = a.serviceProvider.initComponents()

	return nil
}

//func (a *App) Run1(ctx context.Context) {
//	components1 := components.Components[configuration.Config]{}
//
//	for _, comp := range a.components {
//		components.AddComponent(components1, comp)
//	}
//}

func (a *App) Run(ctx context.Context) {
	componentsCtx, componentsStopCtx := signal.NotifyContext(ctx, syscall.SIGHUP,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer componentsStopCtx()

	for _, comp := range a.components {
		err := comp.Start()
		if err != nil {
			a.logger.Printf("error when starting the component %v", err)
		}
	}

	<-componentsCtx.Done()

	for _, comp := range a.components {
		err := comp.Stop(ctx)
		if err != nil {
			a.logger.Printf("error when stopping the component %v", err)
		}
	}

}

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
}

type component interface {
	Start() error
	Stop(ctx context.Context) error
}

func NewApp(ctx context.Context, logger *logrus.Logger) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return &App{}, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
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

func Initialization() {
	return
}

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

package app

import (
	"context"
	roleServiceGRPC "gateway/internal/api/grpc-gateway/role-service"
	roleService "gateway/internal/service/role-service"
	"gateway/internal/service/service"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

type App struct {
	logger     *logrus.Logger
	components []component
}

type component interface {
	Start() error
	Stop(ctx context.Context) error
}

func NewApp(logger *logrus.Logger, components ...component) *App {
	return &App{
		logger:     logger,
		components: components,
	}
}

func Initialization(service *service.GlobalService) map[any]any {
	registrations := make(map[any]any)

	registrations[roleServiceGRPC.NewImplementationRoleService(roleService.NewRoleService(service))] = desc.RegisterRoleServiceServer
	return registrations
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

package app

import (
	"context"
	roleServiceGRPC "gateway/internal/api/grpc-gateway/role-service"
	"gateway/internal/app/configuration"
	server "gateway/internal/server/grpc"
	roleService "gateway/internal/service/role-service"
	"gateway/internal/service/service"
	"gateway/internal/transfer"
	"gateway/pkg/components"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"time"
)

type serviceProvider struct {
	logger *logrus.Logger

	components *components.Components[*configuration.Config]

	config  *configuration.Config
	service *service.GlobalService

	grpcServer *grpc.Server

	transfer transfer.Transfer
}

func newServiceProvider(logger *logrus.Logger) *serviceProvider {
	return &serviceProvider{
		logger: logger,
	}
}

func (s *serviceProvider) initConfig(configPath string) error {
	if s.config == nil {
		config, err := configuration.NewConfig(configPath)
		if err != nil {
			return err
		}

		s.config = config
	}

	return nil
}

func (s *serviceProvider) Transfer() transfer.Transfer {
	if s.transfer == nil {
		s.transfer = transfer.NewTransfer()
	}

	return s.transfer
}

func (s *serviceProvider) GlobalService() *service.GlobalService {
	if s.service == nil {
		s.service = service.NewGlobalService(
			s.Transfer(),
		)
	}

	return s.service
}

func (s *serviceProvider) RoleServiceImpl() *roleServiceGRPC.RoleService {
	//if s.userImpl == nil {
	//	s.userImpl = roleServiceGRPC.NewImplementationRoleService(roleService.NewRoleService(s.GlobalService()))
	//}

	return roleServiceGRPC.NewImplementationRoleService(roleService.NewRoleService(s.GlobalService())) //s.userImpl
}

func (s *serviceProvider) Registration(reg grpc.ServiceRegistrar) {
	desc.RegisterRoleServiceServer(reg, s.RoleServiceImpl())
}

func (s *serviceProvider) initGRPCServer() *grpc.Server {
	if s.grpcServer == nil {
		s.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

		reflection.Register(s.grpcServer)

		s.Registration(s.grpcServer)
	}

	return s.grpcServer
}

func (s *serviceProvider) initComponents() {
	if s.components == nil {
		s.components = components.NewComponents[*configuration.Config]()
	}
}

func (s *serviceProvider) addComponents() error {
	components.AddComponent(s.components,
		server.NewServerGRPC(s.grpcServer, s.logger),
		(*configuration.Config).GetGrpcConfig,
		"GRPC Server") // TODO: Вынести имена компонентов в файл с константами

	return nil
}

func (s *serviceProvider) Start(ctx context.Context) error {
	s.components.Start(ctx)

	return nil
}

func (s *serviceProvider) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute) // TODO: Вынести в конфиг
	defer cancel()

	s.components.Stop(ctx)

	return nil
}

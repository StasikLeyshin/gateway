package app

import (
	"context"
	roleServiceGRPC "gateway/internal/api/grpc-gateway/role"
	"gateway/internal/app/configuration"
	"gateway/internal/repository"
	"gateway/internal/repository/database/mongo"
	"gateway/internal/repository/transfer"
	"gateway/internal/repository/transfer/connector"
	server "gateway/internal/server/grpc"
	serviceInterface "gateway/internal/service"
	"gateway/internal/service/service"
	"gateway/pkg/components"
	"gateway/pkg/log"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"time"
)

type serviceProvider struct {
	logger log.Logger

	components *components.Components[*configuration.Config]

	config *configuration.Config

	service *service.Service

	grpcServer *grpc.Server

	transfer transfer.Transfer

	role serviceInterface.RoleService

	roleImpl *roleServiceGRPC.Implementation

	connector *connector.Connector

	mongoClient *mongo.Client

	fileLog repository.LogRepository

	logService serviceInterface.LogService
}

func newServiceProvider(logger log.Logger) *serviceProvider {
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

func (s *serviceProvider) Registration(reg grpc.ServiceRegistrar) {
	desc.RegisterRoleServiceServer(reg, s.RoleImpl())
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
		server.NewServerGRPC(
			s.grpcServer,
			s.logger.NewNameLogger(ComponentNameGRPCServer.String()),
		),
		(*configuration.Config).GetGrpcConfig,
		ComponentNameGRPCServer.String(),
	)

	components.AddComponent(
		s.components,
		s.Connector(),
		(*configuration.Config).GetConnectorConfig,
		ComponentNameConnector.String(),
	)

	components.AddComponent(
		s.components,
		s.Mongo(),
		(*configuration.Config).GetMongoConfig,
		ComponentNameRepositoryMongo.String(),
	)

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

func (s *serviceProvider) InjectCallbacks() {
	s.service.Inject(&service.AppCallbacks{
		RoleService: s.role,
	})
}

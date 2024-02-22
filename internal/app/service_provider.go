package app

import (
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
)

type serviceProvider struct {
	logger *logrus.Logger

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

//func (s *serviceProvider) initGRPCServer1() *grpc.Server {
//	if s.serverGrpc == nil {
//		s.serverGrpc = grpcServer.NewServerGRPC(config.GrpcConfig.Port, serviceClient, logger)
//	}
//
//	s.serverGrpc =
//
//	return serverGrpc
//}

func (s *serviceProvider) initComponents() []component {

	//components := []component{
	//	server.NewServerGRPC(s.grpcServer, s.logger),
	//}

	components1 := components.Components[configuration.Config]{}

	components.AddComponent(components1, server.NewServerGRPC(s.grpcServer, s.logger))

	return nil
}

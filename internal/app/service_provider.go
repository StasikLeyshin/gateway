package app

import (
	roleServiceGRPC "gateway/internal/api/grpc-gateway/role-service"
	"gateway/internal/app/startup"
	roleService "gateway/internal/service/role-service"
	"gateway/internal/service/service"
	"gateway/internal/transfer"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type serviceProvider struct {
	config  *startup.Config
	service *service.GlobalService

	transfer transfer.Transfer
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) initConfig(configPath string) error {
	config, err := startup.NewConfig(configPath)
	if err != nil {
		return err
	}

	s.config = config

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

func (s *serviceProvider) RoleServiceImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = roleServiceGRPC.NewImplementationRoleService(s.GlobalService())
	}

	return s.userImpl
}

func (s *serviceProvider) Registration() {
	desc.RegisterRoleServiceServer(s, roleServiceGRPC.NewImplementationRoleService(roleService.NewRoleService(service)))
}

func (s *serviceProvider) initGRPCServer() *grpc.Server {
	//if s.service == nil {
	//	s.service = service.NewGlobalService(
	//		s.Transfer(),
	//	)
	//}
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(grpcServer)

	Registration(grpcServer, service)

	return s.service
}

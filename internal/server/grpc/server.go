package grpc

import (
	"context"
	"fmt"
	roleServiceGRPC "gateway/internal/api/grpc-gateway/role-service"
	roleService "gateway/internal/service/role-service"
	"gateway/internal/service/service"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
)

type Config struct {
	Port int `yaml:"port"`
}

type ServerGRPC struct {
	grpcServer *grpc.Server
	port       int
	logger     *logrus.Logger
}

func Registration(s grpc.ServiceRegistrar, service *service.GlobalService) {
	desc.RegisterRoleServiceServer(s, roleServiceGRPC.NewImplementationRoleService(roleService.NewRoleService(service)))
}

func NewServerGRPC(config Config, service *service.GlobalService, logger *logrus.Logger) *ServerGRPC {

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(grpcServer)

	Registration(grpcServer, service)

	return &ServerGRPC{
		grpcServer: grpcServer,
		//implementationServer: implementationServer,
		port:   config.Port,
		logger: logger,
	}
}

func (s *ServerGRPC) Start() error {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen port %d: %v", s.port, err)
	}
	go func() {
		s.logger.Infof("server is listening the port %d", s.port)
		err = s.grpcServer.Serve(list)
		if err != nil {
			s.logger.WithError(err).Fatalf("fail to serve the server on the port %d", s.port)
		}
	}()

	//ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()
	//
	//mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err = desc.RegisterRoleServiceHandlerFromEndpoint(ctx, mux, "localhost:5000", opts)
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("server listening at 5001")
	//if err := http.ListenAndServe(":5001", mux); err != nil {
	//	panic(err)
	//}

	return nil
}

func (s *ServerGRPC) Stop(ctx context.Context) error {
	s.logger.Info("server is stopping")
	s.grpcServer.Stop()
	return nil
}

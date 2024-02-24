package grpc

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ServerGRPC struct {
	grpcServer *grpc.Server
	config     Config
	logger     *logrus.Logger
}

func (c *Config) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}

func NewServerGRPC(grpcServer *grpc.Server, logger *logrus.Logger) *ServerGRPC {

	//grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	//
	//reflection.Register(grpcServer)
	//
	//Registration(grpcServer, service)

	return &ServerGRPC{
		grpcServer: grpcServer,
		//implementationServer: implementationServer,
		//port:   config.Port,
		logger: logger,
	}
}

func (s *ServerGRPC) Start(ctx context.Context) error {
	addres := s.config.Address()
	list, err := net.Listen("tcp", addres)
	if err != nil {
		return fmt.Errorf("failed to listen addres %s: %v", addres, err)
	}
	go func() {
		s.logger.Infof("server is listening the addres %s", addres)
		err = s.grpcServer.Serve(list)
		if err != nil {
			s.logger.WithError(err).Fatalf("fail to serve the server on the addres %s", addres)
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

func (s *ServerGRPC) Configure(ctx context.Context, config Config) error {
	s.config = config

	return nil
}

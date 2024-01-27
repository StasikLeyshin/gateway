package grpc

import (
	"context"
	"fmt"
	"github.com/StasikLeyshin/grpc-kafka-services/internal/api/server"
	desc "github.com/StasikLeyshin/grpc-kafka-services/pkg/server_v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ServerGRPC struct {
	grpcServer           *grpc.Server
	implementationServer *server.Implementation
	port                 int
	logger               *logrus.Logger
}

func NewServerGRPC(port int, implementationServer *server.Implementation, logger *logrus.Logger) *ServerGRPC {

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(grpcServer)

	desc.RegisterManagerServiceServer(grpcServer, implementationServer)

	return &ServerGRPC{
		grpcServer:           grpcServer,
		implementationServer: implementationServer,
		port:                 port,
		logger:               logger,
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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = desc.RegisterManagerServiceHandlerFromEndpoint(ctx, mux, "localhost:5000", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 5001")
	if err := http.ListenAndServe(":5001", mux); err != nil {
		panic(err)
	}

	return nil
}

func (s *ServerGRPC) Stop(ctx context.Context) error {
	s.logger.Info("server is stopping")
	s.grpcServer.Stop()
	return nil
}

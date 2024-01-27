package manage_server_service

import (
	"gateway/internal/service"
	desc "github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"
)

type Implementation struct {
	desc.UnimplementedManageServiceServer
	serverService service.ServerService
}

func NewImplementation(serverService service.ServerService) *Implementation {
	return &Implementation{
		serverService: serverService,
	}
}

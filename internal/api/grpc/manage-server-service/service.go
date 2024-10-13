package manage_server_service

import (
	"gateway/internal/service"
	desc "github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"
)

type Implementation struct {
	desc.UnimplementedManageServiceServer
	serverService service.RoleService
}

func NewImplementation(serverService service.RoleService) *Implementation {
	return &Implementation{
		serverService: serverService,
	}
}

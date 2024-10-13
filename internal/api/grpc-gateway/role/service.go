package role

import (
	"gateway/internal/service"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/grpc_gateway_pb"
)

type Implementation struct {
	desc.UnimplementedRoleServiceServer
	service service.RoleService
}

func NewImplementation(service service.RoleService) *Implementation {
	return &Implementation{
		service: service,
	}
}

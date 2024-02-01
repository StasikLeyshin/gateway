package role_service

import (
	"gateway/internal/service"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
)

type RoleService struct {
	desc.UnimplementedRoleServiceServer
	service service.RoleService
}

func NewImplementationRoleService(service service.RoleService) *RoleService {
	return &RoleService{
		service: service,
	}
}

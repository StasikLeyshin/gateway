package role_service

import (
	"gateway/internal/service"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
)

type Implementation struct {
	desc.UnimplementedRoleServiceServer
	service service.GlobalService
}

func NewImplementation(service service.GlobalService) *Implementation {
	return &Implementation{
		service: service,
	}
}

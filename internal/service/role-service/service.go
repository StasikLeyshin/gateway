package role_service

import "gateway/internal/service/service"

type roleService struct {
	service *service.GlobalService
}

func NewRoleService(service *service.GlobalService) *roleService {
	return &roleService{
		service: service,
	}
}

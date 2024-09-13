package role

import "gateway/internal/service"
import serviceCore "gateway/internal/service/service"

var _ service.RoleService = (*roleService)(nil)

type roleService struct {
	service *serviceCore.Service
}

func NewRoleService(service *serviceCore.Service) *roleService {
	return &roleService{
		service: service,
	}
}

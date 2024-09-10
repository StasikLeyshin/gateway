package role

import "gateway/internal/service"
import serviceCore "gateway/internal/service/service"

var _ service.RoleService = (*roleService)(nil)

type roleService struct {
	internalService *serviceCore.InternalService
	serv            serviceCore.InternalService
}

func NewRoleService(internalService *serviceCore.InternalService) *roleService {
	return &roleService{
		internalService: internalService,
	}
}

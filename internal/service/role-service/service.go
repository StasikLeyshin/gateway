package role_service

import "gateway/internal/service"

var _ service.RoleService = (*roleService)(nil)

type roleService struct {
	internalService service.InternalService
}

func NewRoleService(internalService service.InternalService) *roleService {
	return &roleService{
		internalService: internalService,
	}
}

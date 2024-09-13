package app

import (
	serviceInterface "gateway/internal/service"
	roleService "gateway/internal/service/role"
	"gateway/internal/service/service"
)

func (s *serviceProvider) Service() *service.Service {
	if s.service == nil {
		s.service = service.NewService(
			nil,
		)
	}

	return s.service
}

func (s *serviceProvider) RoleService() serviceInterface.RoleService {
	if s.role == nil {
		roleService.NewRoleService(s.Service())
	}

	return s.role
}

package app

import (
	serviceInterface "gateway/internal/service"
	roleService "gateway/internal/service/role"
	"gateway/internal/service/service"
)

func (s *serviceProvider) Service() *service.Service {
	if s.service == nil {
		s.service = service.NewService(
			s.Transfer(),
			s.Connector(),
		)
	}

	return s.service
}

// RoleService

func (s *serviceProvider) RoleService() serviceInterface.RoleService {
	if s.role == nil {
		s.role = roleService.NewRoleService(s.Service())
	}

	return s.role
}

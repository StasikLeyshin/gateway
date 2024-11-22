package app

import (
	serviceInterface "gateway/internal/service"
	"gateway/internal/service/log"
	roleService "gateway/internal/service/role"
	"gateway/internal/service/service"
)

func (s *serviceProvider) Service() *service.Service {
	if s.service == nil {
		s.service = service.NewService(
			s.logger.NewNameLogger(ComponentNameService.String()),
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

// LogService

func (s *serviceProvider) LogService() serviceInterface.LogService {
	if s.logService == nil {
		s.logService = log.NewLogService(s.Service())
	}

	return s.logService
}

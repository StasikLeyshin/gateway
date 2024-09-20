package app

import roleServiceGRPC "gateway/internal/api/grpc-gateway/role"

func (s *serviceProvider) RoleImpl() *roleServiceGRPC.Implementation {
	if s.roleImpl == nil {
		s.roleImpl = roleServiceGRPC.NewImplementation(s.RoleService())
	}

	return s.roleImpl
}

package service

import (
	"gateway/internal/service/role-service"
)

type MicroServices interface {
	role_service.RoleService
}

//type GlobalService interface {
//	GetRoleServiceMethods() RoleService
//}

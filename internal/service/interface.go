package service

import role_service "gateway/internal/service/role-service"

type RoleService interface {
	role_service.LoginService
}

type MicroServices interface {
	RoleService
}

//type GlobalService interface {
//	GetRoleServiceMethods() RoleService
//}

//type microServices struct {
//	role_service.RoleService
//}
//
//func re() {
//	microServices1 := microServices{}
//	microServices1.Login()
//}

package service

import (
	"context"
	"gateway/internal/service/role-service/model"
)

type (
	LoginService interface {
		Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
		NewRequest() *model.LoginRequest
	}
)

type RoleService interface {
	LoginService
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

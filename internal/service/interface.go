package service

import (
	"context"
	role_service "gateway/internal/service/role-service"
	"gateway/internal/service/role-service/model"
	"gateway/internal/transfer"
)

type (
	LoginSubService interface {
		Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	}
)

type RoleService interface {
	LoginSubService
}

type InternalService interface {
	//GetExternalServices() ExternalServices
	GetTransfer() transfer.Transfer
}

type internalServices struct {
	role_service.roleService
}

//
//func NewInternalServices[com any](coms ...com) *internalServices {
//	return &internalServices{}
//}

//type InternalServices interface {
//	RoleService
//}

//type Global

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

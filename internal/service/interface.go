package service

import (
	"context"
	"gateway/internal/service/role/model"
)

type (
	LoginSubService interface {
		Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	}
)

type RoleService interface {
	LoginSubService
}

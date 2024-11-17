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

type (
	LogSubService interface {
		Write(p []byte) (n int, err error)
	}
)

type RoleService interface {
	LoginSubService
}

type LogService interface {
	LogSubService
}

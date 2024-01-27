package role_service

import (
	"context"
	"gateway/internal/service"
)

type (
	LoginService interface {
		Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error)
	}
)

// Login

type (
	LoginRequest struct {
	}

	LoginResponse struct {
	}
)

func (l *LoginRequest) Do(ctx context.Context, service *service.GlobalService) (*LoginResponse, error) {
	return nil, nil
}

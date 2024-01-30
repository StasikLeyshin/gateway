package role_service

import (
	"context"
)

type (
	LoginService interface {
		Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error)
		NewRequest() *LoginRequest
	}
)

// Login

type (
	LoginRequest struct {
	}

	LoginResponse struct {
		SessionID string
	}
)

//func (l *LoginRequest) Do(ctx context.Context, service *service.GlobalService) (*LoginResponse, error) {
//	return &LoginResponse{
//		SessionID: "TEST228",
//	}, nil
//}

func (r *roleService) NewRequest() *LoginRequest {
	return &LoginRequest{}
}

func (r *roleService) Login(ctx context.Context, service *LoginRequest) (*LoginResponse, error) {
	return &LoginResponse{
		SessionID: "TEST228",
	}, nil
}

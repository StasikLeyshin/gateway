package service

import (
	"context"
	"gateway/internal/service/role-service"
)

func (g *GlobalService) Login(ctx context.Context, request *role_service.LoginRequest) (*role_service.LoginResponse, error) {
	return request.Do(ctx, g)
}

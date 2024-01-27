package transfer

import (
	"context"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
)

type RoleService interface {
	Login(ctx context.Context, request *desc.LoginRequest) (*desc.LoginResponse, error)
}

type Transfer interface {
	RoleService
}

type transfer struct {
}

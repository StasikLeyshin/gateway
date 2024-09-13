package transfer

import (
	"context"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
)

type RoleTransfer interface {
	Login(ctx context.Context, request *desc.LoginRequest) (*desc.LoginResponse, error)
}

type Transfer interface {
	RoleTransfer
}

type transfer struct {
	role RoleTransfer
}

func NewTransfer() *transfer {
	return &transfer{}
}

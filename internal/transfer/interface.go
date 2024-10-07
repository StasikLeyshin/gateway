package transfer

import (
	"context"
	"gateway/internal/transfer/grpc/role/model"
)

type fromTransferToService[T, R any] interface {
	FromTransfer(value T) R
}

type toTransferFromService[T, R any] interface {
	ToTransfer(value T) R
}

type RoleTransfer interface {
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
}

type Transfer interface {
	RoleTransfer
}

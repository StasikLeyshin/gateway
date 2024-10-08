package transfer

import (
	transfer1 "gateway/internal/transfer"
	"gateway/internal/transfer/grpc/role"
)

type transfer struct {
	transfer1.RoleTransfer
}

func NewTransfer() *transfer {
	return &transfer{
		role.NewTransferRole(),
	}
}

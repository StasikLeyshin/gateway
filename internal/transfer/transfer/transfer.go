package transfer

import (
	"gateway/internal/transfer"
	"gateway/internal/transfer/grpc/role"
)

type Transfer struct {
	transfer.RoleTransfer
}

func NewTransfer() *Transfer {
	return &Transfer{
		role.NewTransferRole(),
	}
}

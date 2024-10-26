package manage_server

import (
	"gateway/internal/transfer"
	"gateway/internal/transfer/grpc/manage-server/model"
)

var _ transfer.ServerManagerTransfer = (*serverManagerTransfer)(nil)

type serverManagerTransfer struct {
	serverTypes []model.ServerType
}

func NewServerManagerTransfer(serverTypes []model.ServerType) *serverManagerTransfer {
	return &serverManagerTransfer{
		serverTypes: serverTypes,
	}
}

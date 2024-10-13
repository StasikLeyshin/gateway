package manage_server

import (
	"gateway/internal/transfer"
	"gateway/internal/transfer/grpc/manage-server/model"
)

var _ transfer.ServerManagerTransfer = (*serverManagerTransfer)(nil)

type serverManagerTransfer struct {
	serverType model.ServerType
}

func NewServerManagerTransfer() *serverManagerTransfer {
	return &serverManagerTransfer{
		serverType: model.ManageServerServerType,
	}
}

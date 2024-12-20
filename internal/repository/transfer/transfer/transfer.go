package transfer

import (
	transfer1 "gateway/internal/repository/transfer"
	"gateway/internal/repository/transfer/grpc/manage-server"
	"gateway/internal/repository/transfer/grpc/role"
)

type transfer struct {
	transfer1.RoleTransfer
	transfer1.ServerManagerTransfer
}

func NewTransfer() *transfer {
	return &transfer{
		RoleTransfer:          role.NewRoleTransfer(),
		ServerManagerTransfer: manage_server.NewServerManagerTransfer(nil), // Сервер менеджер сам будет знать, какие типы
		// proto файлов используются на каждом мс и сообщать об этом
	}
}

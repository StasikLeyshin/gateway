package service

import (
	"gateway/internal/service"
	"gateway/internal/transfer"
)

type InternalGrpcServices struct {
	roleService service.RoleService
}

type AppCallbacks struct {
	RoleService service.RoleService
}

type Service struct {
	internalGrpcServices InternalGrpcServices
	Transfer             transfer.Transfer
	appCallbacks         *AppCallbacks
	connector            transfer.Connector
}

func NewService(transfer transfer.Transfer, connector transfer.Connector) *Service {
	return &Service{
		Transfer:  transfer,
		connector: connector,
	}
}

func (g *Service) Inject(appCallbacks *AppCallbacks) {
	g.appCallbacks = appCallbacks
}

func (g *Service) GetTransfer() transfer.Transfer {
	return g.Transfer
}

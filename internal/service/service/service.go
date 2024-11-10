package service

import (
	"gateway/internal/repository/transfer"
	"gateway/internal/service"
	"gateway/pkg/log"
)

type AppCallbacks struct {
	RoleService service.RoleService
}

type Service struct {
	logger log.Logger

	Transfer  transfer.Transfer
	Services  *AppCallbacks
	connector transfer.Connector
}

func NewService(logger log.Logger, transfer transfer.Transfer, connector transfer.Connector) *Service {
	return &Service{
		logger:    logger,
		Transfer:  transfer,
		connector: connector,
	}
}

func (g *Service) Inject(appCallbacks *AppCallbacks) {
	g.Services = appCallbacks
}

func (g *Service) GetTransfer() transfer.Transfer {
	return g.Transfer
}

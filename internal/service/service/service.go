package service

import (
	"gateway/internal/repository"
	"gateway/internal/repository/transfer"
	"gateway/internal/service"
	"gateway/pkg/log"
)

type AppCallbacks struct {
	RoleService service.RoleService
	LogService  service.LogService
}

type AppRepositoryCallbacks struct {
	LogRepository repository.LogRepository
}

type Service struct {
	logger log.Logger

	Transfer  transfer.Transfer
	connector transfer.Connector

	Services   *AppCallbacks
	Repository *AppRepositoryCallbacks
}

func NewService(logger log.Logger, transfer transfer.Transfer, connector transfer.Connector) *Service {
	return &Service{
		logger:    logger,
		Transfer:  transfer,
		connector: connector,
	}
}

func (g *Service) Inject(appCallbacks *AppCallbacks, appRepositoryCallbacks *AppRepositoryCallbacks) {
	g.Services = appCallbacks
	g.Repository = appRepositoryCallbacks
}

func (g *Service) GetTransfer() transfer.Transfer {
	return g.Transfer
}

package connector

import (
	"context"
	"gateway/internal/repository/transfer"
	"gateway/internal/repository/transfer/clients"
	"gateway/internal/repository/transfer/grpc/manage-server/model"
	"gateway/pkg/log"
	"gateway/pkg/utils"
	"sync/atomic"
)

type Config struct {
	ServerManagerAddress string `yaml:"server_manager_address"`
	NodeID               string `yaml:"node_id"`
}

type (
	Connector struct {
		logger log.Logger

		transfer transfer.Transfer
		config   atomic.Pointer[Config]
		clients  *utils.Safe[model.ServerType, string, *clients.Client]
	}
)

func NewConnector(logger log.Logger, transfer transfer.Transfer) *Connector {
	return &Connector{
		logger: logger,

		transfer: transfer,

		clients: utils.NewSafe[model.ServerType, string, *clients.Client](),
	}
}

func (c *Connector) Configure(ctx context.Context, config *Config) error {
	c.config.Store(config)

	return nil
}

func (c *Connector) Start(ctx context.Context) error {
	c.SyncServersAddress(ctx)

	return nil
}

func (c *Connector) Stop(ctx context.Context) error {
	return nil
}

func (c *Connector) SyncServersAddress(ctx context.Context) {
	response, err := c.transfer.GetServersAddresses(
		ctx,
		&model.GetServersAddressesRequest{
			ServerType: model.ManageServerServerType,
		},
		c.GetManagerServerClient(ctx),
	)

	if err != nil {
	}

	newClients := make(map[model.ServerType]map[string]*clients.Client)

	if response != nil {
		for _, server := range response.Servers {
			if _, ok := newClients[server.ServerType]; !ok {
				newClients[server.ServerType] = make(map[string]*clients.Client)
			}
			newClients[server.ServerType][server.NodeID] = clients.NewClient(ctx, server.Host+server.Port)
		}
	}

	c.clients.SetAll(newClients)
}

func (c *Connector) GetManagerServerClient(ctx context.Context) *clients.Client {
	managerServerCount := c.clients.LenValueByKey(model.ManageServerServerType)

	if managerServerCount == 0 {
		c.clients.Set(model.ManageServerServerType,
			map[string]*clients.Client{
				c.config.Load().NodeID: clients.NewClient(ctx, c.config.Load().ServerManagerAddress),
			},
		)
	}

	client, ok := c.clients.GetValue(model.ManageServerServerType, c.config.Load().NodeID)

	if !ok {

	}

	return client
}

func (c *Connector) GetClientFromServerType(ctx context.Context, serverType model.ServerType) any {
	clientCount := c.clients.LenValueByKey(serverType)

	if clientCount == 0 {
		c.SyncServersAddress(ctx)
	}

	client, ok := c.clients.GetValueByIndex(model.ManageServerServerType, utils.RandRange(0, clientCount))

	if !ok {

	}

	return client
}

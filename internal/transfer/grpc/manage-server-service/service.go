package manage_server_service

import (
	"context"
	desc "github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"
)

type transferManageServerService struct {
}

func NewTransferRoleService() *transferManageServerService {
	return &transferManageServerService{}
}

func (i *transferManageServerService) GetServers(ctx context.Context, request *desc.GetServersRequest) (*desc.GetServersResponse, error) {

	return nil, nil
}

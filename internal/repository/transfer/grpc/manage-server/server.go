package manage_server

import (
	"context"
	"gateway/internal/repository/transfer"
	"gateway/internal/repository/transfer/grpc/manage-server/model"
	"github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"
)

func (s *serverManagerTransfer) GetServers(ctx context.Context, request *model.GetServersRequest, client any) (*model.GetServersResponse, error) {
	response, err := transfer.CallGRPCOrHTTP(ctx, client, request, pb.ManageServiceClient.GetServers, new(model.GetServersResponse))
	return response, err
}

func (s *serverManagerTransfer) GetServersAddresses(ctx context.Context, request *model.GetServersAddressesRequest, client any) (*model.GetServersAddressesResponse, error) {
	//response, err := transfer.CallGRPCorHTTP(ctx, client, request, pb.ManageServiceClient.GetServersAddresses, new(model.GetServersAddressesResponse))
	//return response, err
	return &model.GetServersAddressesResponse{
		Servers: []*model.Server{{
			ServerID:   "0000",
			NodeID:     "1111",
			Name:       "Manager Server",
			Host:       "127.0.0.1",
			Port:       ":8080",
			ServerType: model.ManageServerServerType,
		}},
	}, nil
}

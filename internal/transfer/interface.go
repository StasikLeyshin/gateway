package transfer

import (
	"context"
	"gateway/internal/transfer/clients"
	manage_server "gateway/internal/transfer/grpc/manage-server/model"
	role "gateway/internal/transfer/grpc/role/model"
	"github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"
	"google.golang.org/grpc"
)

type RoleTransfer interface {
	Login(ctx context.Context, request *role.LoginRequest) (*role.LoginResponse, error)
}

type ServerManagerTransfer interface {
	GetServers(ctx context.Context, request *manage_server.GetServersRequest, client any) (*manage_server.GetServersResponse, error)
	GetServersAddresses(ctx context.Context, request *manage_server.GetServersAddressesRequest, client any) (*manage_server.GetServersAddressesResponse, error)
}

type Transfer interface {
	RoleTransfer
	ServerManagerTransfer
}

type (
	Connector interface {
		GetClientFromServerType(ctx context.Context, serverType manage_server.ServerType) any
	}
)

type grpcFunc[C, Req, Resp any] func(client C, ctx context.Context, request Req, opts ...grpc.CallOption) (Resp, error)

type grpcToAnotherServerFunc[Req, Resp any] grpcFunc[pb.ManageServiceClient, Req, Resp]

func CallGRPCorHTTP[
	GrpcRequest any,
	GrpcResponse any,
	TransferResponse interface {
		ToTransfer(GrpcResponse) TransferResponse
	},
	TransferRequest interface {
		FromTransfer() GrpcRequest
	},
](
	ctx context.Context,
	client any,
	request TransferRequest,
	fn grpcToAnotherServerFunc[GrpcRequest, GrpcResponse],
	resp TransferResponse,
) (TransferResponse, error) {

	switch client := client.(type) {
	case *clients.Client:

		grpcRequest := request.FromTransfer()

		serviceResponse, err := fn(client, ctx, grpcRequest)

		if err != nil {
			return *new(TransferResponse), err // TODO: Добавить обработку ошибок
		}

		resp = resp.ToTransfer(serviceResponse)

		return resp, nil
	}

	return resp, nil
}

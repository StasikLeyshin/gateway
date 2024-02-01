package role_service

import (
	"context"
	"gateway/internal/api"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *RoleService) Login(ctx context.Context, request *desc.LoginRequest) (*desc.LoginResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	response, _ := api.CallService(ctx, request, i.service.Login, new(desc.LoginResponse), i.service.NewRequest())

	return response, nil
	//i.service.Do()
	//return service.CallService(ctx, i.service, r.ToCore().Do, new(GetLoginsResponse))

	//
	//result, err := i.serverService.CreateServer(ctx, converter.ToCreateServerRequestFromGrpc(serverRequest))
	//if err != nil {
	//	i.logger.WithError(err).Error(err)
	//	return nil, status.Errorf(codes.Internal, "Internal error")
	//}
	//
	//return converter.ToCreateServerResponseToGrpc(result), nil
}

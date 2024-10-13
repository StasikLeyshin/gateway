package role

import (
	"context"
	"gateway/internal/api"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/grpc_gateway_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, request *desc.LoginRequest) (*desc.LoginResponse, error) {
	if request == nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	response, _ := api.CallService(ctx, request, i.service.Login) //, i.service.NewRequest())

	return response, nil

}

package role

import (
	"context"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
)

type transferRoleService struct {
}

func NewTransferRoleService() *transferRoleService {
	return &transferRoleService{}
}

func (i *transferRoleService) Login(ctx context.Context, request *desc.LoginRequest) (*desc.LoginResponse, error) {

	return nil, nil
}

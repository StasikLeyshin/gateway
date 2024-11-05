package role

import (
	"context"
	"gateway/internal/repository/transfer/grpc/role/model"
)

func (r *roleTransfer) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	//NewClient(ctx, "").RoleServiceClient.Login(ctx, request.FromTransfer())
	return nil, nil
}

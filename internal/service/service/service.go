package service

import (
	"gateway/internal/service"
	"gateway/internal/transfer"
)

type GlobalService struct {
	transfer      transfer.Transfer
	microServices service.MicroServices
}

func NewGlobalService(transfer transfer.Transfer) *GlobalService {
	return &GlobalService{
		transfer: transfer,
	}
}

//func (g *GlobalService) GetServer(ctx context.Context) error {
//	role_service.NewRoleService(g)
//	return nil
//}

//func (g *GlobalService) Login(ctx context.Context, request *role_service.LoginRequest) (*role_service.LoginResponse, error) {
//	return request.Do(ctx, g)
//}

package service

import (
	"gateway/internal/service"
	"gateway/internal/transfer"
)

type InternalGrpcServices struct {
	roleService service.RoleService
}

type AppCallbacks struct {
	RoleService service.RoleService
}

type Service struct {
	internalGrpcServices InternalGrpcServices
	transfer             transfer.Transfer
	appCallbacks         *AppCallbacks
}

func NewService(transfer transfer.Transfer) *Service {
	return &Service{
		//internalGrpcServices: internalGrpcServices,
		transfer: transfer,
	}
}

func (g *Service) Inject(appCallbacks *AppCallbacks) {
	g.appCallbacks = appCallbacks
}

func (g *Service) GetTransfer() transfer.Transfer {
	return g.transfer
}

//func (g *GlobalService) GetServer(ctx context.Context) error {
//	role_service.NewRoleService(g)
//	return nil
//}
//
//func (g *GlobalService) GetServer(ctx context.Context) error {
//	role_service.NewRoleService(g)
//	return nil
//}
//
//func (g *GlobalService) Login(ctx context.Context, request *role_service.LoginRequest) (*role_service.LoginResponse, error) {
//	return request.Do(ctx, g)
//}

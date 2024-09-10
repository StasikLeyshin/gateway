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

type InternalService struct {
	internalGrpcServices InternalGrpcServices
	transfer             transfer.Transfer
	appCallbacks         *AppCallbacks
}

func NewInternalService(transfer transfer.Transfer) *InternalService {
	return &InternalService{
		//internalGrpcServices: internalGrpcServices,
		transfer: transfer,
	}
}

func (g *InternalService) Inject(appCallbacks *AppCallbacks) {
	g.appCallbacks = appCallbacks
}

func (g *InternalService) GetTransfer() transfer.Transfer {
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

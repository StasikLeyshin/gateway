package role

import (
	"context"
	"gateway/internal/service/role/model"
	"gateway/internal/service/service"
)

//type (
//	LoginService interface {
//		Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
//		NewRequest() *model.LoginRequest
//	}
//)

// Login

//type (
//	LoginRequest struct {
//	}
//
//	LoginResponse struct {
//		SessionID string
//	}
//)

//func (l *LoginRequest) Do(ctx context.Context, service *service.GlobalService) (*LoginResponse, error) {
//	return &LoginResponse{
//		SessionID: "TEST228",
//	}, nil
//}

//func (r *roleService) NewLoginRequest() *model.LoginRequest {
//	return &model.LoginRequest{}
//}

func (r *roleService) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	//r.internalService.
	response, _ := r.service.Transfer.Login(ctx, request.FromService())

	response, _ := service.DoTransfer(ctx)
	//c.GetTransfer().Login(ctx, request.FromService())
	return new(model.LoginResponse).response, nil
}

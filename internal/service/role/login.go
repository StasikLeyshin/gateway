package role

import (
	"context"
	"gateway/internal/service/role-service/model"
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
	r.internalService.GetTransfer().Login(ctx, request.FromService())
	//c.GetTransfer().Login(ctx, request.FromService())
	return &model.LoginResponse{
		SessionID: "TEST228",
	}, nil
}

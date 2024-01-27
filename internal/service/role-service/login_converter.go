package role_service

import desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"

func (convert *LoginRequest) FromService(request *desc.LoginRequest) *desc.LoginRequest {
	return nil
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func (convert *LoginRequest) ToService(request *desc.LoginRequest) *LoginRequest {
	return nil
}

func (convert *LoginResponse) FromService(request *desc.LoginResponse) *desc.LoginResponse {
	return nil
}

func (convert *LoginResponse) ToService(request *desc.LoginResponse) *LoginResponse {
	return nil
}

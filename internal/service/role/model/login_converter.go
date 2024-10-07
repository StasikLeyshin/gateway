package model

import (
	"gateway/internal/transfer/grpc/role/model"
	desc "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
	//descInternal "github.com/StasikLeyshin/libs-proto/grpc/role-service/pb"
)

func (convert *LoginRequest) ToService(value *desc.LoginRequest) *LoginRequest {
	return nil
}

func (convert *LoginRequest) FromService() *model.LoginRequest {
	return nil
}

func (convert *LoginResponse) ToService(value *desc.LoginResponse) *LoginResponse {
	return nil
}

func (convert *LoginResponse) FromService() *model.LoginResponse {
	return nil
}

//func (convert *LoginRequest) FromTransfer(value *descInternal.LoginRequest) *descInternal.LoginRequest {
//	return nil
//}
//
//func (convert *LoginRequest) ToTransfer(value *descInternal.LoginRequest) *LoginRequest {
//	return nil
//}
//
//func (convert *LoginResponse) FromTransfer(value *descInternal.LoginResponse) *descInternal.LoginResponse {
//	return nil
//}
//
//func (convert *LoginResponse) ToTransfer(value *descInternal.LoginResponse) *LoginResponse {
//	return nil
//}

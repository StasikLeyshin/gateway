package model

import (
	descExternal "github.com/StasikLeyshin/libs-proto/grpc-gateway/role-service/pb"
)

func (convert *LoginRequest) ToService(value *descExternal.LoginRequest) *LoginRequest {
	return nil
}

func (convert *LoginResponse) ToService(value *descExternal.LoginResponse) *LoginResponse {
	return nil
}

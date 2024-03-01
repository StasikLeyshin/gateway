package model

import (
	descInternal "github.com/StasikLeyshin/libs-proto/grpc/manage-server-service/pb"
)

func (convert *LoginRequest) FromTransfer(value *descInternal.GetServersRequest) *descInternal.LoginRequest {
	return nil
}

func (convert *LoginRequest) ToTransfer(value *descInternal.LoginRequest) *LoginRequest {
	return nil
}

func (convert *LoginResponse) FromTransfer(value *descInternal.LoginResponse) *descInternal.LoginResponse {
	return nil
}

func (convert *LoginResponse) ToTransfer(value *descInternal.LoginResponse) *LoginResponse {
	return nil
}

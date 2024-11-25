package service

import (
	"context"
	log "gateway/internal/service/log/model"
	manageServer "gateway/internal/service/manage-server/model"
	"gateway/internal/service/role/model"
)

type (
	LoginSubService interface {
		Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	}
)

type (
	ManageServerSubService interface {
		GetServers(ctx context.Context, request *manageServer.GetServersRequest) (*manageServer.GetServersResponse, error)
	}
)

type (
	LogSubService interface {
		Write(p []byte) (n int, err error)
		AddLog(ctx context.Context, request *log.AddLogRequest) (*log.AddLogResponse, error)
		GetLog(ctx context.Context, request *log.GetLogRequest) (*log.GetLogResponse, error)
	}
)

type RoleService interface {
	LoginSubService
}

type LogService interface {
	LogSubService
}

type ManageServerService interface {
	ManageServerSubService
}

package repository

import (
	"context"
	"gateway/internal/repository/log/model"
)

type (
	LogRepository interface {
		AddLog(ctx context.Context, request *model.AddLogRequest) (*model.AddLogResponse, error)
		GetLog(ctx context.Context, request *model.GetLogRequest) (*model.GetLogResponse, error)
		GetLogs(ctx context.Context) ([]*model.Log, error)
	}
)

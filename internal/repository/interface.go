package repository

import (
	"context"
	"gateway/internal/repository/log/models"
)

type (
	FileLog interface {
		AddLog(ctx context.Context, log *models.Log) error
		GetLog(ctx context.Context, id string) (*models.Log, error)
		GetLogs(ctx context.Context) ([]*models.Log, error)
	}
)

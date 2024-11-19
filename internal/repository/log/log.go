package log

import (
	"context"
	"gateway/internal/repository/log/model"
)

func (l *logRepository) AddLog(ctx context.Context, request *model.AddLogRequest) (*model.AddLogResponse, error) {
	err := l.client.Database.GetCollection(l.CollectionName).Insert(ctx, request)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (l *logRepository) GetLog(ctx context.Context, request *model.GetLogRequest) (*model.GetLogResponse, error) {

	return nil, nil
}

func (l *logRepository) GetLogs(ctx context.Context) ([]*model.Log, error) {

	return nil, nil
}

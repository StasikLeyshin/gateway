package log

import (
	"context"
	"gateway/internal/repository/database/mongo"
	"gateway/internal/repository/log/models"
)

type (
	fileLog struct {
		CollectionName string
		collection     mongo.Collection
	}
)

func NewFileLog(collectionName string) *fileLog {
	return &fileLog{
		CollectionName: collectionName,
	}
}

func (l *fileLog) AddLog(ctx context.Context, log *models.Log) error {
	err := l.collection.Insert(ctx, log)
	if err != nil {
		return err
	}

	return nil
}

func (l *fileLog) GetLog(ctx context.Context, id string) (*models.Log, error) {

	return nil, nil
}

func (l *fileLog) GetLogs(ctx context.Context) ([]*models.Log, error) {

	return nil, nil
}

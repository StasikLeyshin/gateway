package log

import (
	"gateway/internal/repository/database/mongo"
)

type (
	logRepository struct {
		CollectionName string
		client         *mongo.Client
	}
)

func NewLogRepository(collectionName string, client *mongo.Client) *logRepository {
	return &logRepository{
		CollectionName: collectionName,
		client:         client,
	}
}

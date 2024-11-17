package log

import (
	"gateway/internal/repository/database/mongo"
)

type (
	logRepository struct {
		CollectionName string
		collection     *mongo.Database
	}
)

func NewLogRepository(collectionName string, collection *mongo.Database) *logRepository {
	return &logRepository{
		CollectionName: collectionName,
		collection:     collection,
	}
}

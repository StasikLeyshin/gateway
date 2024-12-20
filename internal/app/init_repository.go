package app

import (
	"gateway/internal/repository"
	"gateway/internal/repository/database/mongo"
	logRepository "gateway/internal/repository/log"
)

func (s *serviceProvider) Mongo() *mongo.Client {
	if s.mongoClient == nil {
		s.mongoClient = mongo.NewClientMongo(s.logger.NewNameLogger(ComponentNameRepositoryMongo.String()))
	}

	return s.mongoClient
}

func (s *serviceProvider) LogRepository() repository.LogRepository {
	if s.fileLog == nil {
		s.fileLog = logRepository.NewLogRepository("Log", s.Mongo()) // TODO: Вынести в константы или в конфиг
	}

	return s.fileLog
}

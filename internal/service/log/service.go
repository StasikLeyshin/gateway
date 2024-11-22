package log

import (
	"gateway/internal/service"
	serviceCore "gateway/internal/service/service"
)

var _ service.LogService = (*logService)(nil)

type (
	logService struct {
		service *serviceCore.Service
	}
)

func NewLogService(service *serviceCore.Service) *logService {
	return &logService{
		service: service,
	}
}

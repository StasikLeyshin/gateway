package log

import (
	"context"
	"encoding/json"
	modelRep "gateway/internal/repository/log/model"
	"gateway/internal/service/log/model"
)

func (l *logService) Write(p []byte) (n int, err error) {
	var addLog model.AddLogRequest

	err = json.Unmarshal(p, &addLog)
	if err != nil {
		return 0, err
	}

	_, err = l.service.Repository.LogRepository.AddLog(context.Background(), new(modelRep.AddLogRequest).ToRepository(&addLog))
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (l *logService) AddLog(ctx context.Context, request *model.AddLogRequest) (*model.AddLogResponse, error) {

	return nil, nil
}

func (l *logService) GetLog(ctx context.Context, request *model.GetLogRequest) (*model.GetLogResponse, error) {

	return nil, nil
}

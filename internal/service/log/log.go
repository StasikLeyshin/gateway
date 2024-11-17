package log

import (
	"context"
	"encoding/json"
	modelRep "gateway/internal/repository/log/model"
	"gateway/internal/service/log/model"
)

func (f *FileLog) Write(p []byte) (n int, err error) {
	var addLog model.AddLogRequest

	err = json.Unmarshal(p, &addLog)
	if err != nil {
		return 0, err
	}

	_, err = f.fileLog.AddLog(context.Background(), new(modelRep.AddLogRequest).ToRepository(&addLog))
	if err != nil {
		return 0, err
	}

	return 1, nil
}

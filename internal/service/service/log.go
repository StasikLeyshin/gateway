package service

import (
	"context"
	"gateway/internal/repository"
)

type (
	FileLog struct {
		fileLog repository.FileLog
	}
)

func NewFileLog(fileLog repository.FileLog) *FileLog {
	return &FileLog{
		fileLog: fileLog,
	}
}

func (f *FileLog) Write(p []byte) (n int, err error) {
	err = f.fileLog.AddLog(context.Background())
	if err != nil {
		return 0, err
	}
	return 1, nil
}

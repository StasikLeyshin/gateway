package log

import "gateway/internal/repository"

type (
	FileLog struct {
		fileLog repository.LogRepository
	}
)

func NewFileLog(fileLog repository.LogRepository) *FileLog {
	return &FileLog{
		fileLog: fileLog,
	}
}

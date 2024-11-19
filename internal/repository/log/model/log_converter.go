package model

import (
	"gateway/internal/service/log/model"
)

func (convert *Log) ToRepository(value *model.Log) *Log {
	return &Log{
		ID:       value.ID,
		Datetime: value.Datetime,
		Level:    value.Level,
		Logger:   value.Logger,
		Caller:   value.Caller,
		FuncName: value.FuncName,
		Message:  value.Message,
	}
}

func (convert *Log) FromSRepository() *model.Log {
	value := &model.Log{}
	value.ID = convert.ID
	value.Datetime = convert.Datetime
	value.Level = convert.Level
	value.Logger = convert.Logger
	value.Caller = convert.Caller
	value.FuncName = convert.FuncName
	value.Message = convert.Message

	return value
}

func (convert *AddLogRequest) ToRepository(value *model.AddLogRequest) *AddLogRequest {
	return &AddLogRequest{
		Datetime: value.Datetime,
		Level:    value.Level.CapitalString(),
		Logger:   value.Logger,
		Caller:   value.Caller,
		FuncName: value.FuncName,
		Message:  value.Message,
	}
}

func (convert *AddLogResponse) FromSRepository() *model.AddLogResponse {
	value := &model.AddLogResponse{}

	return value
}

package models

import "time"

type (
	Log struct {
		Datetime time.Time `json:"date_time"`
		Level    string    `json:"level"`
		Logger   string    `json:"logger"`
		Caller   string    `json:"caller"`
		FuncName string    `json:"func_name"`
		Message  string    `json:"message"`
	}
)

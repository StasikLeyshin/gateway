package model

import "time"

type (
	Log struct {
		ID       string    `json:"id"`
		Datetime time.Time `json:"date_time"`
		Level    string    `json:"level"`
		Logger   string    `json:"logger"`
		Caller   string    `json:"caller"`
		FuncName string    `json:"func_name"`
		Message  string    `json:"message"`
	}
)

// AddLog
type (
	AddLogRequest struct {
		Datetime time.Time `json:"date_time"`
		Level    string    `json:"level"`
		Logger   string    `json:"logger"`
		Caller   string    `json:"caller"`
		FuncName string    `json:"func_name"`
		Message  string    `json:"message"`
	}

	AddLogResponse struct{}
)

// GetLog
type (
	GetLogRequest struct {
		StartTime time.Time `json:"start_time"`
		FinalTime time.Time `json:"final_time"`
		Level     string    `json:"level"`
		Logger    string    `json:"logger"`
		Caller    string    `json:"caller"`
		FuncName  string    `json:"func_name"`
		Message   string    `json:"message"`
	}

	GetLogResponse struct {
		Log Log
	}
)

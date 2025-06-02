package models

import "time"

type APILogs struct {
	APILog []APILog
}

type APILog struct {
	StatusCode int       `json:"status_code"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
	Path       string    `json:"path"`
	RequestID  string    `json:"request_id"`
}

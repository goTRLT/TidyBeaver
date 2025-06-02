package models

import "time"

type MSVCLogs struct {
	MSVCLog []MSVCLog
}

type MSVCLog struct {
	Timestamp     time.Time `json:"timestamp"`
	Service       string    `json:"service"`
	Level         string    `json:"level"`
	Message       string    `json:"message"`
	CorrelationID string    `json:"correlation_id"`
	RequestID     string    `json:"request_id"`
	Host          string    `json:"host"`
}

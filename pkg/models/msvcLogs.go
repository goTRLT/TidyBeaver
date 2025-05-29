package models

type LogEntry struct {
	Timestamp     string `json:"timestamp"`
	Service       string `json:"service"`
	Level         string `json:"level"`
	Message       string `json:"message"`
	CorrelationID string `json:"correlation_id"`
	RequestID     string `json:"request_id"`
	Host          string `json:"host"`
}

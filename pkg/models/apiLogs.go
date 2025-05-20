package models

type APILogs struct {
	APILog []APILog
}

type APILog struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
	Path       string `json:"path"`
	RequestID  string `json:"request_id"`
}

package types

import "time"

type APIResponse struct {
	StatusCode int       `json:"status_code"`
	Status     string    `json:"status"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
	Path       string    `json:"path"`
	RequestID  string    `json:"request_id"`
}

type APIResponseVariant struct {
	StatusCode int
	Status     string
	Messages   []string
}



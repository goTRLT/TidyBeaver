package models

import (
	"time"
)

type FSLogs struct {
	FSLog []struct {
		Level      string    `json:"Level"`
		Service    string    `json:"Service"`
		Message    string    `json:"Message"`
		Time       time.Time `json:"Time"`
		Index      int       `json:"Index"`
		EntryType  int       `json:"EntryType"`
		Source     string    `json:"Source"`
		InstanceID int64     `json:"InstanceID"`
	}
}

package types

import (
	"time"
)

type SampleLogs struct {
	SampleLog []SampleLog
}

type SampleLog struct {
	Level   string    `json:"Level"`
	Service string    `json:"Service"`
	Message string    `json:"Message"`
	Time    time.Time `json:"Time"`
}

var SampleLevels = []string{"INFO", "DEBUG", "WARN", "ERROR"}
var SampleServices = []string{"auth-service", "payment-service", "user-service", "inventory-service"}
var SampleErrorMessages = []string{
	"User authentication failed",
	"Payment transaction error",
	"Database connection timeout",
}
var SampleInfoMessages = []string{
	"Operation completed successfully",
	"Success!",
}

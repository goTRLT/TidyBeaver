package models

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

var AvailableLevels = []string{"INFO", "DEBUG", "WARN", "ERROR"}
var AvailableServices = []string{"auth-service", "payment-service", "user-service", "inventory-service"}
var AvailableErrorMessages = []string{
	"User authentication failed",
	"Payment transaction error",
	"Database connection timeout",
}
var AvailableInfoMessages = []string{
	"Successfully completed operation",
	"Success!",
}

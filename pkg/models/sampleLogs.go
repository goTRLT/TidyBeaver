package models

import (
	"time"
)

type SampleLogs struct {
	SampleLog []SampleLog
}

type SampleLog struct {
	Level   string
	Service string
	Message string
	Time    time.Time
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

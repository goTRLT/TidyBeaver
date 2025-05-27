package models

import (
	"time"
)

type MockedLogs struct {
	MockedLog []MockedLog
}

type MockedLog struct {
	Level   string    `json:"Level"`
	Service string    `json:"Service"`
	Message string    `json:"Message"`
	Time    time.Time `json:"Time"`
}

var MockedLevels = []string{"INFO", "DEBUG", "WARN", "ERROR"}
var MockedServices = []string{"auth-service", "payment-service", "user-service", "inventory-service"}
var MockedErrorMessages = []string{
	"User authentication failed",
	"Payment transaction error",
	"Database connection timeout",
}
var MockedInfoMessages = []string{
	"Operation completed successfully",
	"Success!",
}

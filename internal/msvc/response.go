package msvc

import (
	"math/rand"
	models "tidybeaver/pkg/models"
	"time"

	"github.com/google/uuid"
)

func randomLogEntry(serviceName string, hostname string) models.LogEntry {
	return models.LogEntry{
		Timestamp:     time.Now().Format(time.RFC3339),
		Service:       serviceName,
		Level:         logLevels[rand.Intn(len(logLevels))],
		Message:       messages[rand.Intn(len(messages))],
		CorrelationID: uuid.New().String(),
		RequestID:     uuid.New().String(),
		Host:          hostname,
	}
}

var logLevels = []string{"INFO", "DEBUG", "WARN", "ERROR"}
var messages = []string{
	"User authenticated successfully",
	"Database connection established",
	"Cache miss for key",
	"Service timeout reached",
	"Unhandled exception occurred",
	"Data validation failed",
	"Request processed successfully",
	"External API call failed",
}

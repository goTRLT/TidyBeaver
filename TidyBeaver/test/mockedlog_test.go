package test

import (
    "testing"
    "tidybeaver/pkg/models"
    "time"
)

func TestMockedLog_ToAggregatedLog(t *testing.T) {
    ml := models.MockedLog{
        Level:   "INFO",
        Service: "test-service",
        Message: "test message",
        Time:    time.Now(),
    }
    agg := ml.ToAggregatedLog()
    if agg.Service != "test-service" {
        t.Errorf("Expected Service 'test-service', got '%s'", agg.Service)
    }
}
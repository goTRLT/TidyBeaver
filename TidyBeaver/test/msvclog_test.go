package test

import (
    "testing"
    "tidybeaver/pkg/models"
    "time"
)

func TestMSVCLog_ToAggregatedLog(t *testing.T) {
    msvc := models.MSVCLog{
        Service: "svc",
        Level:   "DEBUG",
        Message: "msg",
        Timestamp: time.Now(),
    }
    agg := msvc.ToAggregatedLog()
    if agg.Service != "svc" {
        t.Errorf("Expected Service 'svc', got '%s'", agg.Service)
    }
}
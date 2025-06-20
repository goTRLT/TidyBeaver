package test

import (
    "testing"
    "tidybeaver/pkg/models"
    "time"
)

func TestFSLog_ToAggregatedLog(t *testing.T) {
    fs := models.FSLog{
        Category: "test",
        Message:  "hello",
        TimeGenerated: time.Now(),
    }
    agg := fs.ToAggregatedLog()
    if agg.Category != "test" {
        t.Errorf("Expected Category 'test', got '%s'", agg.Category)
    }
    if agg.Message != "hello" {
        t.Errorf("Expected Message 'hello', got '%s'", agg.Message)
    }
}
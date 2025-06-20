package test

import (
    "testing"
    "tidybeaver/pkg/models"
    "tidybeaver/utils"
    "time"
)

func TestTransformSlice_FSLog(t *testing.T) {
    fsLogs := []models.FSLog{
        {Category: "cat", Message: "msg", TimeGenerated: time.Now()},
    }
    result := utils.TransformSlice(fsLogs)
    if len(result) != 1 {
        t.Errorf("Expected 1 result, got %d", len(result))
    }
    if result[0].Category != "cat" {
        t.Errorf("Expected Category 'cat', got '%s'", result[0].Category)
    }
}
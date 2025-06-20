package test

import (
    "testing"
    "tidybeaver/pkg/models"
)

func TestDBLog_ToAggregatedLog(t *testing.T) {
    db := models.DBLog{
        Level: "INFO",
        Table_name: "users",
    }
    agg := db.ToAggregatedLog()
    if agg.TableName != "users" {
        t.Errorf("Expected TableName 'users', got '%s'", agg.TableName)
    }
}
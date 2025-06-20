package test

import (
    "testing"
    "tidybeaver/internal/config"
)

func TestConfigInit(t *testing.T) {
    cfg := config.Init()
    if cfg.App.LogAmount == "" {
        t.Error("Expected LogAmount to be set")
    }
}
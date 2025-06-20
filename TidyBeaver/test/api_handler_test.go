package test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "tidybeaver/internal/api"
)

func TestResponseHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/api/random-response?amount=2", nil)
    w := httptest.NewRecorder()
    api.ResponseHandler(w, req)
    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }
}
package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	models "tidybeaver/pkg/models"
)

const (
	MaxAmount     = 100
	DefaultAmount = 1
)

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	var result []models.APIResponse
	countStr := r.URL.Query().Get("amount")
	amount := DefaultAmount
	if countStr != "" {
		if parsed, err := strconv.Atoi(countStr); err == nil && parsed > 0 && parsed <= MaxAmount {
			amount = parsed
		}
	}

	for range amount {
		result = append(result, CreateRandomResponse(r.URL.Path))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

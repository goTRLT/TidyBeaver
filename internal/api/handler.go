package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	models "tidybeaver/pkg/models"
)

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	countStr := r.URL.Query().Get("count")
	count := 1
	if countStr != "" {
		if parsed, err := strconv.Atoi(countStr); err == nil && parsed > 0 && parsed <= 100 {
			count = parsed
		}
	}

	var result []models.APIResponse
	for i := 0; i < count; i++ {
		result = append(result, CreateRandomResponse(r.URL.Path))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

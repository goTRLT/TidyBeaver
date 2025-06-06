package msvc

import (
	"encoding/json"
	"net/http"
	"strconv"
	models "tidybeaver/pkg/models"
)

func MsvcLogHandler(serviceName, hostname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result []models.MSVCLog
		countStr := r.URL.Query().Get("amount")
		amount := 1
		if countStr != "" {
			if parsed, err := strconv.Atoi(countStr); err == nil && parsed > 0 && parsed <= 100 {
				amount = parsed
			}
		}

		for i := 0; i < amount; i++ {
			result = append(result, CreateRandomResponse(serviceName, hostname))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

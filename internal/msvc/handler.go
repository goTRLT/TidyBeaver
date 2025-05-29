package msvc

import (
	"encoding/json"
	"net/http"
)

func MsvcLogHandler(serviceName, hostname string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entry := CreateRandomResponse(serviceName, hostname)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entry)
	}
}

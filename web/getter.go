package web

import (
	"encoding/json"
	"net/http"
	"os"
	"tidybeaver/pkg/models"
)

// FetchLogs retrieves logs from various sources and returns them as JSON.
func FetchLogs(w http.ResponseWriter) (logs models.AggregatedLogs) {
	path := os.Getenv("LOGS_FOLDER_PATH")
	fileName := os.Getenv("LOGS_FILE_NAME")

	logFilePath := path + fileName
	file, err := os.Open(logFilePath)
	if err != nil {
		http.Error(w, "Could not open log file: "+err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&logs); err != nil {
		http.Error(w, "Error decoding log file: "+err.Error(), http.StatusInternalServerError)
	}
	return logs
}

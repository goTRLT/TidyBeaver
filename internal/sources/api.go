package sources

import (
	"encoding/json"
	"log"
	"net/http"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
)

func FetchAPILogs() (APILogs models.APILogs, err error) {

	var APILogEntry []models.APILog
	var responses []models.APILog
	//TODO Add Timeout
	resp, err := http.Get(config.ConfigValues.API.BaseURL + config.ConfigValues.App.LogAmount)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		log.Fatal(err)
	}

	APILogs.APILog = append(APILogEntry, responses...)

	return APILogs, err
}

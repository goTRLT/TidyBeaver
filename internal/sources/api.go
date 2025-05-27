package sources

import (
	"encoding/json"
	"log"
	"net/http"
	"tidybeaver/internal/config"
	"tidybeaver/pkg/models"
)

func FetchAPILogs() (APILogs models.APILogs, err error) {

	var APILogEntry []models.APILog
	//TODO Add Timeout
	resp, err := http.Get(config.ConfigValues.API.BaseURL + config.ConfigValues.App.LogAmount)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// fmt.Println("Resp ", resp)

	var responses []models.APILog

	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("ActResp ", responses)

	APILogs.APILog = append(APILogEntry, responses...)

	return APILogs, err
}

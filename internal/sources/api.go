package sources

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
	"time"
)

func GetAPILogs() (APILogs models.APILogs, err error) {
	var APILogEntry []models.APILog
	var responses []models.APILog

	timeoutSecondsStr := config.EnvVar["API_TIMEOUTSECONDS"]
	timeoutSeconds, err := strconv.Atoi(timeoutSecondsStr)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: -time.Duration(timeoutSeconds) * time.Second,
	}

	resp, err := client.Get(config.EnvVar["API_BASEURL"] + config.ConfigValues.App.LogAmount)

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

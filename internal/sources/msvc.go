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

func GetMSVCLogs() (MSVCLogs models.MSVCLogs, err error) {
	var MSVCLogEntry []models.MSVCLog
	var responses []models.MSVCLog

	timeoutSecondsStr := config.EnvVar["MSVC_TIMEOUTSECONDS"]
	timeoutSeconds, err := strconv.Atoi(timeoutSecondsStr)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: -time.Duration(timeoutSeconds) * time.Second,
	}

	resp, err := client.Get(config.EnvVar["MSVC_BASEURL"] + config.ConfigValues.App.LogAmount)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		log.Fatal(err)
	}

	MSVCLogs.MSVCLog = append(MSVCLogEntry, responses...)

	return MSVCLogs, err
}

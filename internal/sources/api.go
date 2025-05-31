package sources

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
	"time"
)

func GetAPILogs() (APIL models.APILogs, err error) {
	var APILtemp []models.APILog
	var responses []models.APILog

	timeoutSecondsStr := os.Getenv("API_TIMEOUTSECONDS")
	timeoutSeconds, err := strconv.Atoi(timeoutSecondsStr)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: -time.Duration(timeoutSeconds) * time.Second,
	}

	resp, err := client.Get(os.Getenv("API_BASEURL") + config.CFG.App.LogAmount)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		log.Fatal(err)
	}

	APIL.APIL = append(APILtemp, responses...)

	return APIL, err
}

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

func GetMSVCLogs() (msvcls models.MSVCLogs, err error) {
	var MSVCLtemp []models.MSVCLog
	var responses []models.MSVCLog

	timeoutSecondsStr := os.Getenv("MSVC_TIMEOUTSECONDS")
	timeoutSeconds, err := strconv.Atoi(timeoutSecondsStr)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: -time.Duration(timeoutSeconds) * time.Second,
	}

	resp, err := client.Get(os.Getenv("MSVC_BASEURL") + os.Getenv("MSVC_PORT") + os.Getenv("MSVC_REQUESTURL") + config.CFG.App.LogAmount)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		log.Fatal(err)
	}

	msvcls.MSVCLog = append(MSVCLtemp, responses...)

	return msvcls, err
}

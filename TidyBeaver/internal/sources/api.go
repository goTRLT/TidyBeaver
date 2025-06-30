package sources

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
	"time"
)

func GetAPILogs() (apiLogs *models.APILogs, err error) {
	apiLogs = &models.APILogs{}
	var responses []models.APILog

	timeoutSeconds, err := strconv.Atoi(os.Getenv("API_TIMEOUTSECONDS"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, os.Getenv("API_BASEURL")+os.Getenv("API_PORT")+os.Getenv("API_REQUESTURL")+config.CFG.App.LogAmount, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	apiLogs.APILog = append(apiLogs.APILog, responses...)

	return apiLogs, err
}

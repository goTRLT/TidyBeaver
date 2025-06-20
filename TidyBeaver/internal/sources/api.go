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

func GetAPILogs() (APILS models.APILogs, err error) {
	var APILtemp []models.APILog
	var responses []models.APILog

	timeoutSeconds, err := strconv.Atoi(os.Getenv("API_TIMEOUTSECONDS"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, os.Getenv("API_BASEURL")+config.CFG.App.LogAmount, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	APILS.APILog = append(APILtemp, responses...)

	return APILS, err
}

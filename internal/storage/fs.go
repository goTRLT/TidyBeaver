package storage

import (
	"encoding/json"
	"fmt"
	"os"
	models "tidybeaver/pkg/models"
	"time"
)

func JSONSaveLogs(Logs *models.AggregatedLogs) {
	encodedLogs, err := json.Marshal(Logs)
	if err != nil {
		return
	} else {
		os.WriteFile((`.\logs\TidyBeaverAdaptedLogs.json`), encodedLogs, 0644)
		fmt.Println("Logs saved as Json")
		time.Sleep(1000000000)
	}
}

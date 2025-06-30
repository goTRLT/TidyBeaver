package sources

import (
	"encoding/json"
	"log"
	"os"
	models "tidybeaver/pkg/models"
)

func GetFSLogs() (fsLogs *models.FSLogs, err error) {
	fsLogs = &models.FSLogs{}
	var tempLogs models.FSLogs
	logFile, err := os.Open(`.\logs\` + "InputLogs.json")

	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()

	decodedJson := json.NewDecoder(logFile)
	decodedJson.Decode(&tempLogs)

	if err != nil {
		log.Printf("Error decoding file %s: %v", "InputLogs.json", err)
	}

	fsLogs.FSLog = append(fsLogs.FSLog, tempLogs.FSLog...)
	return fsLogs, err
}

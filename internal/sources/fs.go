package sources

import (
	"encoding/json"
	"log"
	"os"
	models "tidybeaver/pkg/models"
)

func FetchFSLogs() (FSLogs models.FSLogs, err error) {
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

	FSLogs.FSLog = append(FSLogs.FSLog, tempLogs.FSLog...)
	return FSLogs, err
}

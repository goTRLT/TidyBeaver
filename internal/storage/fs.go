package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	models "tidybeaver/pkg/models"
	"time"
)

func JSONSaveLogs(Logs *models.AggregatedLogs) {
	path := os.Getenv("LOGS_FOLDER_PATH")
	fileName := os.Getenv("LOGS_FILE_NAME")

	err := os.Mkdir(path, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	file, err := os.Create(path + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, logEntry := range Logs.AggregatedLog {
		line, err := json.Marshal(logEntry)
		if err != nil {
			log.Println("Error encoding log entry:", err)
			continue
		}
		_, err = file.Write(append(line, '\n'))
		if err != nil {
			log.Println("Error writing to file:", err)
			continue
		}
	}

	fmt.Println("Logs saved as NDJSON: " + path + fileName)
	time.Sleep(500 * time.Millisecond)
}

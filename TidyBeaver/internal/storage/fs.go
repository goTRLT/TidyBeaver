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
	// path := os.Getenv("OUTPUT_FOLDER_PATH")
	// fileName := os.Getenv("OUTPUT_FILE_NAME")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current working directory:", dir)

	// err = os.Mkdir(path, 0750)
	err = os.Mkdir(`/logs/output/`, 0750)
	if err != nil && !os.IsExist(err) {
		log.Println(err)
		return
	}

	// file, err := os.Create(path + fileName)
	file, err := os.OpenFile("/logs/output/OutputLogs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	for _, logEntry := range Logs.AggregatedLog {
		logEntry.Timestamp = time.Now().UTC()
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

	fmt.Println("Logs saved as NDJSON: " + dir + "/logs/output/OutputLogs.json")
	time.Sleep(500 * time.Millisecond)
}

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
	encodedLogs, err := json.Marshal(Logs)
	if err != nil {
		return
	} else {
		path := os.Getenv("LOGS_FOLDER_PATH")
		fileName := os.Getenv("LOGS_FILE_NAME")

		err := os.Mkdir(path, 0750)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}

		err = os.WriteFile((path + fileName), encodedLogs, 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("jFilepath: ", path)
		fmt.Println("jfileName: ", fileName)

		fmt.Println(`Logs saved as Json: ` + path + fileName)
		time.Sleep(500 * time.Millisecond)
	}
}

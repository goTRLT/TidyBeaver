package sources

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	models "tidybeaver/pkg/models"
)

var FileDetailedLogs models.FSLogs

func FetchFSLogs() (model models.FSLogs, err error) {
	files, err := os.ReadDir(`.\Logs`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Files: ", files)

	for _, file := range files {
		logFile, err := os.Open(`.\Logs\` + file.Name())

		if err != nil {
			log.Fatal(err)
		}

		defer logFile.Close()
		decodedJson := json.NewDecoder(logFile)
		decodedJson.Decode(&FileDetailedLogs)
		// indentedDetailedLog, err := json.MarshalIndent(FileDetailedLogs.StandardLog, "", "  ")

		if err != nil {
			log.Fatal("Error marshalling the Indented Detailed Log:", err)
			return FileDetailedLogs, err
		}

		// fmt.Println("logFile: ", logFile.Name())
		// fmt.Println("Detailed Log: ", FileDetailedLogs)
	}
	return FileDetailedLogs, err
}

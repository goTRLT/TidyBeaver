package sources

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	models "tidybeaver/pkg/models"
)

var FileDetailedLogs models.TransformedLogs
var FileDetailedLog models.TransformedLog

func GetLogsFromFS() models.TransformedLogs {
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
		indentedDetailedLog, err := json.MarshalIndent(FileDetailedLog, "", "  ")

		if err != nil {
			fmt.Println("Error marshalling the Indented Detailed Log:", err)
			return FileDetailedLogs
		}
		
		fmt.Println("Detailed Log: ", string(indentedDetailedLog))
	}
	return FileDetailedLogs
}

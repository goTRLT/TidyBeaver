package sources

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	models "tidybeaver/pkg/models"
)

var FileDetailedLogs models.DetailedLogs
var FileDetailedLog models.DetailedLog

func GetLogsFromFileSystem() {
	files, err := os.ReadDir(`.\Logs`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files: ", files)

	for _, file := range files {
		test, err := os.Open(`.\Logs\` + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer test.Close()
		test2 := json.NewDecoder(test)
		test2.Decode(&FileDetailedLog)
		indentedDetailedLog, err := json.MarshalIndent(FileDetailedLog, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling the Indented Detailed Log:", err)
			return
		}
		fmt.Println("Detailed Log: ", string(indentedDetailedLog))
	}
}

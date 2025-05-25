package sources

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	typess "tidybeaver/pkg/types"
)

func FetchFSLogs() (types typess.FSLogs, err error) {
	var FSLogs typess.FSLogs
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
		var tempLogs typess.FSLogs

		decodedJson := json.NewDecoder(logFile)
		decodedJson.Decode(&tempLogs)

		if err != nil {
			log.Printf("Error decoding file %s: %v", file.Name(), err)
			continue
		}

		FSLogs.FSLog = append(FSLogs.FSLog, tempLogs.FSLog...)

		// indentedFSLog, err := json.MarshalIndent(tempLogs, "", "  ")

		// if err != nil {
		// 	log.Fatal("Error marshalling the Indented Detailed Log:", err)
		// 	return FSLogs, err
		// }

		// fmt.Println("logFile: ", logFile.Name())
		// fmt.Println("indentedFSLog: ", string(indentedFSLog))
		// fmt.Println("FSLogs: ", FSLogs)
	}
	return FSLogs, err
}

package storage

import (
	"encoding/json"
	"fmt"
	"os"
	models "tidybeaver/pkg/models"
)

func WriteSampleLogsToFile(sampleLogs models.SampleLogs) {
	receivedLogsByte := fmt.Sprint(sampleLogs)
	os.WriteFile((`.\Logs\TidyBeaverSampleLogs.txt`), []byte(receivedLogsByte), 0644)
	fmt.Println("Sample Logs saved as Text")

	encodedSampleLogs, err := json.Marshal(sampleLogs)
	if err != nil {
		return
	} else {
		os.WriteFile((`.\Logs\TidyBeaverSampleLogs.json`), encodedSampleLogs, 0644)
		fmt.Println("Sample Logs saved as Json")
	}
}

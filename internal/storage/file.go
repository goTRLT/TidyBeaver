package storage

import (
	"encoding/json"
	"fmt"
	"os"
	models "tidybeaver/pkg/models"
)

func WriteLogsToFile(receivedLogs models.SampleLogs) {
	fmt.Println("Logs received to be writen to file (text):")
	fmt.Println(receivedLogs)
	receivedLogsByte := fmt.Sprint(receivedLogs)
	os.WriteFile(("TidyBeaverLogs.txt"), []byte(receivedLogsByte), 0644)

	encodedLogs, err := json.Marshal(receivedLogs)
	if err != nil {
		return
	}

	fmt.Println("Logs received to be writen to file (json):")
	fmt.Println(string(encodedLogs))
	os.WriteFile(("TidyBeaverLogs.json"), encodedLogs, 0644)
}

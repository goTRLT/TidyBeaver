package storage

import (
	"fmt"
	"os"
	models "tidybeaver/pkg/models"
)

func WriteLogsToFile(receivedLogs models.SampleLogs) {
	fmt.Println("Logs received to be writen to file:")
	fmt.Println(receivedLogs)
	receivedLogsByte := fmt.Sprint(receivedLogs)
	os.WriteFile(("TidyBeaverLogs.txt"), []byte(receivedLogsByte), 0644)
}

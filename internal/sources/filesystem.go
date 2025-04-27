package sources

import (
	"fmt"
	models "tidybeaver/pkg/models"
)

func GetLogsFromFS() {
	generatedSampleLogs := models.GenerateLogs()
	fmt.Println(generatedSampleLogs)
}

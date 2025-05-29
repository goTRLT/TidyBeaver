package sources

import (
	"fmt"
	models "tidybeaver/pkg/models"
)

func GetMSVCLogs() (m models.MSVCLogs, err error) {
	fmt.Println("msvcLogs")
	return m, err
}

package sources

import (
	"encoding/json"
	"fmt"
	"os/exec"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
)

func GetLogsFromOS() (windowsEventLogs models.WindowsEventLogs, err error) {
	cmd := exec.Command("powershell", "-Command", "Get-EventLog -LogName Application -Newest "+config.ConfigValues.App.LogAmount+" | ConvertTo-Json -Depth 2")
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error running PowerShell command: %v\n", err)
		return windowsEventLogs, err
	}

	err = json.Unmarshal(out, &windowsEventLogs.WindowsEventLogs)

	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return windowsEventLogs, err
	}

	fmt.Print(windowsEventLogs)
	return windowsEventLogs, err
}

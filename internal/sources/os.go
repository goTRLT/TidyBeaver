package sources

import (
	"encoding/json"
	"fmt"
	"os/exec"
	config "tidybeaver/internal/config"
)

type WindowsEventLogs struct {
	WindowsEventLogs []WindowsEventLog
}

type WindowsEventLog struct {
	MachineName        string   `json:"MachineName"`
	Data               []string `json:"Data"`
	Index              int      `json:"Index"`
	Category           string   `json:"Category"`
	CategoryNumber     int      `json:"CategoryNumber"`
	EventID            int      `json:"EventID"`
	EntryType          int      `json:"EntryType"`
	Message            string   `json:"Message"`
	Source             string   `json:"Source"`
	ReplacementStrings []string `json:"ReplacementStrings"`
	InstanceID         int      `json:"InstanceID"`
	TimeGenerated      string   `json:"TimeGenerated"`
	TimeWritten        string   `json:"TimeWritten"`
	UserName           string   `json:"UserName"`
	SplitLines         string   `json:"SplitLines"`
	Container          string   `json:"Container"`
}

var windowsEventLogs WindowsEventLogs

func GetLogsFromOS() WindowsEventLogs {
	cmd := exec.Command("powershell", "-Command", "Get-EventLog -LogName Application -Newest "+config.ConfigValues.App.LogAmount+" | ConvertTo-Json -Depth 2")
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error running PowerShell command: %v\n", err)
		return windowsEventLogs
	}

	err = json.Unmarshal(out, &windowsEventLogs.WindowsEventLogs)
	
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return windowsEventLogs
	}

	fmt.Print(windowsEventLogs)
	return windowsEventLogs

}

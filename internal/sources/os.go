package sources

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	config "tidybeaver/internal/config"
)

type WindowsEventLogs struct {
	WindowsEventLogs []WindowsEventLogs
}

type WindowsEventLog struct {
	Index      int    `json:"Index"`
	Time       string `json:"Time"`
	EntryType  string `json:"EntryType"`
	Source     string `json:"Source"`
	InstanceID int    `json:"InstanceID"`
	Message    string `json:"Message"`
}

func GetLogsFromOS() {
	cmd := exec.Command("powershell", "-Command", "Get-EventLog -LogName Application -Newest "+config.ConfigValues.App.LogAmount+" | Format-Table -AutoSize")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running PowerShell command: %v\n", err)
		return
	}
	fmt.Println("Application Event Logs:")
	fmt.Println(out.String())
	os.WriteFile("TBEventLogs.txt", out.Bytes(), 0644)

	stringz := out.String()
	strLines := SplitLines(stringz)
	fmt.Println("strLines:")
	fmt.Println(strLines)

	// receivedLogs := ""
	// encodedLogs, err := json.Marshal(receivedLogs)
	// if err != nil {
	// 	return
	// }

	// fmt.Println("Logs received to be writen to file (json):")
	// fmt.Println(string(encodedLogs))
	// os.WriteFile(("TBEventLogs.json"), encodedLogs, 0644)
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

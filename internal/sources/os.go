package sources

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	config "tidybeaver/internal/config"
)

type WindowsEventLogs struct {
	WindowsEventLogs []WindowsEventLogs
}

type WindowsEventLog struct {
	Index      string `json:"Index"`
	Time       string `json:"Time"`
	EntryType  string `json:"EntryType"`
	Source     string `json:"Source"`
	InstanceID string `json:"InstanceID"`
	Message    string `json:"Message"`
}

var windowsEventLog WindowsEventLog

func GetLogsFromOS() {
	cmd := exec.Command("powershell", "-Command", "Get-EventLog -LogName Application -Newest "+config.ConfigValues.App.LogAmount+" | Format-Table -AutoSize")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running PowerShell command: %v\n", err)
		return
	}
	fmt.Println("Application Event Logs: ", out.String())
	os.WriteFile("TBEventLogs.txt", out.Bytes(), 0644)

	stringz := out.String()
	strLines := SplitLines(stringz)

	for i := 0; i < len(strLines); i++ {

		regex := regexp.MustCompile(`^.{1}(.{6})`)
		test := regex.FindString(strLines[i])
		fmt.Println("test: ", test)

		fmt.Println("strLine: ", strings.Split(strLines[i], strLines[i]))
		// windowsEventLog.Index =	string(regex.Split(strLines[i], -1))
		// windowsEventLog.Time = regexp.MustCompile("^.{7}(.{19})")
		// windowsEventLog.EntryType = regexp.MustCompile("^.{10}(.{5})")
		// windowsEventLog.Source = regexp.MustCompile("^.{10}(.{5})")
		// windowsEventLog.Message = regexp.MustCompile("^.{10}(.{5})")
	}

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

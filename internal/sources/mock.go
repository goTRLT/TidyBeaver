package sources

import (
	"fmt"
	"math/rand"
	"strconv"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
	"time"
)

var SampleLogsEntry models.SampleLogs
var SampleLogEntry models.SampleLog

func GetLogsFromMock() models.SampleLogs {
	generatedMockLogs := GenerateMockLogs()
	fmt.Println(generatedMockLogs)
	return generatedMockLogs
}

func randomChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}

func randomTime() time.Time {
	now := time.Now()
	maxOffset := int64(30 * 24 * 60 * 60)
	randomizedOffset := rand.Int63n(maxOffset)
	return now.Add(-time.Duration(randomizedOffset) * time.Second)
}

func selectMessage(selectedLevel string) string {
	var selectMessage string
	if selectedLevel == "WARN" || selectedLevel == "ERROR" {
		selectMessage = randomChoice(models.AvailableErrorMessages)
	} else if selectedLevel == "INFO" || selectedLevel == "DEBUG" {
		selectMessage = randomChoice(models.AvailableInfoMessages)
	}
	return selectMessage
}

func GenerateMockLogs() models.SampleLogs {
	rand.Seed(time.Now().UnixNano())
	logAmount, _ := strconv.ParseInt(config.ConfigValues.App.LogAmount, 0, 0)
	logAmountCast := int(logAmount)

	for i := 0; i < logAmountCast; i++ {
		SampleLogEntry = GenerateLog()
		fmt.Println(SampleLogEntry)
		SampleLogsEntry.SampleLog = append(SampleLogsEntry.SampleLog, SampleLogEntry)
	}
	return SampleLogsEntry
}

func GenerateLog() (sampleLog models.SampleLog) {
	sampleLog.Level = randomChoice(models.AvailableLevels)
	sampleLog.Service = randomChoice(models.AvailableServices)
	sampleLog.Message = selectMessage(sampleLog.Level)
	sampleLog.Time = randomTime()
	return sampleLog
}

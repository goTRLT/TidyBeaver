package sources

import (
	"fmt"
	"math/rand"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
	"time"
)

var SampleLogsEntry models.SampleLogs
var SampleLogEntry models.SampleLog

func GetLogsFromMock() models.SampleLogs {
	generatedSampleLogs := GenerateSampleLogs()
	fmt.Println(generatedSampleLogs)
	return generatedSampleLogs
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

func selectSampleMessage(sampleLevel string) string {
	var sampleMessage string
	if sampleLevel == "WARN" || sampleLevel == "ERROR" {
		sampleMessage = randomChoice(models.SampleErrorMessages)
	} else if sampleLevel == "INFO" || sampleLevel == "DEBUG" {
		sampleMessage = randomChoice(models.SampleInfoMessages)
	}
	return sampleMessage
}

func GenerateSampleLogs() models.SampleLogs {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < config.LogAmountSet; i++ {
		SampleLogEntry = GenerateLog()
		fmt.Println(SampleLogEntry)
		SampleLogsEntry.SampleLog = append(SampleLogsEntry.SampleLog, SampleLogEntry)
	}
	return SampleLogsEntry
}

func GenerateLog() (sampleLog models.SampleLog) {
	sampleLog.Level = randomChoice(models.SampleLevels)
	sampleLog.Service = randomChoice(models.SampleServices)
	sampleLog.Message = selectSampleMessage(sampleLog.Level)
	sampleLog.Time = randomTime()
	return sampleLog
}

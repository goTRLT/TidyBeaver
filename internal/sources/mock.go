package sources

import (
	"errors"
	"log"
	"math/rand"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
	"time"
)

var SampleLogsEntry models.SampleLogs
var SampleLogEntry models.SampleLog

func CreateSampleLogs() (model models.SampleLogs, err error) {
	generatedSampleLogs, err := GenerateSampleLogs()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(generatedSampleLogs)
	return generatedSampleLogs, err
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

func GenerateSampleLogs() (model models.SampleLogs, err error) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < config.LogAmountSet; i++ {
		SampleLogEntry, err = GenerateLog()
		// fmt.Println(SampleLogEntry)
		SampleLogsEntry.SampleLog = append(SampleLogsEntry.SampleLog, SampleLogEntry)
	}

	if err != nil {
		log.Fatal(err)
	}

	return SampleLogsEntry, err
}

func GenerateLog() (sampleLog models.SampleLog, err error) {
	sampleLog.Level = randomChoice(models.SampleLevels)
	sampleLog.Service = randomChoice(models.SampleServices)
	sampleLog.Message = selectSampleMessage(sampleLog.Level)
	sampleLog.Time = randomTime()

	if sampleLog.Level == "" || sampleLog.Service == "" || sampleLog.Message == "" {
		err = errors.New("error while setting level, service or message for the sample log")
		log.Fatal(err)
		return sampleLog, err
	}

	return sampleLog, err
}

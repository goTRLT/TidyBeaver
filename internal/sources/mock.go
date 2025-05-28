package sources

import (
	"errors"
	"log"
	"math/rand"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
	"time"
)

var MockedLogsEntry models.MockedLogs
var MockedLogEntry models.MockedLog

func CreateMockedLogs() (model models.MockedLogs, err error) {
	generatedMockedLogs, err := GenerateMockedLogs()

	if err != nil {
		log.Fatal(err)
	}

	return generatedMockedLogs, err
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

func selectMockedMessage(MockedLevel string) string {
	var MockedMessage string
	if MockedLevel == "WARN" || MockedLevel == "ERROR" {
		MockedMessage = randomChoice(models.MockedErrorMessages)
	} else if MockedLevel == "INFO" || MockedLevel == "DEBUG" {
		MockedMessage = randomChoice(models.MockedInfoMessages)
	}
	return MockedMessage
}

func GenerateMockedLogs() (model models.MockedLogs, err error) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < config.LogAmountSet; i++ {
		MockedLogEntry, err = GenerateLog()
		MockedLogsEntry.MockedLog = append(MockedLogsEntry.MockedLog, MockedLogEntry)
	}

	if err != nil {
		log.Fatal(err)
	}

	return MockedLogsEntry, err
}

func GenerateLog() (MockedLog models.MockedLog, err error) {
	MockedLog.Level = randomChoice(models.MockedLevels)
	MockedLog.Service = randomChoice(models.MockedServices)
	MockedLog.Message = selectMockedMessage(MockedLog.Level)
	MockedLog.Time = randomTime()

	if MockedLog.Level == "" || MockedLog.Service == "" || MockedLog.Message == "" {
		err = errors.New("error while setting level, service or message for the Mocked log")
		log.Fatal(err)
		return MockedLog, err
	}

	return MockedLog, err
}

package sources

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	config "tidybeaver/internal/config"
	models "tidybeaver/pkg/models"
	"time"
)

var mls models.MockedLogs
var ml models.MockedLog

func CreateMockedLogs() (model models.MockedLogs, err error) {
	logAmount, _ := strconv.ParseInt(config.CFG.App.LogAmount, 0, 0)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(logAmount); i++ {
		ml, err = GenerateLog()
		mls.MockedLog = append(mls.MockedLog, ml)
	}

	if err != nil {
		log.Fatal(err)
	}

	return mls, err
}

func GenerateLog() (ml models.MockedLog, err error) {
	ml.Level = randomChoice(models.MockedLevels)
	ml.Service = randomChoice(models.MockedServices)
	ml.Message = selectMockedMessage(ml.Level)
	ml.Time = randomTime()

	if ml.Level == "" || ml.Service == "" || ml.Message == "" {
		err = errors.New("error while setting level, service or message for the Mocked log")
		log.Fatal(err)
		return ml, err
	}

	return ml, err
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

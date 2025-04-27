package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"tidybeaver/internal/config"
	"time"
)

type SampleLogs struct {
	sampleLog []SampleLog
}

type SampleLog struct {
	level   string
	service string
	message string
	time    time.Time
}

var sampleLogs SampleLogs
var sampleLog SampleLog
var levels = []string{"INFO", "DEBUG", "WARN", "ERROR"}
var services = []string{"auth-service", "payment-service", "user-service", "inventory-service"}
var errorMessages = []string{
	"Successfully completed operation",
	"User authentication failed",
	"Payment transaction error",
	"Database connection timeout",
}
var infoMessages = []string{
	"Successfully completed operation",
	"User authentication failed",
	"Payment transaction error",
	"Database connection timeout",
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
		selectMessage = randomChoice(errorMessages)
	} else if selectedLevel == "INFO" || selectedLevel == "DEBUG" {
		selectMessage = randomChoice(infoMessages)
	}
	return selectMessage
}

func GenerateLogs() SampleLogs {
	rand.Seed(time.Now().UnixNano())
	logAmount, _ := strconv.ParseInt(config.ConfigValues.App.LogAmount, 0, 0)
	logAmountCast := int(logAmount)

	for i := 0; i < logAmountCast; i++ {
		sampleLog = GenerateLog()
		fmt.Println(sampleLog)
		sampleLogs.sampleLog = append(sampleLogs.sampleLog, sampleLog)
	}
	return sampleLogs
}

func GenerateLog() (sampleLog SampleLog) {
	sampleLog.level = randomChoice(levels)
	sampleLog.service = randomChoice(services)
	sampleLog.message = selectMessage(sampleLog.level)
	sampleLog.time = randomTime()
	return sampleLog
}

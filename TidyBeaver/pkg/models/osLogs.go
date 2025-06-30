package models

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type OSLogs struct {
	OSLog []OSLog
}

type OSLog struct {
	Category           string   `json:"Category"`
	CategoryNumber     int      `json:"CategoryNumber"`
	Container          string   `json:"Container"`
	Data               []int    `json:"Data"`
	EntryType          int      `json:"EntryType"`
	EventID            int      `json:"EventID"`
	Index              int      `json:"Index"`
	InstanceID         int64    `json:"InstanceID"`
	MachineName        string   `json:"MachineName"`
	Message            string   `json:"Message"`
	ReplacementStrings []string `json:"ReplacementStrings"`
	Source             string   `json:"Source"`
	SplitLines         string   `json:"SplitLines"`
	TimeGenerated      string   `json:"TimeGenerated"`
	TimeWritten        string   `json:"TimeWritten"`
	UserName           string   `json:"UserName"`
}

func (v OSLog) ToAggregatedLog() AggregatedLog {
	v.TimeWritten = strings.TrimPrefix(v.TimeWritten, "/Date(")
	v.TimeWritten = strings.TrimSuffix(v.TimeWritten, ")/")

	parsedTime, err := strconv.ParseInt(v.TimeWritten, 10, 64)
	seconds := parsedTime / 1000
	nanoseconds := (parsedTime % 1000) * 1000000
	unixTime := time.Unix(seconds, nanoseconds)
	unixTime = unixTime.Round(time.Millisecond)

	if err != nil {
		log.Println(err.Error())
	}

	return AggregatedLog{
		Category:           v.Category,
		CategoryNumber:     v.CategoryNumber,
		Container:          v.Container,
		Data:               v.Data,
		EntryType:          v.EntryType,
		EventID:            v.EventID,
		Index:              v.Index,
		InstanceID:         v.InstanceID,
		MachineName:        v.MachineName,
		Message:            v.Message,
		ReplacementStrings: v.ReplacementStrings,
		Source:             "Operational System: " + v.Source,
		SplitLines:         v.SplitLines,
		TimeGenerated:      unixTime,
		TimeWritten:        time.Now(),
		UserName:           v.UserName,
	}
}

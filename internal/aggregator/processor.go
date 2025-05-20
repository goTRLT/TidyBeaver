package aggregator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"tidybeaver/pkg/models"
	"time"
)

func TransformSampleLogs(SampleLogs *models.SampleLogs) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range SampleLogs.SampleLog {
		transformedLog := models.AggregatedLog{
			Level:         val.Level,
			Message:       val.Message,
			Service:       val.Service,
			Source:        "Sample Log",
			TimeGenerated: val.Time,
			TimeWritten:   time.Now(),
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}
	// fmt.Println(transformedLogs)
	return transformedLogs
}

func TransformFSLogs(FSLogs *models.FSLogs) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range FSLogs.FSLog {
		transformedLog := models.AggregatedLog{
			EntryType:     val.EntryType,
			Index:         val.Index,
			InstanceID:    val.InstanceID,
			Level:         val.Level,
			Message:       "FileSystem: " + val.Message,
			Service:       val.Service,
			Source:        val.Source,
			TimeGenerated: val.Time,
			TimeWritten:   time.Now(),
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}
	// fmt.Println(transformedLogs)
	return transformedLogs
}

func TransformDBLogs(DBLogs *models.DBLogs) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range DBLogs.DBLog {
		transformedLog := models.AggregatedLog{
			Column:        val.Column,
			CorrelationID: val.Constraint,
			Datatype:      val.Datatype,
			Detail:        val.Detail,
			Errcode:       val.Errcode,
			Level:         val.Level,
			Schema:        val.Schema,
			Source:        "Database",
			TableName:     val.Table_name,
			TimeGenerated: time.Now(),
			TimeWritten:   time.Now(),
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}
	// fmt.Println(transformedLogs)
	return transformedLogs
}

func TransformOSLogs(OSLogs *models.OSLogs) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range OSLogs.OS {
		val.TimeWritten = strings.TrimPrefix(val.TimeWritten, "/Date(")
		val.TimeWritten = strings.TrimSuffix(val.TimeWritten, ")/")
		fmt.Println(val.TimeWritten)

		parsedTime, err := strconv.ParseInt(val.TimeWritten, 10, 64)
		if err != nil {
			log.Fatal("failed to parse milliseconds: %w", err)
		}

		seconds := parsedTime / 1000
		nanoseconds := (parsedTime % 1000) * 1000000
		unixTime := time.Unix(seconds, nanoseconds)
		unixTime = unixTime.Round(time.Millisecond)
		fmt.Println("unix: ", unixTime)

		if err != nil {
			log.Fatal(err)
		}

		transformedLog := models.AggregatedLog{
			Category:           val.Category,
			CategoryNumber:     val.CategoryNumber,
			Container:          val.Container,
			Data:               val.Data,
			EntryType:          val.EntryType,
			EventID:            val.EventID,
			Index:              val.Index,
			InstanceID:         val.InstanceID,
			MachineName:        val.MachineName,
			Message:            val.Message,
			ReplacementStrings: val.ReplacementStrings,
			Source:             "Operational System: " + val.Source,
			SplitLines:         val.SplitLines,
			TimeGenerated:      unixTime,
			TimeWritten:        time.Now(),
			UserName:           val.UserName,
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}
	// fmt.Println(transformedLogs)
	return transformedLogs
}

// func TransformAPILogs(APILogs *[]string) models.AggregatedLogs    {}
// func TransformMSVLogs(MSVLogs *[]string) models.AggregatedLogs    {}

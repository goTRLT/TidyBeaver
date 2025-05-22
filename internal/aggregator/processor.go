package aggregator

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
	"time"
)

func TransformLogs() {
	if len(SampleLogs.SampleLog) != 0 {
		TransformedLogs, err := TransformSampleLogs(&SampleLogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if len(OSLogs.OS) != 0 {
		TransformedLogs, err := TransformOSLogs(&OSLogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if len(FSLogs.FSLog) != 0 {
		TransformedLogs, err := TransformFSLogs(&FSLogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if len(DBLogs.DBLog) != 0 {
		TransformedLogs, err := TransformDBLogs(&DBLogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if len(APILogs.APILog) != 0 {
		TransformedLogs, err := TransformAPILogs(&APILogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if Errors != nil {
		TransformedLogs := TransformErrors(Errors)
		AggregateLogs(&TransformedLogs)
	}
	// if len(MSVLogs) != 0 {
	// 	//TODO
	// }
}

func TransformSampleLogs(SampleLogs *models.SampleLogs) (aggregatedLogs []models.AggregatedLog, err error) {
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

	if transformedLogs == nil {
		return transformedLogs, fmt.Errorf("error on Transforming Sample Logs into Standard Logs")
	}
	// fmt.Println(transformedLogs)
	return transformedLogs, err
}

func TransformFSLogs(FSLogs *models.FSLogs) (aggregatedLogs []models.AggregatedLog, err error) {
	var transformedLogs []models.AggregatedLog
	for _, val := range FSLogs.FSLog {
		transformedLog := models.AggregatedLog{
			Category:           val.Category,
			CategoryNumber:     val.CategoryNumber,
			Checksum:           val.Checksum,
			ClientIP:           val.ClientIP,
			Column:             val.Column,
			Component:          val.Component,
			ComputerName:       val.ComputerName,
			Constraint:         val.Constraint,
			Container:          val.Container,
			CorrelationID:      val.CorrelationID,
			Data:               val.Data,
			Datatype:           val.Datatype,
			Detail:             val.Detail,
			Endpoint:           val.Endpoint,
			EntryType:          val.EntryType,
			Environment:        val.Environment,
			Errcode:            val.Errcode,
			EventID:            val.EventID,
			EventType:          val.EventType,
			Path:               val.Path,
			FileSize:           val.FileSize,
			Host:               val.Host,
			HTTPMethod:         val.HTTPMethod,
			Index:              val.Index,
			InstanceID:         val.InstanceID,
			LatencyMs:          val.LatencyMs,
			Level:              val.Level,
			LineNumber:         val.LineNumber,
			LogName:            val.LogName,
			MachineName:        val.MachineName,
			Message:            "FileSystem: " + val.Message,
			RequestBody:        val.RequestBody,
			ReplacementStrings: val.ReplacementStrings,
			ResponseBody:       val.ResponseBody,
			RowsAffected:       val.RowsAffected,
			Schema:             val.Schema,
			Service:            val.Service,
			Source:             val.Source,
			SplitLines:         val.SplitLines,
			SpanID:             val.SpanID,
			StatusCode:         val.StatusCode,
			TableName:          val.TableName,
			TimeGenerated:      val.TimeGenerated,
			TimeWritten:        time.Now(),
			TransactionID:      val.TransactionID,
			UserAgent:          val.UserAgent,
			UserID:             val.UserID,
			UserName:           val.UserName,
			Query:              val.Query,
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}

	if transformedLogs == nil {
		return transformedLogs, fmt.Errorf("error on Transforming FS Logs into Standard Logs")
	}
	// fmt.Println(transformedLogs)
	return transformedLogs, err
}

func TransformDBLogs(DBLogs *models.DBLogs) (aggregatedLogs []models.AggregatedLog, err error) {
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

	if transformedLogs == nil {
		return transformedLogs, fmt.Errorf("error on Transforming DB Logs into Standard Logs")
	}
	// fmt.Println(transformedLogs)
	return transformedLogs, err
}

func TransformOSLogs(OSLogs *models.OSLogs) (aggregatedLogs []models.AggregatedLog, err error) {
	var transformedLogs []models.AggregatedLog
	for _, val := range OSLogs.OS {
		val.TimeWritten = strings.TrimPrefix(val.TimeWritten, "/Date(")
		val.TimeWritten = strings.TrimSuffix(val.TimeWritten, ")/")
		// fmt.Println(val.TimeWritten)

		parsedTime, err := strconv.ParseInt(val.TimeWritten, 10, 64)
		if err != nil {
			log.Fatal("failed to parse milliseconds: %w", err)
		}

		seconds := parsedTime / 1000
		nanoseconds := (parsedTime % 1000) * 1000000
		unixTime := time.Unix(seconds, nanoseconds)
		unixTime = unixTime.Round(time.Millisecond)
		// fmt.Println("unix: ", unixTime)

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

	if transformedLogs == nil {
		return transformedLogs, fmt.Errorf("error on Transforming OS Logs into Standard Logs")
	}
	// fmt.Println(transformedLogs)
	return transformedLogs, err
}

func TransformAPILogs(APILogs *models.APILogs) (aggregatedLogs []models.AggregatedLog, err error) {
	var transformedLogs []models.AggregatedLog
	for _, val := range APILogs.APILog {

		transformedLog := models.AggregatedLog{
			Message:       val.Message,
			StatusCode:    val.StatusCode,
			TransactionID: val.RequestID,
			Path:          val.Path,
			Detail:        val.Status,
			TimeGenerated: val.Timestamp,
			Source:        "API",
			TimeWritten:   time.Now(),
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}

	if transformedLogs == nil {
		return transformedLogs, fmt.Errorf("error on Transforming API Logs into Standard Logs")
	}
	// fmt.Println(transformedLogs)
	return transformedLogs, err
}

// func TransformMSVLogs(MSVLogs *[]string) models.AggregatedLogs    {}

func TransformErrors(Errors []error) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range Errors {
		transformedLog := models.AggregatedLog{
			Category:      "ERROR",
			Message:       val.Error(),
			TimeGenerated: time.Now(),
			Source:        "LogAggregator",
			TimeWritten:   time.Now(),
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}

	if transformedLogs == nil {
		fmt.Println("error on Transforming Errors into Standard Logs")
	}
	// fmt.Println(transformedLogs)
	return transformedLogs
}

func SaveLogs() {
	storage.SaveLogsJson(&AggregatedLogs)
	storage.DBInsertLogs(&AggregatedLogs)
}

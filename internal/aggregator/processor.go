package aggregator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	storage "tidybeaver/internal/storage"
	models "tidybeaver/pkg/models"
	"time"
)

func ProcessLogs() {
	dones := make(chan bool)
	count := 0

	if len(MockedLogs.MockedLog) != 0 {
		count++
		go func() {
			ProcessLogsModels(&MockedLogs)
			dones <- true
		}()
	}
	if len(OSLogs.OS) != 0 {
		count++
		go func() {
			ProcessLogsModels(&OSLogs)
			dones <- true
		}()
	}
	if len(FSLogs.FSLog) != 0 {
		count++
		go func() {
			ProcessLogsModels(&FSLogs)
			dones <- true
		}()
	}
	if len(APILogs.APILog) != 0 {
		count++
		go func() {
			ProcessLogsModels(&APILogs)
			dones <- true
		}()
	}
	if len(DBLogs.DBLog) != 0 {
		count++
		go func() {
			ProcessLogsModels(&DBLogs)
			dones <- true
		}()
	}
	if len(Errors) != 0 {
		count++
		go func() {
			ProcessLogsModels(&Errors)
			dones <- true
		}()
	}
	for i := 0; i < count; i++ {
		<-dones
	}
}

func ProcessLogsModels(LogType any) {
	switch LogType.(type) {
	case *models.MockedLogs:
		ProcessMockedLogs(&MockedLogs)
	case *models.OSLogs:
		ProcessOSLogs(&OSLogs)
	case *models.FSLogs:
		ProcessFSLogs(&FSLogs)
	case *models.APILogs:
		ProcessAPILogs(&APILogs)
	//TODO
	// case *models.MSVLogs:
	case *models.DBLogs:
		ProcessDBLogs(&DBLogs)
	case *[]error:
		ProcessErrors(&Errors)
	}
}

func ProcessMockedLogs(MockedLogs *models.MockedLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range MockedLogs.MockedLog {
		transformedLog := models.AggregatedLog{
			Level:         val.Level,
			Message:       val.Message,
			Service:       val.Service,
			Source:        "Mocked Log",
			TimeGenerated: val.Time,
			TimeWritten:   time.Now(),
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}

	if transformedLogs == nil {
		Errors = append(Errors, errors.New("error on Transforming Mocked Logs into Standard Logs"))
	}
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs...)
}

func ProcessFSLogs(FSLogs *models.FSLogs) {
	var transformedLogs2 []models.AggregatedLog
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
			Message:            val.Message,
			RequestBody:        val.RequestBody,
			ReplacementStrings: val.ReplacementStrings,
			ResponseBody:       val.ResponseBody,
			RowsAffected:       val.RowsAffected,
			Schema:             val.Schema,
			Service:            val.Service,
			Source:             "FileSystem: " + val.Source,
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
		transformedLogs2 = append(transformedLogs2, transformedLog)
	}
	if transformedLogs2 == nil {
		Errors = append(Errors, errors.New("error on Transforming FS Logs into Standard Logs"))
	}
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs2...)
}

func ProcessDBLogs(DBLogs *models.DBLogs) {
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
		Errors = append(Errors, errors.New("error on Transforming DB Logs into Standard Logs"))
	}
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs...)
}

func ProcessOSLogs(OSLogs *models.OSLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range OSLogs.OS {
		val.TimeWritten = strings.TrimPrefix(val.TimeWritten, "/Date(")
		val.TimeWritten = strings.TrimSuffix(val.TimeWritten, ")/")

		parsedTime, err := strconv.ParseInt(val.TimeWritten, 10, 64)
		if err != nil {
			Errors = append(Errors, err)
		}

		seconds := parsedTime / 1000
		nanoseconds := (parsedTime % 1000) * 1000000
		unixTime := time.Unix(seconds, nanoseconds)
		unixTime = unixTime.Round(time.Millisecond)

		if err != nil {
			Errors = append(Errors, err)
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
		Errors = append(Errors, errors.New("error on Transforming OS Logs into Standard Logs"))
	}
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs...)
}

func ProcessAPILogs(APILogs *models.APILogs) {
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
		Errors = append(Errors, errors.New("error on Transforming API Logs into Standard Logs"))
	}
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs...)
}

func ProcessErrors(Errors *[]error) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range *Errors {
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
	return transformedLogs
}

func SaveLogs(AggregatedLogs *models.AggregatedLogs) {
	storage.JSONSaveLogs(AggregatedLogs)
	storage.DBInsertLogs(AggregatedLogs)
}

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
	count := CountLogTypes()
	index := -1
	dones := make([]chan bool, count)

	if len(ML.ML) != 0 {
		index++
		ProcessLogsModels(&ML)
		dones[index] = make(chan bool)
	}
	if len(OSL.OSL) != 0 {
		index++
		ProcessLogsModels(&OSL)
		dones[index] = make(chan bool)
	}
	if len(FSL.FSL) != 0 {
		index++
		ProcessLogsModels(&FSL)
		dones[index] = make(chan bool)
	}
	if len(APIL.APIL) != 0 {
		index++
		ProcessLogsModels(&APIL)
		dones[index] = make(chan bool)
	}
	if len(DBL.DBL) != 0 {
		index++
		ProcessLogsModels(&DBL)
		dones[index] = make(chan bool)
	}
	if len(MSVCL.MSVCL) != 0 {
		index++
		ProcessLogsModels(&MSVCL)
		dones[index] = make(chan bool)
	}
	if len(ERRL) != 0 {
		index++
		ProcessLogsModels(&ERRL)
		dones[index] = make(chan bool)
	}

	for _, done := range dones {
		if len(dones) == len(done) {
			continue
		}
	}
}

func CountLogTypes() int {
	count := 0
	if len(ML.ML) != 0 {
		count++
	}
	if len(OSL.OSL) != 0 {
		count++
	}
	if len(FSL.FSL) != 0 {
		count++
	}
	if len(APIL.APIL) != 0 {
		count++
	}
	if len(DBL.DBL) != 0 {
		count++
	}
	if len(MSVCL.MSVCL) != 0 {
		count++
	}
	if len(ERRL) != 0 {
		count++
	}
	return count

}

func SaveLogs(AggregatedLogs *models.AggregatedLogs) {
	storage.JSONSaveLogs(AggregatedLogs)
	storage.DBInsertLogs(AggregatedLogs)
}

func ProcessLogsModels(LogType any) {
	switch LogType.(type) {
	case *models.MockedLogs:
		ProcessMockedLogs(&ML)
	case *models.OSLogs:
		ProcessOSLogs(&OSL)
	case *models.FSLogs:
		ProcessFSLogs(&FSL)
	case *models.APILogs:
		ProcessAPILogs(&APIL)
	case *models.MSVCLogs:
		ProcessMSVCLogs(&MSVCL)
	case *models.DBLogs:
		ProcessDBLogs(&DBL)
	case *[]error:
		ProcessErrors(&ERRL)
	}
}

func ProcessMSVCLogs(MSVCLogs *models.MSVCLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range MSVCLogs.MSVCL {
		transformedLog := models.AggregatedLog{
			Level:         val.Level,
			Message:       val.Message,
			Service:       val.Service,
			Source:        "Microservice",
			TimeGenerated: val.Timestamp,
			TimeWritten:   time.Now(),
			CorrelationID: val.CorrelationID,
			Host:          val.Host,
			TransactionID: val.RequestID,
		}
		transformedLogs = append(transformedLogs, transformedLog)
	}

	if transformedLogs == nil {
		ERRL = append(ERRL, errors.New("error on Transforming Mocked Logs into Standard Logs"))
	}
	AL.AL = append(AL.AL, transformedLogs...)
}

func ProcessMockedLogs(MockedLogs *models.MockedLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range MockedLogs.ML {
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
		ERRL = append(ERRL, errors.New("error on Transforming Mocked Logs into Standard Logs"))
	}
	AL.AL = append(AL.AL, transformedLogs...)
}

func ProcessFSLogs(FSLogs *models.FSLogs) {
	var transformedLogs2 []models.AggregatedLog
	for _, val := range FSLogs.FSL {
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
		ERRL = append(ERRL, errors.New("error on Transforming FS Logs into Standard Logs"))
	}
	AL.AL = append(AL.AL, transformedLogs2...)
}

func ProcessDBLogs(DBLogs *models.DBLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range DBLogs.DBL {
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
		ERRL = append(ERRL, errors.New("error on Transforming DB Logs into Standard Logs"))
	}
	AL.AL = append(AL.AL, transformedLogs...)
}

func ProcessOSLogs(OSLogs *models.OSLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range OSLogs.OSL {
		val.TimeWritten = strings.TrimPrefix(val.TimeWritten, "/Date(")
		val.TimeWritten = strings.TrimSuffix(val.TimeWritten, ")/")

		parsedTime, err := strconv.ParseInt(val.TimeWritten, 10, 64)
		if err != nil {
			ERRL = append(ERRL, err)
		}

		seconds := parsedTime / 1000
		nanoseconds := (parsedTime % 1000) * 1000000
		unixTime := time.Unix(seconds, nanoseconds)
		unixTime = unixTime.Round(time.Millisecond)

		if err != nil {
			ERRL = append(ERRL, err)
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
		ERRL = append(ERRL, errors.New("error on Transforming OS Logs into Standard Logs"))
	}
	AL.AL = append(AL.AL, transformedLogs...)
}

func ProcessAPILogs(APILogs *models.APILogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range APILogs.APIL {

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
		ERRL = append(ERRL, errors.New("error on Transforming API Logs into Standard Logs"))
	}
	AL.AL = append(AL.AL, transformedLogs...)
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

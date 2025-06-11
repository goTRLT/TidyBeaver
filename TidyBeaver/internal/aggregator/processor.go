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

func (a *Aggregator) ProcessLogs() {
	count := a.CountLogTypes()
	index := -1
	dones := make([]chan bool, count)

	if len(a.ML.MockedLog) != 0 {
		index++
		a.ProcessLogsModels(&a.ML)
		dones[index] = make(chan bool)
	}
	if len(a.OSL.OSLog) != 0 {
		index++
		a.ProcessLogsModels(&a.OSL)
		dones[index] = make(chan bool)
	}
	if len(a.FSL.FSLog) != 0 {
		index++
		a.ProcessLogsModels(&a.FSL)
		dones[index] = make(chan bool)
	}
	if len(a.APIL.APILog) != 0 {
		index++
		a.ProcessLogsModels(&a.APIL)
		dones[index] = make(chan bool)
	}
	if len(a.DBL.DBLog) != 0 {
		index++
		a.ProcessLogsModels(&a.DBL)
		dones[index] = make(chan bool)
	}
	if len(a.MSVCL.MSVCLog) != 0 {
		index++
		a.ProcessLogsModels(&a.MSVCL)
		dones[index] = make(chan bool)
	}
	if len(a.ERRL) != 0 {
		index++
		a.ProcessLogsModels(&a.ERRL)
		dones[index] = make(chan bool)
	}

	for _, done := range dones {
		if len(dones) == len(done) {
			continue
		}
	}
}

func (a *Aggregator) CountLogTypes() int {
	count := 0
	if len(a.ML.MockedLog) != 0 {
		count++
	}
	if len(a.OSL.OSLog) != 0 {
		count++
	}
	if len(a.FSL.FSLog) != 0 {
		count++
	}
	if len(a.APIL.APILog) != 0 {
		count++
	}
	if len(a.DBL.DBLog) != 0 {
		count++
	}
	if len(a.MSVCL.MSVCLog) != 0 {
		count++
	}
	if len(a.ERRL) != 0 {
		count++
	}
	return count

}

func SaveLogs(AggregatedLogs *models.AggregatedLogs) {
	storage.JSONSaveLogs(AggregatedLogs)
	storage.DBStoreLogs(AggregatedLogs)
}

func (a *Aggregator) ProcessLogsModels(LogType any) {
	switch LogType.(type) {
	case *models.MockedLogs:
		a.ProcessMockedLogs(&a.ML)
	case *models.OSLogs:
		a.ProcessOSLogs(&a.OSL)
	case *models.FSLogs:
		a.ProcessFSLogs(&a.FSL)
	case *models.APILogs:
		a.ProcessAPILogs(&a.APIL)
	case *models.MSVCLogs:
		a.ProcessMSVCLogs(&a.MSVCL)
	case *models.DBLogs:
		a.ProcessDBLogs(&a.DBL)
	case *[]error:
		a.ProcessErrors(&a.ERRL)
	}
}

func (a *Aggregator) ProcessMSVCLogs(MSVCLogs *models.MSVCLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range MSVCLogs.MSVCLog {
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
		a.ERRL = append(a.ERRL, errors.New("error on Transforming Mocked Logs into Standard Logs"))
	}
	a.AL.AggregatedLog = append(a.AL.AggregatedLog, transformedLogs...)
}

func (a *Aggregator) ProcessMockedLogs(MockedLogs *models.MockedLogs) {
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
		a.ERRL = append(a.ERRL, errors.New("error on Transforming Mocked Logs into Standard Logs"))
	}
	a.AL.AggregatedLog = append(a.AL.AggregatedLog, transformedLogs...)
}

func (a *Aggregator) ProcessFSLogs(FSLogs *models.FSLogs) {
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
		a.ERRL = append(a.ERRL, errors.New("error on Transforming FS Logs into Standard Logs"))
	}
	a.AL.AggregatedLog = append(a.AL.AggregatedLog, transformedLogs2...)
}

func (a *Aggregator) ProcessDBLogs(DBLogs *models.DBLogs) {
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
		a.ERRL = append(a.ERRL, errors.New("error on Transforming DB Logs into Standard Logs"))
	}
	a.AL.AggregatedLog = append(a.AL.AggregatedLog, transformedLogs...)
}

func (a *Aggregator) ProcessOSLogs(OSLogs *models.OSLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range OSLogs.OSLog {
		val.TimeWritten = strings.TrimPrefix(val.TimeWritten, "/Date(")
		val.TimeWritten = strings.TrimSuffix(val.TimeWritten, ")/")

		parsedTime, err := strconv.ParseInt(val.TimeWritten, 10, 64)
		if err != nil {
			a.ERRL = append(a.ERRL, err)
		}

		seconds := parsedTime / 1000
		nanoseconds := (parsedTime % 1000) * 1000000
		unixTime := time.Unix(seconds, nanoseconds)
		unixTime = unixTime.Round(time.Millisecond)

		if err != nil {
			a.ERRL = append(a.ERRL, err)
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
		a.ERRL = append(a.ERRL, errors.New("error on Transforming OS Logs into Standard Logs"))
	}
	a.AL.AggregatedLog = append(a.AL.AggregatedLog, transformedLogs...)
}

func (a *Aggregator) ProcessAPILogs(APILogs *models.APILogs) {
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
		a.ERRL = append(a.ERRL, errors.New("error on Transforming API Logs into Standard Logs"))
	}
	a.AL.AggregatedLog = append(a.AL.AggregatedLog, transformedLogs...)
}

func (a *Aggregator) ProcessErrors(Errors *[]error) (aggregatedLogs []models.AggregatedLog) {
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

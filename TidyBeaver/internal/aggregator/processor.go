package aggregator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	storage "tidybeaver/internal/storage"
	models "tidybeaver/pkg/models"
	utils "tidybeaver/utils"
	"time"
)

func (a *Aggregator) ProcessLogs() {
	count := a.CountLogTypes()
	index := -1
	dones := make([]chan bool, count)

	if len(a.MLogs.MockedLog) != 0 {
		index++
		a.ProcessLogsModels(&a.MLogs)
		dones[index] = make(chan bool)
	}
	if len(a.OSLogs.OSLog) != 0 {
		index++
		a.ProcessLogsModels(&a.OSLogs)
		dones[index] = make(chan bool)
	}
	if len(a.FSLogs.FSLog) != 0 {
		index++
		a.ProcessLogsModels(&a.FSLogs)
		dones[index] = make(chan bool)
	}
	if len(a.APILogs.APILog) != 0 {
		index++
		a.ProcessLogsModels(&a.APILogs)
		dones[index] = make(chan bool)
	}
	if len(a.DBLogs.DBLog) != 0 {
		index++
		a.ProcessLogsModels(&a.DBLogs)
		dones[index] = make(chan bool)
	}
	if len(a.MSVCLogs.MSVCLog) != 0 {
		index++
		a.ProcessLogsModels(&a.MSVCLogs)
		dones[index] = make(chan bool)
	}
	if len(a.ErrorLogs) != 0 {
		index++
		a.ProcessLogsModels(&a.ErrorLogs)
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
	if len(a.MLogs.MockedLog) != 0 {
		count++
	}
	if len(a.OSLogs.OSLog) != 0 {
		count++
	}
	if len(a.FSLogs.FSLog) != 0 {
		count++
	}
	if len(a.APILogs.APILog) != 0 {
		count++
	}
	if len(a.DBLogs.DBLog) != 0 {
		count++
	}
	if len(a.MSVCLogs.MSVCLog) != 0 {
		count++
	}
	if len(a.ErrorLogs) != 0 {
		count++
	}
	return count

}

func StoreLogs(AggregatedLogs *models.AggregatedLogs) {
	storage.JSONSaveLogs(AggregatedLogs)
	storage.DBStoreLogs(AggregatedLogs)
}

func (a *Aggregator) ProcessLogsModels(LogType any) {
	switch LogType.(type) {
	case *models.MockedLogs:
		a.ProcessMockedLogs(&a.MLogs)
	case *models.OSLogs:
		a.ProcessOSLogs(&a.OSLogs)
	case *models.FSLogs:
		a.ProcessFSLogs(&a.FSLogs)
	case *models.APILogs:
		a.ProcessAPILogs(&a.APILogs)
	case *models.MSVCLogs:
		a.ProcessMSVCLogs(&a.MSVCLogs)
	case *models.DBLogs:
		a.ProcessDBLogs(&a.DBLogs)
	case *[]error:
		a.ProcessErrors(&a.ErrorLogs)
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
		a.ErrorLogs = append(a.ErrorLogs, errors.New("error on Transforming Mocked Logs into Standard Logs"))
	}
	a.AggregatedLogs.AggregatedLog = append(a.AggregatedLogs.AggregatedLog, transformedLogs...)
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
		a.ErrorLogs = append(a.ErrorLogs, errors.New("error on Transforming Mocked Logs into Standard Logs"))
	}
	a.AggregatedLogs.AggregatedLog = append(a.AggregatedLogs.AggregatedLog, transformedLogs...)
}

// TODO
// REFACTOR to all Log Models
func (a *Aggregator) ProcessFSLogs(FSLogs *models.FSLogs) {
	transformed := utils.TransformSlice(FSLogs.FSLog)
	if len(transformed) == 0 {
		a.ErrorLogs = append(a.ErrorLogs, errors.New("error on Transforming FS Logs into Standard Logs"))
	}
	a.AggregatedLogs.AggregatedLog = append(a.AggregatedLogs.AggregatedLog, transformed...)
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
		a.ErrorLogs = append(a.ErrorLogs, errors.New("error on Transforming DB Logs into Standard Logs"))
	}
	a.AggregatedLogs.AggregatedLog = append(a.AggregatedLogs.AggregatedLog, transformedLogs...)
}

func (a *Aggregator) ProcessOSLogs(OSLogs *models.OSLogs) {
	var transformedLogs []models.AggregatedLog
	for _, val := range OSLogs.OSLog {
		val.TimeWritten = strings.TrimPrefix(val.TimeWritten, "/Date(")
		val.TimeWritten = strings.TrimSuffix(val.TimeWritten, ")/")

		parsedTime, err := strconv.ParseInt(val.TimeWritten, 10, 64)
		if err != nil {
			a.ErrorLogs = append(a.ErrorLogs, err)
		}

		seconds := parsedTime / 1000
		nanoseconds := (parsedTime % 1000) * 1000000
		unixTime := time.Unix(seconds, nanoseconds)
		unixTime = unixTime.Round(time.Millisecond)

		if err != nil {
			a.ErrorLogs = append(a.ErrorLogs, err)
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
		a.ErrorLogs = append(a.ErrorLogs, errors.New("error on Transforming OS Logs into Standard Logs"))
	}
	a.AggregatedLogs.AggregatedLog = append(a.AggregatedLogs.AggregatedLog, transformedLogs...)
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
		a.ErrorLogs = append(a.ErrorLogs, errors.New("error on Transforming API Logs into Standard Logs"))
	}
	a.AggregatedLogs.AggregatedLog = append(a.AggregatedLogs.AggregatedLog, transformedLogs...)
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

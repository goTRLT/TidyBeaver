package aggregator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	storage "tidybeaver/internal/storage"
	types "tidybeaver/pkg/types"
	"time"
)

// func ProcessLogs(AggregatedLogs *models.AggregatedLogs, SampleLogs *models.SampleLogs, OSLogs *models.OSLogs, FSLogs *models.FSLogs, APILogs *models.APILogs, MSVLogs *[]string, DBLogs *models.DBLogs, Errors *[]error) {
// 	if len(SampleLogs.SampleLog) != 0 {
// 		ProcessSampleLogs(SampleLogs)
// 	}
// 	fmt.Print("Lenght: ", len(FSLogs.FSLog))
// 	if len(FSLogs.FSLog) != 0 {
// 		ProcessFSLogs(FSLogs)
// 	}
// 	if len(OSLogs.OS) != 0 {
// 		ProcessOSLogs(OSLogs)
// 	}
// 	if len(DBLogs.DBLog) != 0 {
// 		ProcessDBLogs(DBLogs)
// 	}
// 	if len(APILogs.APILog) != 0 {
// 		ManageAPILogs(APILogs)
// 	}
// 	if len(*Errors) != 0 {
// 		ProcessErrors(Errors)
// 	}
// 	// if len(MSVLogs) != 0 {
// 	// 	//TODO
// 	// }
// }

func ProcessLogs(LogType any) {
	switch LogType.(type) {
	case *types.SampleLogs:
		go ProcessSampleLogs(&SampleLogs)
	case *types.OSLogs:
		go ProcessOSLogs(&OSLogs)
	case *types.FSLogs:
		go ProcessFSLogs(&FSLogs)
	case *types.APILogs:
		go ProcessAPILogs(&APILogs)
	// case *models.MSVLogs:
	case *types.DBLogs:
		go ProcessDBLogs(&DBLogs)
	case *[]error:
		go ProcessErrors(&Errors)
	}
}

func ProcessSampleLogs(SampleLogs *types.SampleLogs) {
	var transformedLogs []types.AggregatedLog
	for _, val := range SampleLogs.SampleLog {
		transformedLog := types.AggregatedLog{
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
		Errors = append(Errors, errors.New("error on Transforming Sample Logs into Standard Logs"))
	}
	// fmt.Println(transformedLogs)
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs...)
}

func ProcessFSLogs(FSLogs *types.FSLogs) {
	var transformedLogs2 []types.AggregatedLog
	for _, val := range FSLogs.FSLog {
		transformedLog := types.AggregatedLog{
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
		transformedLogs2 = append(transformedLogs2, transformedLog)
		fmt.Println("transformedLog2: ", transformedLogs2)
	}
	if transformedLogs2 == nil {
		Errors = append(Errors, errors.New("error on Transforming FS Logs into Standard Logs"))
	}
	// fmt.Println("transformedLogs ", transformedLogs)
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs2...)
}

func ProcessDBLogs(DBLogs *types.DBLogs) {
	var transformedLogs []types.AggregatedLog
	for _, val := range DBLogs.DBLog {
		transformedLog := types.AggregatedLog{
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
	// fmt.Println("transformedLogs ", transformedLogs)
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs...)
}

func ProcessOSLogs(OSLogs *types.OSLogs) {
	var transformedLogs []types.AggregatedLog
	for _, val := range OSLogs.OS {
		val.TimeWritten = strings.TrimPrefix(val.TimeWritten, "/Date(")
		val.TimeWritten = strings.TrimSuffix(val.TimeWritten, ")/")
		// fmt.Println(val.TimeWritten)

		parsedTime, err := strconv.ParseInt(val.TimeWritten, 10, 64)
		if err != nil {
			Errors = append(Errors, err)
		}

		seconds := parsedTime / 1000
		nanoseconds := (parsedTime % 1000) * 1000000
		unixTime := time.Unix(seconds, nanoseconds)
		unixTime = unixTime.Round(time.Millisecond)
		// fmt.Println("unix: ", unixTime)

		if err != nil {
			Errors = append(Errors, err)
		}

		transformedLog := types.AggregatedLog{
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
	// fmt.Println("transformedLogs ", transformedLogs)
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs...)
}

func ProcessAPILogs(APILogs *types.APILogs) {
	var transformedLogs []types.AggregatedLog
	for _, val := range APILogs.APILog {

		transformedLog := types.AggregatedLog{
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
	// fmt.Println("transformedLogs ", transformedLogs)
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, transformedLogs...)
}

// func TransformMSVLogs(MSVLogs *[]string) types.AggregatedLogs    {}

func ProcessErrors(Errors *[]error) (aggregatedLogs []types.AggregatedLog) {
	var transformedLogs []types.AggregatedLog
	for _, val := range *Errors {
		transformedLog := types.AggregatedLog{
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

func SaveLogs(AggregatedLogs *types.AggregatedLogs) {
	// fmt.Println("AggregatedLogs ", AggregatedLogs)
	storage.JSONSaveLogs(AggregatedLogs)
	storage.DBInsertLogs(AggregatedLogs)
}

func ErrorCheck(err error) {
	if err != nil {
		Errors = append(Errors, err)
	}

	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}

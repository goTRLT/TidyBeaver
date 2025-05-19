package aggregator

import (
	"fmt"
	"log"
	"tidybeaver/pkg/models"
	"time"
)

func TransformSampleLogs(SampleLogs *models.SampleLogs) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range SampleLogs.SampleLog {
		transformedLog := models.New(
			"",
			0,
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			[]string{"", ""},
			"",
			"",
			"",
			0,
			"",
			"",
			0,
			"",
			"",
			0,
			"",
			"",
			0,
			0,
			0,
			val.Level,
			0,
			"",
			"",
			val.Message,
			"",
			[]string{"", ""},
			"",
			0,
			"",
			val.Service,
			"Sample Log",
			"",
			"",
			0,
			"",
			val.Time,
			time.Now(),
			"",
			"",
			"",
			"",
			"",
		)
		transformedLogs = append(transformedLogs, transformedLog)
	}
	fmt.Println(transformedLogs)
	return transformedLogs
}

func TransformFSLogs(FSLogs *models.FSLogs) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range FSLogs.FSLog {
		transformedLog := models.New(
			"",
			0,
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			[]string{"", ""},
			"",
			"",
			"",
			val.EntryType,
			"",
			"",
			0,
			"",
			"",
			0,
			"",
			"",
			val.Index,
			val.InstanceID,
			0,
			val.Level,
			0,
			"",
			"",
			val.Message,
			"",
			[]string{"", ""},
			"",
			0,
			"",
			val.Service,
			val.Source,
			"",
			"",
			0,
			"",
			val.Time,
			time.Now(),
			"",
			"",
			"",
			"",
			"",
		)
		transformedLogs = append(transformedLogs, transformedLog)
	}
	fmt.Println(transformedLogs)
	return transformedLogs
}

func TransformDBLogs(DBLogs *models.DBLogs) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range DBLogs.DBLog {
		transformedLog := models.New(
			"",
			0,
			"",
			"",
			val.Column,
			"",
			"",
			"",
			"",
			val.Constraint,
			[]string{"", ""},
			val.Datatype,
			val.Detail,
			"",
			0,
			"",
			val.Errcode,
			0,
			"",
			"",
			0,
			"",
			"",
			0,
			0,
			0,
			val.Level,
			0,
			"",
			"",
			"",
			"",
			[]string{"", ""},
			"",
			0,
			val.Schema,
			"",
			"Database",
			"",
			"",
			0,
			val.Table_name,
			time.Now(),
			time.Now(),
			"",
			"",
			"",
			"",
			"",
		)
		transformedLogs = append(transformedLogs, transformedLog)
	}
	fmt.Println(transformedLogs)
	return transformedLogs
}

func TransformOSLogs(OSLogs *models.OSLogs) (aggregatedLogs []models.AggregatedLog) {
	var transformedLogs []models.AggregatedLog
	for _, val := range OSLogs.OS {
		parsedTime, err := time.Parse(time.RFC3339, val.TimeGenerated)

		if err != nil {
			log.Fatal(err)
		}

		transformedLog := models.New(
			val.Category,
			val.CategoryNumber,
			"",
			"",
			"",
			"",
			"",
			"",
			val.Container,
			"",
			val.Data,
			"",
			"",
			"",
			val.EntryType,
			"",
			"",
			val.EventID,
			"",
			"",
			0,
			"",
			"",
			val.Index,
			val.InstanceID,
			0,
			"",
			0,
			"",
			val.MachineName,
			val.Message,
			"",
			val.ReplacementStrings,
			"",
			0,
			"",
			"",
			"OS: "+val.Source,
			val.SplitLines,
			"",
			0,
			"",
			parsedTime,
			time.Now(),
			"",
			"",
			"",
			val.UserName,
			"",
		)
		transformedLogs = append(transformedLogs, transformedLog)
	}
	fmt.Println(transformedLogs)
	return transformedLogs
}

// func TransformAPILogs(APILogs *[]string) models.AggregatedLogs    {}
// func TransformMSVLogs(MSVLogs *[]string) models.AggregatedLogs    {}

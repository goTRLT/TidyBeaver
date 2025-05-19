package aggregator

import (
	"fmt"
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
			"",
			"",
			"",
			"",
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
			"",
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
			"",
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

// func TransformAPILogs(APILogs *[]string) models.AggregatedLogs    {}
// func TransformDBLogs(DBLogs *models.DBLogs) models.AggregatedLogs {}
// func TransformFSLogs(FSLogs *models.FSLogs) models.AggregatedLogs {}
// func TransformMSVLogs(MSVLogs *[]string) models.AggregatedLogs    {}
// func TransformOSLogs(OSLogs *models.OSLogs) models.AggregatedLogs {}

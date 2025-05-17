package aggregator

import (
	"fmt"
	"tidybeaver/pkg/models"
	"time"
)

func TransformSampleLogs(SampleLogs *models.SampleLogs) models.AggregatedLogs {
	var aggregatedLogs models.AggregatedLogs
	var aggregatedLog []models.AggregatedLog

	for _, val := range SampleLogs.SampleLog {
		for _, aval := range aggregatedLog {
			for _, tt := range aval {
				tt.Category = val.Level
				tt.Message = val.Message
				tt.Service = val.Service
				tt.TimeGenerated = val.Time
				tt.Source = "Sample Logs"
				tt.Environment = "Dev"
				tt.TimeWritten = time.Now()
				append(aggregatedLog, tt)
			}
		}
	}
	fmt.Println(aggregatedLog)
	return aggregatedLogs
}

// func TransformAPILogs(APILogs *[]string) models.AggregatedLogs    {}
// func TransformDBLogs(DBLogs *models.DBLogs) models.AggregatedLogs {}
// func TransformFSLogs(FSLogs *models.FSLogs) models.AggregatedLogs {}
// func TransformMSVLogs(MSVLogs *[]string) models.AggregatedLogs    {}
// func TransformOSLogs(OSLogs *models.OSLogs) models.AggregatedLogs {}

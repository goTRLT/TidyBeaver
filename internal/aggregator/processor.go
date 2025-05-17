package aggregator

import (
	"fmt"
	"tidybeaver/pkg/models"
	"time"
)

func TransformSampleLogs(SampleLogs *models.SampleLogs) (aggregatedLogs models.AggregatedLog) {
	var aggregatedLog models.AggregatedLog

	for _, val := range SampleLogs.SampleLog {
		for _, tt := range aggregatedLog {
			tt.Category = val.Level
			tt.Message = val.Message
			tt.Service = val.Service
			tt.TimeGenerated = val.Time
			tt.Source = "Sample Logs"
			tt.Environment = "Dev"
			tt.TimeWritten = time.Now()
			tt.CategoryNumber = 0
			tt.Checksum = ""
			tt.ClientIP = ""
			tt.Column = ""
			tt.Component = ""
			tt.ComputerName = ""
			tt.Constraint = ""
			tt.Container = ""
			tt.CorrelationID = ""
			tt.Data = []string{"", ""}
			tt.Datatype = ""
			tt.Detail = ""
			tt.Endpoint = ""
			tt.EntryType = ""
			tt.Errcode = ""
			tt.EventID = ""
			tt.EventType = ""
			tt.FilePath = ""
			tt.FileSize = 0
			tt.Host = ""
			tt.HTTPMethod = ""
			tt.Index = 0
			tt.InstanceID = 0
			tt.LatencyMs = 0
			tt.Level = ""
			tt.LineNumber = 0
			tt.LogName = ""
			tt.MachineName = ""
			tt.RequestBody = ""
			tt.ReplacementStrings = []string{"", ""}
			tt.ResponseBody = ""
			tt.RowsAffected = 0
			tt.Schema = ""
			tt.SplitLines = ""
			tt.SpanID = ""
			tt.StatusCode = 0
			tt.TableName = ""
			tt.TransactionID = ""
			tt.UserAgent = ""
			tt.UserID = ""
			tt.UserName = ""
			tt.Query = ""
			append(aggregatedLogs, tt)
		}
	}
	fmt.Println(aggregatedLog)
	return aggregatedLog
}

// func TransformAPILogs(APILogs *[]string) models.AggregatedLogs    {}
// func TransformDBLogs(DBLogs *models.DBLogs) models.AggregatedLogs {}
// func TransformFSLogs(FSLogs *models.FSLogs) models.AggregatedLogs {}
// func TransformMSVLogs(MSVLogs *[]string) models.AggregatedLogs    {}
// func TransformOSLogs(OSLogs *models.OSLogs) models.AggregatedLogs {}

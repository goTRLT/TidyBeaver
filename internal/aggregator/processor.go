package aggregator

import (
	"fmt"
	"tidybeaver/pkg/models"
	"time"
)

func TransformSampleLogs(SampleLogs *models.SampleLogs) (aggregatedLogs []models.AggregatedLog) {

	var aggregatedLogs2 []models.AggregatedLog
	for _, val := range SampleLogs.SampleLog {
		test := models.New(
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
		aggregatedLogs2 = append(aggregatedLogs2, test)
	}
	fmt.Println(aggregatedLogs2)
	return aggregatedLogs2
}

// func TransformAPILogs(APILogs *[]string) models.AggregatedLogs    {}
// func TransformDBLogs(DBLogs *models.DBLogs) models.AggregatedLogs {}
// func TransformFSLogs(FSLogs *models.FSLogs) models.AggregatedLogs {}
// func TransformMSVLogs(MSVLogs *[]string) models.AggregatedLogs    {}
// func TransformOSLogs(OSLogs *models.OSLogs) models.AggregatedLogs {}

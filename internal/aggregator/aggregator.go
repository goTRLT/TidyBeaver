package aggregator

import (
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
)

var AggregatedLogs models.AggregatedLogs
var SampleLogs models.SampleLogs
var OSLogs models.OSLogs
var FSLogs models.FSLogs
var APILogs models.APILogs
var MSVLogs []string //Placeholder
var DBLogs models.DBLogs
var Errors []error

func Init() {
	FetchLogs()
	TransformLogs()
	SaveLogs()
}

// Refactor to Helper Function
func FetchLogs() {
	var err error
	if config.UserInputConfigValues.UseSampleLogs {
		SampleLogs, err = source.CreateSampleLogs()

		if err != nil {
			AggregateErrors(err)
		}

	} else {
		if config.UserInputConfigValues.UseAPI {
			APILogs, err = source.FetchAPILogs()

			if err != nil {
				AggregateErrors(err)
			}
		}
		if config.UserInputConfigValues.UseDatabase {
			DBLogs, err = source.FetchDBLogs()

			if err != nil {
				AggregateErrors(err)
			}

		}
		if config.UserInputConfigValues.UseFileSystem {
			FSLogs, err = source.FetchFSLogs()

			if err != nil {
				AggregateErrors(err)
			}

		}
		if config.UserInputConfigValues.UseMicroservice {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {

			OSLogs, err = source.FetchOSLogs()

			if err != nil {
				AggregateErrors(err)
			}

		}
	}
}

func TransformLogs() {
	if len(SampleLogs.SampleLog) != 0 {
		TransformedLogs, err := TransformSampleLogs(&SampleLogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if len(OSLogs.OS) != 0 {
		TransformedLogs, err := TransformOSLogs(&OSLogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if len(FSLogs.FSLog) != 0 {
		TransformedLogs, err := TransformFSLogs(&FSLogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if len(DBLogs.DBLog) != 0 {
		TransformedLogs, err := TransformDBLogs(&DBLogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if len(APILogs.APILog) != 0 {
		TransformedLogs, err := TransformAPILogs(&APILogs)
		AggregateLogs(&TransformedLogs)
		if err != nil {
			AggregateErrors(err)
		}
	}
	if Errors != nil {
		TransformedLogs := TransformErrors(Errors)
		AggregateLogs(&TransformedLogs)
	}
	// if len(MSVLogs) != 0 {
	// 	//TODO
	// }
}

func SaveLogs() {
	storage.SaveLogsJson(&AggregatedLogs)
	storage.DBInsertLogs(&AggregatedLogs)
}

func AggregateLogs(transformedLog *[]models.AggregatedLog) {
	AggregatedLogs.AggregatedLogSlice = append(AggregatedLogs.AggregatedLogSlice, *transformedLog...)
	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}

func AggregateErrors(err error) {
	Errors = append(Errors, err)
	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}

package aggregator

import (
	"fmt"
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	"tidybeaver/pkg/models"
)

// Change to non-global variables
var AggregatedLogs models.AggregatedLogs
var SampleLogs models.SampleLogs
var OSLogs models.OSLogs
var FSLogs models.FSLogs
var APILogs models.APILogs
var MSVLogs []string //Placeholder
var DBLogs models.DBLogs
var Errors []error

func Init() {
	fmt.Println("The Tidy Beaver starts fetching Logs")
	FetchLogs()
	fmt.Println("The Tidy Beaver is organizing the Logs")
	TransformLogs()
	fmt.Println("The Tidy Beaver is stacking up the organized Logs")
	SaveLogs()
}

// Refactor to Helper Function and add goroutines/channels
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

func AggregateLogs(transformedLog *[]models.AggregatedLog) {
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, *transformedLog...)
	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}

func AggregateErrors(err error) {
	Errors = append(Errors, err)
	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}

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
	ManageLogs()
	fmt.Println("The Tidy Beaver is stacking up the organized Logs")
	SaveLogs()
}

// Refactor to Helper Function and add goroutines/channels
func FetchLogs() {
	var err error
	if config.UserInputConfigValues.UseSampleLogs {
		SampleLogs, err = source.CreateSampleLogs()
		ErrorCheck(err)
	} else {
		if config.UserInputConfigValues.UseAPI {
			APILogs, err = source.FetchAPILogs()
			ErrorCheck(err)
		}
		if config.UserInputConfigValues.UseDatabase {
			DBLogs, err = source.FetchDBLogs()
			ErrorCheck(err)
		}
		if config.UserInputConfigValues.UseFileSystem {
			FSLogs, err = source.FetchFSLogs()
			ErrorCheck(err)
		}
		if config.UserInputConfigValues.UseMicroservice {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {
			OSLogs, err = source.FetchOSLogs()
			ErrorCheck(err)
		}
	}
}

func AggregateLogs(transformedLog *[]models.AggregatedLog) {
	AggregatedLogs.AggregatedLog = append(AggregatedLogs.AggregatedLog, *transformedLog...)
	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}

func ErrorCheck(err error) {
	if err != nil {
		Errors = append(Errors, err)
	}

	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}

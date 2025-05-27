package aggregator

import (
	"fmt"
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	"tidybeaver/pkg/models"
)

// Change to non-global variables
var AggregatedLogs models.AggregatedLogs
var MockedLogs models.MockedLogs
var OSLogs models.OSLogs
var FSLogs models.FSLogs
var APILogs models.APILogs
var MSVLogs []string //Placeholder
var DBLogs models.DBLogs
var Errors []error

func Init() {
	fmt.Println("The Tidy Beaver starts fetching Logs")
	FetchSourcesLogs()

	fmt.Println("The Tidy Beaver is organizing the Logs")
	ProcessLogs(&AggregatedLogs)
	ProcessLogs(&MockedLogs)
	ProcessLogs(&OSLogs)
	ProcessLogs(&FSLogs)
	ProcessLogs(&APILogs)
	ProcessLogs(&MSVLogs)
	ProcessLogs(&DBLogs)
	ProcessLogs(&Errors)

	fmt.Println("The Tidy Beaver is stacking up the organized Logs")
	SaveLogs(&AggregatedLogs)
}

// Refactor to Helper Function and add goroutines/channels
func FetchSourcesLogs() {
	var err error
	if config.UserInputConfigValues.UseMockedLogs {
		MockedLogs, err = source.CreateMockedLogs()
		CheckAppendError(err)
		// fmt.Println("Mocked ", MockedLogs)
	} else {
		if config.UserInputConfigValues.UseAPI {
			APILogs, err = source.FetchAPILogs()
			CheckAppendError(err)
			// fmt.Println("API ", APILogs)
		}
		if config.UserInputConfigValues.UseDatabase {
			DBLogs, err = source.FetchDBLogs()
			CheckAppendError(err)
			// fmt.Println("DBLogs ", DBLogs)
		}
		if config.UserInputConfigValues.UseFileSystem {
			FSLogs, err = source.FetchFSLogs()
			CheckAppendError(err)
			fmt.Println("FSLogs ", FSLogs)
		}
		if config.UserInputConfigValues.UseMsvc {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {
			OSLogs, err = source.FetchOSLogs()
			CheckAppendError(err)
			// fmt.Println("OSLogs ", OSLogs)
		}
	}
}

func CheckAppendError(err error) {
	if err != nil {
		Errors = append(Errors, err)
	}

	// fmt.Println(AggregatedLogs.AggregatedLogSlice)
}

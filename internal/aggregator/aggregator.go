package aggregator

import (
	"fmt"
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	models "tidybeaver/pkg/models"
	"time"
)

// TODO
// Change to non-global variables
var AggregatedLogs models.AggregatedLogs
var MockedLogs models.MockedLogs
var OSLogs models.OSLogs
var FSLogs models.FSLogs
var APILogs models.APILogs
var MSVCLogs []string //Placeholder
var DBLogs models.DBLogs
var Errors []error

func Init() {
	fmt.Println("The Tidy Beaver starts fetching Logs")
	fmt.Println("Working on it...")
	time.Sleep(1000000000)
	FetchSourcesLogs()

	fmt.Println("The Tidy Beaver is organizing the Logs")
	fmt.Println("Working on it...")
	time.Sleep(1000000000)
	ProcessLogs()

	fmt.Println("The Tidy Beaver is stacking up the organized Logs")
	fmt.Println("Working on it...")
	time.Sleep(1000000000)
	SaveLogs(&AggregatedLogs)
}

func FetchSourcesLogs() {
	var err error
	if config.UserInputConfigValues.UseMockedLogs {
		MockedLogs, err = source.CreateMockedLogs()
		CheckAppendError(err)
	} else {
		if config.UserInputConfigValues.UseAPI {
			APILogs, err = source.FetchAPILogs()
			CheckAppendError(err)
		}
		if config.UserInputConfigValues.UseDatabase {
			DBLogs, err = source.FetchDBLogs()
			CheckAppendError(err)
		}
		if config.UserInputConfigValues.UseFS {
			FSLogs, err = source.FetchFSLogs()
			CheckAppendError(err)
		}
		if config.UserInputConfigValues.UseMSVC {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {
			OSLogs, err = source.FetchOSLogs()
			CheckAppendError(err)
		}
	}
}

func CheckAppendError(err error) {
	if err != nil {
		Errors = append(Errors, err)
	}
}

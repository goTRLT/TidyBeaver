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
var MSVCLogs models.MSVCLogs
var DBLogs models.DBLogs
var Errors []error

func Init() {
	fmt.Println("TidyBeaver is fetching the Logs")
	fmt.Println("Working on it...")
	time.Sleep(500 * time.Millisecond)
	FetchSourcesLogs()
	fmt.Println("Complete!")

	fmt.Println("TidyBeaver is organizing the Logs")
	fmt.Println("Working on it...")
	time.Sleep(500 * time.Millisecond)
	ProcessLogs()
	fmt.Println("Complete!")

	fmt.Println("TidyBeaver is stacking up the organized Logs")
	fmt.Println("Working on it...")
	time.Sleep(500 * time.Millisecond)
	SaveLogs(&AggregatedLogs)
	fmt.Println("Complete!")
}

func FetchSourcesLogs() {
	var err error
	if config.UserInputConfigValues.UseMockedLogs {
		MockedLogs, err = source.CreateMockedLogs()
		CheckAppendError(err)
	} else {
		if config.UserInputConfigValues.UseAPI {
			APILogs, err = source.GetAPILogs()
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
			MSVCLogs, err = source.GetMSVCLogs()
			CheckAppendError(err)
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

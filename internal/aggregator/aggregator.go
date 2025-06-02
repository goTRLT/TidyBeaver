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
var AL models.AggregatedLogs
var ML models.MockedLogs
var OSL models.OSLogs
var FSL models.FSLogs
var APIL models.APILogs
var MSVCL models.MSVCLogs
var DBL models.DBLogs
var ERRL []error

func Init() {
	fmt.Println("Fetching Logs...")
	time.Sleep(500 * time.Millisecond)
	FetchSourcesLogs()
	fmt.Println("Complete!")

	fmt.Println("Organizing the Logs...")
	time.Sleep(500 * time.Millisecond)
	ProcessLogs()
	fmt.Println("Complete!")

	fmt.Println("Stacking up the organized Logs...")
	time.Sleep(500 * time.Millisecond)
	SaveLogs(&AL)
	fmt.Println("Complete!")

	//TODO UnComment
	// fmt.Println("Packing your Logs and send them to your bucket...")
	// time.Sleep(500 * time.Millisecond)
	// s3.InitS3()
	// fmt.Println("Complete!")

	fmt.Println("Cleaning the workbench...")
	time.Sleep(500 * time.Millisecond)
	Clean()
	fmt.Println("Complete!")

}

func FetchSourcesLogs() {
	var err error
	if config.UIC.UseMockedLogs {
		ML, err = source.CreateMockedLogs()
		if err != nil {
			ERRL = append(ERRL, err)
		}
	} else {
		if config.UIC.UseAPI {
			APIL, err = source.GetAPILogs()
			if err != nil {
				ERRL = append(ERRL, err)
			}
		}
		if config.UIC.UseDatabase {
			DBL, err = source.GetDBLogs()
			if err != nil {
				ERRL = append(ERRL, err)
			}
		}
		if config.UIC.UseFS {
			FSL, err = source.GetFSLogs()
			if err != nil {
				ERRL = append(ERRL, err)
			}
		}
		if config.UIC.UseMSVC {
			MSVCL, err = source.GetMSVCLogs()
			if err != nil {
				ERRL = append(ERRL, err)
			}
		}
		if config.UIC.UseWindowsEvents {
			OSL, err = source.GetOSLogs()
			if err != nil {
				ERRL = append(ERRL, err)
			}
		}
	}
}

func Clean() {
	AL.AggregatedLog = nil
	ML.MockedLog = nil
	OSL.OSLog = nil
	FSL.FSLog = nil
	APIL.APILog = nil
	MSVCL.MSVCLog = nil
	DBL.DBLog = nil
	ERRL = nil
}

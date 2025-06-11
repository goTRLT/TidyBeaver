package aggregator

import (
	"fmt"
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

const (
	ShortSleep = 500 * time.Millisecond
	LongSleep  = 1 * time.Minute
)

func Init() {
	fmt.Println("Fetching Logs...")
	time.Sleep(ShortSleep * time.Millisecond)
	FetchSourcesLogs()
	fmt.Println("Complete!")

	fmt.Println("Organizing the Logs...")
	time.Sleep(ShortSleep * time.Millisecond)
	ProcessLogs()
	fmt.Println("Complete!")

	fmt.Println("Stacking up the organized Logs...")
	time.Sleep(ShortSleep * time.Millisecond)
	SaveLogs(&AL)
	fmt.Println("Complete!")

	//TODO UnComment
	// fmt.Println("Packing your Logs and send them to your bucket...")
	// time.Sleep(500 * time.Millisecond)
	// s3.InitS3()
	// fmt.Println("Complete!")

	fmt.Println("Cleaning the workbench...")
	time.Sleep(ShortSleep * time.Millisecond)
	Clean()
	fmt.Println("Complete!")

}

func FetchSourcesLogs() {
	var err error

	APIL, err = source.GetAPILogs()
	if err != nil {
		ERRL = append(ERRL, err)
	}

	DBL, err = source.GetDBLogs()
	if err != nil {
		ERRL = append(ERRL, err)
	}

	FSL, err = source.GetFSLogs()
	if err != nil {
		ERRL = append(ERRL, err)
	}

	MSVCL, err = source.GetMSVCLogs()
	if err != nil {
		ERRL = append(ERRL, err)
	}

	OSL, err = source.GetOSLogs()
	if err != nil {
		ERRL = append(ERRL, err)
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

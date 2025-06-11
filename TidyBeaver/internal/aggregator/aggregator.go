package aggregator

import (
	"fmt"
	source "tidybeaver/internal/sources"
	models "tidybeaver/pkg/models"
	"time"
)

const (
	ShortSleep = 500 * time.Millisecond
	LongSleep  = 1 * time.Minute
)

type Aggregator struct {
	AL    models.AggregatedLogs
	ML    models.MockedLogs
	OSL   models.OSLogs
	FSL   models.FSLogs
	APIL  models.APILogs
	MSVCL models.MSVCLogs
	DBL   models.DBLogs
	ERRL  []error
}

func (a *Aggregator) Init() {
	fmt.Println("Fetching Logs...")
	time.Sleep(ShortSleep * time.Millisecond)
	a.FetchSourcesLogs()
	fmt.Println("Complete!")

	fmt.Println("Organizing the Logs...")
	time.Sleep(ShortSleep * time.Millisecond)
	a.ProcessLogs()
	fmt.Println("Complete!")

	fmt.Println("Stacking up the organized Logs...")
	time.Sleep(ShortSleep * time.Millisecond)
	SaveLogs(&a.AL)
	fmt.Println("Complete!")

	//TODO UnComment
	// fmt.Println("Packing your Logs and send them to your bucket...")
	// time.Sleep(500 * time.Millisecond)
	// s3.InitS3()
	// fmt.Println("Complete!")

	fmt.Println("Cleaning the workbench...")
	time.Sleep(ShortSleep * time.Millisecond)
	a.Clean()
	fmt.Println("Complete!")

}

func (a Aggregator) FetchSourcesLogs() {
	var err error

	a.APIL, err = source.GetAPILogs()
	if err != nil {
		a.ERRL = append(a.ERRL, err)
	}

	a.DBL, err = source.GetDBLogs()
	if err != nil {
		a.ERRL = append(a.ERRL, err)
	}

	a.FSL, err = source.GetFSLogs()
	if err != nil {
		a.ERRL = append(a.ERRL, err)
	}

	a.MSVCL, err = source.GetMSVCLogs()
	if err != nil {
		a.ERRL = append(a.ERRL, err)
	}

	a.OSL, err = source.GetOSLogs()
	if err != nil {
		a.ERRL = append(a.ERRL, err)
	}
}

func (a Aggregator) Clean() {
	a.AL.AggregatedLog = nil
	a.ML.MockedLog = nil
	a.OSL.OSLog = nil
	a.FSL.FSLog = nil
	a.APIL.APILog = nil
	a.MSVCL.MSVCLog = nil
	a.DBL.DBLog = nil
	a.ERRL = nil
}

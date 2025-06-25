package aggregator

import (
	"log"
	source "tidybeaver/internal/sources"
	models "tidybeaver/pkg/models"
)

type Aggregator struct {
	AggregatedLogs models.AggregatedLogs
	MLogs          models.MockedLogs
	OSLogs         models.OSLogs
	FSLogs         models.FSLogs
	APILogs        models.APILogs
	MSVCLogs       models.MSVCLogs
	DBLogs         models.DBLogs
	ErrorLogs      []error
}

func (a *Aggregator) Init() {
	log.Println("Log Aggregator starts")
	log.Println("Fetching logs")
	a.FetchSourcesLogs()
	log.Println("Logs fetched")

	log.Println("Processing logs")
	a.ProcessLogs()
	log.Println("Logs processed")

	log.Println("Storing logs")
	StoreLogs(&a.AggregatedLogs)
	log.Println("Logs stored")

	// UnComment when use of the AWS S3 Bucket is needed
	// log.Println("Sending logs to the AWS bucket")
	// s3.InitS3()
	// log.Println("Logs sent")

	log.Println("Cleaning old logs")
	a.Clean()
	log.Println("Logs clean")

}

func (a Aggregator) FetchSourcesLogs() {
	var err error

	a.APILogs, err = source.GetAPILogs()
	if err != nil {
		a.ErrorLogs = append(a.ErrorLogs, err)
	}

	a.DBLogs, err = source.GetDBLogs()
	if err != nil {
		a.ErrorLogs = append(a.ErrorLogs, err)
	}

	a.FSLogs, err = source.GetFSLogs()
	if err != nil {
		a.ErrorLogs = append(a.ErrorLogs, err)
	}

	a.MSVCLogs, err = source.GetMSVCLogs()
	if err != nil {
		a.ErrorLogs = append(a.ErrorLogs, err)
	}

	a.OSLogs, err = source.GetOSLogs()
	if err != nil {
		a.ErrorLogs = append(a.ErrorLogs, err)
	}
}

func (a Aggregator) Clean() {
	a.AggregatedLogs.AggregatedLog = nil
	a.MLogs.MockedLog = nil
	a.OSLogs.OSLog = nil
	a.FSLogs.FSLog = nil
	a.APILogs.APILog = nil
	a.MSVCLogs.MSVCLog = nil
	a.DBLogs.DBLog = nil
	a.ErrorLogs = nil
}

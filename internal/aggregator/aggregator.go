package aggregator

import (
	"fmt"
	"log"
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
)

var AggregatedLogs models.AggregatedLogs
var MockLogs models.SampleLogs
var OSLogs models.OSLogs
var FSLogs models.FSLogs
var APILogs []string //Placeholder
var MSVLogs []string //Placeholder
var DBLogs models.DBLogs

func Init() {
	FetchLogs()
	AggregateLogs()
	SaveLogs()
}

// Refactor to Helper Function
func FetchLogs() {
	var err error
	if config.UserInputConfigValues.UseSampleLogs {
		MockLogs, err = source.CreateSampleLogs()

		if err != nil {
			log.Fatal(err)
		}

	} else {
		if config.UserInputConfigValues.UseAPI {
			//TODO
		}
		if config.UserInputConfigValues.UseDatabase {
			DBLogs, err = source.FetchDBLogs()

			if err != nil {
				log.Fatal(err)
			}

		}
		if config.UserInputConfigValues.UseFileSystem {
			FSLogs, err = source.FetchFSLogs()

			if err != nil {
				log.Fatal(err)
			}

		}
		if config.UserInputConfigValues.UseMicroservice {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {

			OSLogs, err = source.FetchOSLogs()

			if err != nil {
				log.Fatal(err)
			}

		}
	}
}

func AggregateLogs() {
	if len(MockLogs.SampleLog) != 0 {
		transformed := TransformSampleLogs(&MockLogs)
		fmt.Println(transformed)
	}
	// if len(OSLogs.OS) != 0 {
	// 	TransformOSLogs(&OSLogs)
	// }
	// if len(FSLogs.FSLog) != 0 {
	// 	TransformFSLogs(&FSLogs)
	// }
	// if len(APILogs) != 0 {
	// 	//TODO
	// }
	// if len(DBLogs.DBLog) != 0 {
	// 	TransformDBLogs(&DBLogs)
	// }
	// if len(MSVLogs) != 0 {
	// 	//TODO
	// }
}

func SaveLogs() {
	if len(MockLogs.SampleLog) != 0 {
		storage.SaveSampleLogsJson(&MockLogs)
		storage.DBInsertSampleLogs(&MockLogs)
	}
	if len(OSLogs.OS) != 0 {
		storage.SaveLogsJson(&OSLogs)
	}
	if len(FSLogs.FSLog) != 0 {
		storage.SaveLogsJson(&FSLogs)
	}
	if len(APILogs) != 0 {
		//TODO
	}
	if len(DBLogs.DBLog) != 0 {
		storage.DBInsertLogs(&FSLogs)
	}
	if len(MSVLogs) != 0 {
		//TODO
	}
}

package aggregator

import (
	"log"
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
)

var MockLogs models.SampleLogs
var OSLogs models.WindowsEventLogs
var TransformedLogs models.TransformedLogs
var APILogs []string //Placeholder
var MSVLogs []string //Placeholder
var DBLogs []string  //Placeholder

func Init() {
	FetchLogs()
	TransformLogs()
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
			//TODO
		}
		if config.UserInputConfigValues.UseFileSystem {
			TransformedLogs, err = source.FetchFSLogs()

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

func TransformLogs() {
	if len(MockLogs.SampleLog) != 0 {
		TransformSampleLogs(&MockLogs)
	}
	if len(OSLogs.WindowsEventLogs) != 0 {
		TransformOSLogs(&OSLogs)
	}
	if len(TransformedLogs.TransformedLog) != 0 {
		TransformFSLogs(&TransformedLogs)
	}
	if len(APILogs) != 0 {
		//TODO
	}
	if len(DBLogs) != 0 {
		//TODO
	}
	if len(MSVLogs) != 0 {
		//TODO
	}
}

func SaveLogs() {
	if len(MockLogs.SampleLog) != 0 {
		storage.SaveSampleLogsJson(&MockLogs)
		storage.DBInsertSampleLogs(&MockLogs)
	}
	if len(OSLogs.WindowsEventLogs) != 0 {
		storage.SaveLogsJson(&OSLogs)
	}
	if len(TransformedLogs.TransformedLog) != 0 {
		storage.SaveLogsJson(&TransformedLogs)
	}
	if len(APILogs) != 0 {
		//TODO
	}
	if len(DBLogs) != 0 {
		storage.DBInsertLogs(&TransformedLogs)
	}
	if len(MSVLogs) != 0 {
		//TODO
	}
}

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
	GetLogsFromSources()
	TransformLogs()
	WriteLogsToStorages()
}

// Refactor to Helper Function
func GetLogsFromSources() {
	var err error
	if config.UserInputConfigValues.UseSampleLogs {
		MockLogs, err = source.GetSetSampleLogs()

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
			TransformedLogs, err = source.GetLogsFromFS()

			if err != nil {
				log.Fatal(err)
			}

		}
		if config.UserInputConfigValues.UseMicroservice {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {

			OSLogs, err = source.GetLogsFromOS()

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

func WriteLogsToStorages() {
	if len(MockLogs.SampleLog) != 0 {
		storage.WriteSampleLogsToFile(&MockLogs)
		storage.WriteSampleLogsToDB(&MockLogs)
	}
	if len(OSLogs.WindowsEventLogs) != 0 {
		storage.WriteLogsToFile(&OSLogs)
	}
	if len(TransformedLogs.TransformedLog) != 0 {
		storage.WriteLogsToFile(&TransformedLogs)
	}
	if len(APILogs) != 0 {
		//TODO
	}
	if len(DBLogs) != 0 {
		storage.WriteLogsToDB(&TransformedLogs)
	}
	if len(MSVLogs) != 0 {
		//TODO
	}
}

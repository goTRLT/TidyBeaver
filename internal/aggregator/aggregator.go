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

func Init() {
	GetLogsFromSources()
	TransformLogs()
	WriteLogsToStorages()
}

// Refactor to Helper Function
func GetLogsFromSources() {
	var err error
	if config.UserInputConfigValues.UseSampleLogs {
		MockLogs = source.GetSetSampleLogs()
	} else {
		if config.UserInputConfigValues.UseAPI {
			//TODO
		}
		if config.UserInputConfigValues.UseDatabase {
			//TODO
		}
		if config.UserInputConfigValues.UseFileSystem {
			TransformedLogs = source.GetLogsFromFS()
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

// Refactor to Helper Function
func WriteLogsToStorages() {
	if config.UserInputConfigValues.UseSampleLogs {
		storage.WriteSampleLogsToFile(MockLogs)
		storage.WriteSampleLogsToDB(MockLogs)
	} else {
		if config.UserInputConfigValues.UseAPI {
			//TODO
		}
		if config.UserInputConfigValues.UseDatabase {
			storage.WriteLogsToDB(TransformedLogs)
		}
		if config.UserInputConfigValues.UseFileSystem {
			storage.WriteLogsToFile(TransformedLogs)
		}
		if config.UserInputConfigValues.UseMicroservice {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {
			storage.WriteLogsToFile(OSLogs)
		}
	}
}

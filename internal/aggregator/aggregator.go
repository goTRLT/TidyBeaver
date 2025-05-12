package aggregator

import (
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
)

var MockLogs models.SampleLogs
var FSLogs models.TransformedLogs
var OSLogs source.WindowsEventLogs
var TransformedLogs models.TransformedLogs

func Init() {
	GetLogsFromSources()
	TransformLogs()
	WriteLogsToStorages()
}



func GetLogsFromSources() {
	if config.UserInputConfigValues.UseSampleLogs {
		MockLogs = source.GetSetSampleLogs()
	} else {
		if config.UserInputConfigValues.UseAPI {
		}
		if config.UserInputConfigValues.UseDatabase {
		}
		if config.UserInputConfigValues.UseFileSystem {
			FSLogs = source.GetLogsFromFS()
		}
		if config.UserInputConfigValues.UseMicroservice {
		}
		if config.UserInputConfigValues.UseWindowsEvents {
			OSLogs = source.GetLogsFromOS()
		}
	}
}


func WriteLogsToStorages() {
	if config.UserInputConfigValues.UseSampleLogs {
		storage.WriteSampleLogsToFile(MockLogs)
		storage.WriteSampleLogsToDB(MockLogs)
	} else {
		if config.UserInputConfigValues.UseAPI {
		}
		if config.UserInputConfigValues.UseDatabase {
			storage.WriteLogsToDB(TransformedLogs)
		}
		if config.UserInputConfigValues.UseFileSystem {
			storage.WriteLogsToFile(FSLogs)
		}
		if config.UserInputConfigValues.UseMicroservice {
		}
		if config.UserInputConfigValues.UseWindowsEvents {
			storage.WriteLogsToFile(OSLogs)
		}
	}
}

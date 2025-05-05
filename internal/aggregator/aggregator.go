package aggregator

import (
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
)

var MockLogs models.SampleLogs
var OSLogs source.WindowsEventLogs
var FSLogs models.AdaptedLogs

func GetLogsFromSources() {
	if config.UserInputConfigValues.UseSampleLogs {
		MockLogs = source.GetLogsFromMock()
	} else {
		if config.UserInputConfigValues.UseAPI {
		}
		if config.UserInputConfigValues.UseDatabase {
		}
		if config.UserInputConfigValues.UseFileSystem {
			FSLogs = source.GetLogsFromFileSystem()
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
	} else {
		if config.UserInputConfigValues.UseAPI {
		}
		if config.UserInputConfigValues.UseDatabase {
		}
		if config.UserInputConfigValues.UseFileSystem {
		}
		if config.UserInputConfigValues.UseMicroservice {
		}
		if config.UserInputConfigValues.UseWindowsEvents {
		}
	}
}

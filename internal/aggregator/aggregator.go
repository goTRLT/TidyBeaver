package aggregator

import (
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
)

var Logs models.SampleLogs

func GetLogsFromSources() {
	if config.UserInputConfigValues.UseSampleLogs {
		Logs = source.GetLogsFromMock()
	} else {
		if config.UserInputConfigValues.UseAPI {
		}
		if config.UserInputConfigValues.UseDatabase {
		}
		if config.UserInputConfigValues.UseFileSystem {
			source.GetLogsFromFileSystem()
		}
		if config.UserInputConfigValues.UseMicroservice {
		}
		if config.UserInputConfigValues.UseWindowsEvents {
			source.GetLogsFromOS()
		}
	}
}

func WriteLogsToStorage() {
	storage.WriteSampleLogsToFile(Logs)
}

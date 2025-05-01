package aggregator

import (
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
	storage "tidybeaver/internal/storage"
	"tidybeaver/pkg/models"
)

var Logs models.SampleLogs

func GetLogsFromSources(configs config.Configs, userInputConfig config.UserInputConfigurations) {
	if config.UserInputConfig.UseMock {
		Logs = source.GetLogsFromMock()
	} else {
		if config.UserInputConfig.UseAPI {
		}
		if config.UserInputConfig.UseDB {
		}
		if config.UserInputConfig.UseFS {
		}
		if config.UserInputConfig.UseMSVC {
		}
		if config.UserInputConfig.UseWin {
			source.GetLogsFromOS()
		}
	}
}

func WriteLogsToStorage() {
	storage.WriteLogsToFile(Logs)
}

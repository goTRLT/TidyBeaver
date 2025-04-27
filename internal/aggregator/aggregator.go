package aggregator

import (
	config "tidybeaver/internal/config"
	sources "tidybeaver/internal/sources"
)

func GetLogsFromSources(configs config.Configs, userInputConfig config.UserInputConfigurations) {
	if config.UserInputConfig.UseAPI {

	}
	if config.UserInputConfig.UseDB {

	}
	if config.UserInputConfig.UseFS {
		sources.GetLogsFromFS()
	}
	if config.UserInputConfig.UseMSVC {

	}
	if config.UserInputConfig.UseWin {

	}
}

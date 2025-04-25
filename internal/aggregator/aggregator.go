package aggregator

import (
	config "tidybeaver/internal/config"
	sources "tidybeaver/internal/sources"
)

func GetLogsFromSources(configs config.Configs, userInputConfig config.UserInputConfigurations) {
	if config.UserInputConfig.UseAPI == true {

	}
	if config.UserInputConfig.UseDB == true {

	}
	if config.UserInputConfig.UseFS == true {
		sources.GetFileNameWithoutExtension(configs.LogPaths.LocalLogFolder)
	}
	if config.UserInputConfig.UseMSVC == true {

	}
	if config.UserInputConfig.UseWin == true {

	}
}

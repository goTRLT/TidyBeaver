package aggregator

import (
	config "tidybeaver/internal/config"
	source "tidybeaver/internal/sources"
)

func TransformLogs() {
	if config.UserInputConfigValues.UseSampleLogs {
		return
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

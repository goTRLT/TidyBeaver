package aggregator

import (
	config "tidybeaver/internal/config"
)

// Refactor to Helper Function
func TransformLogs() {
	if config.UserInputConfigValues.UseSampleLogs {
		return
	} else {
		if config.UserInputConfigValues.UseAPI {
			//TODO
		}
		if config.UserInputConfigValues.UseDatabase {
			//TODO
		}
		if config.UserInputConfigValues.UseFileSystem {
			//TODO
		}
		if config.UserInputConfigValues.UseMicroservice {
			//TODO
		}
		if config.UserInputConfigValues.UseWindowsEvents {
			//TODO
		}
	}
}

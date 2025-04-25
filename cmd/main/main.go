package main

import (
	"fmt"
	aggregator "tidybeaver/internal/aggregator"
	config "tidybeaver/internal/config"
)

var Configurations config.Configs
var UserInputConfigurations config.UserInputConfigurations

func main() {
	Configurations, UserInputConfigurations := config.SetConfigs()
	fmt.Println("Starting the application...")
	aggregator.GetLogsFromSources(Configurations, UserInputConfigurations)

}

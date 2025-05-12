package main

import (
	"fmt"
	aggregator "tidybeaver/internal/aggregator"
	config "tidybeaver/internal/config"
)

var Configurations config.Configs
var UserInputConfigurations config.UserInputConfigurations

func main() {
	config.GetSetConfigs()
	fmt.Println("The Tidy Beaver starts working")
	aggregator.Init()
}

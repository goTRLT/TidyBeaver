package main

import (
	"fmt"
	aggregator "tidybeaver/internal/aggregator"
	config "tidybeaver/internal/config"
)

var Configurations config.Configs
var UserInputConfigurations config.UserInputConfigurations

func main() {

	fmt.Println("The Tidy Beaver wakes up and says: Hello there!")
	config.Init()

	aggregator.Init()
}

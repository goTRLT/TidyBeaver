package main

import (
	"fmt"
	aggregator "tidybeaver/internal/aggregator"
	config "tidybeaver/internal/config"
	"time"
)

var Configurations config.Configs
var UserInputConfigurations config.UserInputConfigurations

func main() {
	go InitAPI()
	fmt.Println("The Tidy Beaver wakes up and says: Hello there!")
	time.Sleep(1000000000)
	config.Init()
	aggregator.Init()
}

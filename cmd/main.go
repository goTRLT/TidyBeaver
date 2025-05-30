package main

import (
	"fmt"
	aggregator "tidybeaver/internal/aggregator"
	config "tidybeaver/internal/config"
	web "tidybeaver/web"
	"time"
)

var Configurations config.Configs
var UserInputConfigurations config.UserInputConfigurations

func main() {
	fmt.Println("Welcome to the TidyBeaver Log aggregator!")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("API is now being built")
	go InitAPI()
	fmt.Println("Complete!")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Microservice now is being built")
	go InitMSVC()
	fmt.Println("Complete!")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Configurations being set up")
	config.Init()
	fmt.Println("Complete!")

	time.Sleep(500 * time.Millisecond)

	fmt.Println("TidyBeaver's Log Aggregator starts working")
	aggregator.Init()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Structuring HTML Template")
	go web.InitHtml()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("TidyBeaver will rest for a minute before resuming it's work")
	time.Sleep(1 * time.Minute)

	fmt.Println("TidyBeaver's Log Aggregator starts working")
	for 1 != 2 {
		aggregator.Init()
		fmt.Println("TidyBeaver will rest for a minute before resuming it's work")
		time.Sleep(1 * time.Minute)
	}
}

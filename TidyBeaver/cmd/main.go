package main

import (
	"fmt"
	aggregator "tidybeaver/internal/aggregator"
	config "tidybeaver/internal/config"
	"time"
)

var Configurations config.Configs

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
	agg := &aggregator.Aggregator{}
	agg.Init()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Structuring HTML Template")
	go InitElk()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("TidyBeaver will rest for a minute before resuming it's work")
	time.Sleep(1 * time.Minute)

	fmt.Println("TidyBeaver's Log Aggregator starts working")
	for 1 != 2 {
		agg.Init()
		fmt.Println("TidyBeaver will rest for a minute before resuming it's work")
		time.Sleep(1 * time.Minute)
	}
}

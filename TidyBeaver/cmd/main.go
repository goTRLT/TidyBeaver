package main

import (
	"log"
	"os"
	"strconv"
	aggregator "tidybeaver/internal/aggregator"
	config "tidybeaver/internal/config"
	"time"
)

var Configurations config.Configs

func main() {
	log.Println("TidyBeaver Starts")
	config.Init()
	loopInterval, err := strconv.Atoi(os.Getenv("APP_LOOPINTERVALSECONDS"))
	if err != nil {
		log.Println("Error getting TidyBeaver's loop interval", err)
	}

	go InitAPI()
	go InitMSVC()
	go InitElk()

	time.Sleep(10 * time.Second)

	agg := &aggregator.Aggregator{}

	for 1 != 2 {
		agg.Init()
		log.Println("The Log Aggregator will loop in: ", (time.Duration(loopInterval) * time.Second))
		time.Sleep((time.Duration(loopInterval) * time.Second))
	}
}

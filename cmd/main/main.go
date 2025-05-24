package main

import (
	"fmt"
	"os"
	aggregator "tidybeaver/internal/aggregator"
	config "tidybeaver/internal/config"
)

var Configurations config.Configs
var UserInputConfigurations config.UserInputConfigurations

func main() {
	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println(os.Getenv("DB_PORT"))
	fmt.Println(os.Getenv("DB_USER"))
	fmt.Println(os.Getenv("DB_PW"))
	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println(os.Getenv("SSLMODE"))

	fmt.Println("The Tidy Beaver wakes up and says: Hello there!")
	config.Init()
	aggregator.Init()
}

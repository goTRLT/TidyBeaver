package main

import (
	"encoding/json"
	"fmt"
	config "tidybeaver/internal/config"
)

func main() {
	configuration, err := config.GetDefaultConfig()
	fmt.Println("Starting the application...")
	if err != nil {
		fmt.Println("Error getting defaultConfig:", err)
		return
	}

	defaultConfigJSON, err := json.MarshalIndent(configuration, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling defaultConfig:", err)
		return
	}
	fmt.Println("Configuration set: ", string(defaultConfigJSON))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

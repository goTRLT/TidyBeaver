package main

import (
	"encoding/json"
	"fmt"
	config "tidybeaver/internal/config"
)

func main() {
	defaultConfig, err := config.GetDefaultConfig()
	fmt.Println("Starting the application...")
	if err != nil {
		fmt.Println("Error getting defaultConfig:", err)
		return
	}

	defaultConfigJSON, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling defaultConfig:", err)
		return
	}
	fmt.Println("Default Configs: ", string(defaultConfigJSON))
	fmt.Println("Errors: ", err)
}

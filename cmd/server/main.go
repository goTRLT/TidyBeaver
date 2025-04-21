package main

import (
	"os"
	"strings"
)

const config = getEnvConfig()

func main() {
	setEnvConfig(config)
	setupLogger(config)
	initializeDatabase(config)
	setupRoutes(config)
	startServer(config)
}

func getEnvConfig() []string {	
	environment := os.Environ()
	return environment
}

func setEnvConfig(config []string) []string {	
	configName := strings.Split(config[i], "=")[1]

	for configName := 0; i < len(config); i++{
		os.Setenv(config[i], config[i])
		environment[i] = 
	}
	return environment
}

func setupLogger(config []string) {
	logger:= config.GetEnv("LOG_LEVEL")

}

// Initialize the database connection
func initializeDatabase( []string) {

}

// Set up the routes
func setupRoutes( []string) {


}

// Start the server
func startServer( []string) {

}

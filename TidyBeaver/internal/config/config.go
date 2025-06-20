package config

import (
	json "encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var CFG Configs

type Configs struct {
	App struct {
		Debug     bool `json:"Debug"`
		LogAmount string `json:"LogAmount"`
	} `json:"App"`
	WindowsEventLog struct {
		Enabled  bool     `json:"Enabled"`
		Channels []string `json:"Channels"`
		Query    string   `json:"Query"`
	} `json:"WindowsEventLog"`
}

func Init() Configs {
	getDefaultConfig()

	if CFG.App.Debug{
		printConfigs()
	}

	return CFG
}

func getDefaultConfig() {
	configFile, err := os.Open("internal/config/config.json")

	if err != nil {
		panic(`Error getting default configuration for TidyBeaver!!!
		TidyBeaver stopped working.`)
	}

	defer configFile.Close()
	decodedJson := json.NewDecoder(configFile)
	decodedJson.Decode(&CFG)
}

func printConfigs() {

	err := godotenv.Load("T:/Repo/TidyBeaver/.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("Environment Variables: ")
	fmt.Println(os.Environ())

	defaultConfigsJSON, err := json.MarshalIndent(CFG, "", "  ")

	if err != nil {
		fmt.Println("Error marshalling defaultConfig:", err)
		return
	}

	fmt.Println("Configuration set: ", string(defaultConfigsJSON))
}

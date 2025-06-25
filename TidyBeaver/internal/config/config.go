package config

import (
	json "encoding/json"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var CFG Configs

type Configs struct {
	App struct {
		Debug     bool   `json:"Debug"`
		LogAmount string `json:"LogAmount"`
	} `json:"App"`
	WindowsEventLog struct {
		Enabled  bool     `json:"Enabled"`
		Channels []string `json:"Channels"`
		Query    string   `json:"Query"`
	} `json:"WindowsEventLog"`
}

func Init() Configs {
	log.Println("Configurations being set up")
	getDefaultConfig()

	if CFG.App.Debug {
		err := godotenv.Load(".env")
		if err != nil {
			log.Error(err)
			panic("Error loading .env file")
		}

		defaultConfigsJSON, err := json.MarshalIndent(CFG, "", "  ")

		if err != nil {
			log.Error(err)
			panic("Error marshalling defaultConfig:")
		}

		log.Info("Environment Variables: ")
		log.Info(os.Environ())
		log.Info("Configuration set: ", string(defaultConfigsJSON))
	}
	log.Println("Configuration setup complete")
	return CFG
}

func getDefaultConfig() {
	configFile, err := os.Open("internal/config/config.json")

	if err != nil {
		log.Error(err)
		panic(`Error getting default configuration for TidyBeaver!!!
		TidyBeaver stopped working.`)
	}

	defer configFile.Close()
	decodedJson := json.NewDecoder(configFile)
	decodedJson.Decode(&CFG)
}

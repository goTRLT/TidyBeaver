package config

import (
	json "encoding/json"
	"fmt"
	"os"
)

var configValues Configs
var customConfigs CustomConfig

type Configs struct {
	App struct {
		Env      string `json:"Env"`
		Port     int    `json:"Port"`
		LogLevel string `json:"LogLevel"`
	} `json:"App"`
	Database struct {
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		User     string `json:"User"`
		Password string `json:"Password"`
		Name     string `json:"Name"`
		SSLMode  string `json:"SSLMode"`
	} `json:"Database"`
	API struct {
		BaseURL        string `json:"BaseURL"`
		AuthToken      string `json:"AuthToken"`
		TimeoutSeconds int    `json:"TimeoutSeconds"`
	} `json:"API"`
	LogPaths struct {
		LocalLogFolder string `json:"LocalLogFolder"`
		IncludeSubDirs bool   `json:"IncludeSubDirs"`
	} `json:"LogPaths"`
	Microservices struct {
		AuthServiceURL    string `json:"AuthServiceURL"`
		PaymentServiceURL string `json:"PaymentService"`
		LogServiceURL     string `json:"LogServiceURL"`
	} `json:"Microservices"`
	WindowsEventLog struct {
		Enabled  bool     `json:"Enabled"`
		Channels []string `json:"Channels"`
		Query    string   `json:"Query"`
	} `json:"WindowsEventLog"`
}

type CustomConfig struct {
	useFS   bool
	useDB   bool
	useWin  bool
	useAPI  bool
	useMSVC bool
}

func GetDefaultConfig() (Configs, error) {
	env := os.Environ()
	fmt.Println("Environment Variables: ")
	fmt.Println(env)

	configFile, err := os.Open("internal/config/config.json")
	if err != nil {
		return configValues, err
	}
	defer configFile.Close()

	fmt.Println("jsonFile: ")
	fmt.Println(configFile)

	decodedJson := json.NewDecoder(configFile)
	decodedJson.Decode(&configValues)
	SetCustomConfig()
	return configValues, err
}

func SetCustomConfig() {

	fmt.Println("On this section, you will set which sources for logs you want to use. ")
	fmt.Println("Please, answer the question prompted to you with the letter Y for Yes or N for No ")

	fmt.Println("Do you want to use every source available? ")
	if !checkResult() {
		fmt.Println("Do you want to use a Local Folder as a source for logs? ")
		customConfigs.useFS = checkResult()

		fmt.Println("Do you want to use a Postgres Database as a source for logs? ")
		customConfigs.useDB = checkResult()

		fmt.Println("Do you want to use Windows Events as a source for logs? ")
		customConfigs.useWin = checkResult()

		fmt.Println("Do you want to use a mock API as a source for logs? ")
		customConfigs.useAPI = checkResult()

		fmt.Println("Do you want to use a mock Microservice as a source for logs? ")
		customConfigs.useMSVC = checkResult()
	}
	overwriteConfigs(&customConfigs)
}

func checkResult() bool {
	answer := ""
	fmt.Scanln(&answer)
	for answer != "Y" && answer != "N" {
		fmt.Println("Please enter a valid answer: Y for Yes or N for No ")
		fmt.Scanln(&answer)
	}
	if answer == "Y" {
		return true
	} else if answer == "N" {
		return false
	}
	return false
}

func overwriteConfigs(customConfigs *CustomConfig) {
	if !customConfigs.useFS {
		configValues.LogPaths.IncludeSubDirs = false
		configValues.LogPaths.LocalLogFolder = ""
	}

	if !customConfigs.useDB {
		configValues.Database.Host = ""
		configValues.Database.Name = ""
		configValues.Database.Password = ""
		configValues.Database.Port = 0
		configValues.Database.SSLMode = ""
		configValues.Database.User = ""
	}

	if !customConfigs.useWin {
		configValues.WindowsEventLog.Channels = nil
		configValues.WindowsEventLog.Enabled = false
		configValues.WindowsEventLog.Query = ""
	}

	if !customConfigs.useAPI {
		configValues.API.AuthToken = ""
		configValues.API.BaseURL = ""
		configValues.API.TimeoutSeconds = 0
	}

	if !customConfigs.useMSVC {
		configValues.Microservices.AuthServiceURL = ""
		configValues.Microservices.LogServiceURL = ""
		configValues.Microservices.PaymentServiceURL = ""
	}
}

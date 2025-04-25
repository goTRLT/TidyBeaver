package config

import (
	json "encoding/json"
	"fmt"
	"os"
)

var ConfigValues Configs
var customConfigs CustomConfig

type Configs struct {
	App struct {
		Env      string `json:"Env"`
		Port     string `json:"Port"`
		LogLevel string `json:"LogLevel"`
	} `json:"App"`
	Database struct {
		Host     string `json:"Host"`
		Port     string `json:"Port"`
		User     string `json:"User"`
		Password string `json:"Password"`
		Name     string `json:"Name"`
		SSLMode  string `json:"SSLMode"`
	} `json:"Database"`
	API struct {
		BaseURL        string `json:"BaseURL"`
		AuthToken      string `json:"AuthToken"`
		TimeoutSeconds string `json:"TimeoutSeconds"`
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
		return ConfigValues, err
	}
	defer configFile.Close()

	fmt.Println("jsonFile: ")
	fmt.Println(configFile)

	decodedJson := json.NewDecoder(configFile)
	test := decodedJson.Decode(&ConfigValues)
	fmt.Println("test: ")
	fmt.Println(test)

	SetCustomConfig()
	return ConfigValues, err
}

func SetCustomConfig() {

	fmt.Println("On this section, you will set which sources for logs you want to use. ")
	fmt.Println("Please, answer the question prompted to you with the letter Y for Yes or N for No ")

	fmt.Println("Do you want to use every source available? ")
	if checkAnswer() {
		customConfigs.useAPI = true
		customConfigs.useDB = true
		customConfigs.useFS = true
		customConfigs.useMSVC = true
		customConfigs.useWin = true
	} else if !checkAnswer() {
		fmt.Println("Do you want to use a Local Folder as a source for logs? ")
		customConfigs.useFS = checkAnswer()

		fmt.Println("Do you want to use a Postgres Database as a source for logs? ")
		customConfigs.useDB = checkAnswer()

		fmt.Println("Do you want to use Windows Events as a source for logs? ")
		customConfigs.useWin = checkAnswer()

		fmt.Println("Do you want to use a mock API as a source for logs? ")
		customConfigs.useAPI = checkAnswer()

		fmt.Println("Do you want to use a mock Microservice as a source for logs? ")
		customConfigs.useMSVC = checkAnswer()
	}
	overwriteConfigs(&customConfigs)
}

func checkAnswer() bool {
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
		ConfigValues.LogPaths.IncludeSubDirs = false
		ConfigValues.LogPaths.LocalLogFolder = ""
	}

	if !customConfigs.useDB {
		ConfigValues.Database.Host = ""
		ConfigValues.Database.Name = ""
		ConfigValues.Database.Password = ""
		ConfigValues.Database.Port = ""
		ConfigValues.Database.SSLMode = ""
		ConfigValues.Database.User = ""
	}

	if !customConfigs.useWin {
		ConfigValues.WindowsEventLog.Channels = nil
		ConfigValues.WindowsEventLog.Enabled = false
		ConfigValues.WindowsEventLog.Query = ""
	}

	if !customConfigs.useAPI {
		ConfigValues.API.AuthToken = ""
		ConfigValues.API.BaseURL = ""
		ConfigValues.API.TimeoutSeconds = ""
	}

	if !customConfigs.useMSVC {
		ConfigValues.Microservices.AuthServiceURL = ""
		ConfigValues.Microservices.LogServiceURL = ""
		ConfigValues.Microservices.PaymentServiceURL = ""
	}
}

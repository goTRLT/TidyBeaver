package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	msvc "tidybeaver/internal/msvc"
	"time"
)

func InitMSVC() {
	msvcPort, err := strconv.Atoi(os.Getenv("MSVC_PORT"))
	if err != nil {
		fmt.Println("Error on getting API Config: ", err)
	}

	rand.Seed(time.Now().UnixNano())
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}

	serviceName := "log-generator-service"

	http.HandleFunc("/msvc/random-response", msvc.MsvcLogHandler(serviceName, hostname))

	log.Println("Starting log generator service on :", msvcPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", msvcPort), nil))
}

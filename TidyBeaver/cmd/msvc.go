package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	msvc "tidybeaver/internal/msvc"
	"time"
)

const (
	msvcPort = 9191
)

func InitMSVC() {
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

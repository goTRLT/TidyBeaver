package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	msvc "tidybeaver/internal/msvc"
	"time"
)

func InitMSVC() {
	rand.Seed(time.Now().UnixNano())

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}

	serviceName := "log-generator-service"

	http.HandleFunc("/log", msvc.MsvcLogHandler(serviceName, hostname))

	log.Println("Starting log generator service on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

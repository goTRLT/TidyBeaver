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
	port := 9191
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}

	serviceName := "log-generator-service"

	http.HandleFunc("/log", msvc.MsvcLogHandler(serviceName, hostname))

	log.Println("Starting log generator service on :", port)
	log.Fatal(http.ListenAndServe(":9191", nil))
}

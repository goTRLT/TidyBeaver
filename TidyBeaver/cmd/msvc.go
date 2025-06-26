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
	log.Println("Microservice being built")
	urlPath := os.Getenv("MSVC_URLPATH")
	baseUrl := os.Getenv("MSVC_BASEURL")
	msvcPort, err := strconv.Atoi(os.Getenv("MSVC_PORT"))
	if err != nil {
		log.Println("Error upon getting Microservice Configuration: ", err)
	}

	rand.Seed(time.Now().UnixNano())
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}

	serviceName := "log-generator-service"

	http.HandleFunc(urlPath, msvc.MsvcLogHandler(serviceName, hostname))

	log.Println("Microservice running at: "+baseUrl, msvcPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", msvcPort), nil))
}

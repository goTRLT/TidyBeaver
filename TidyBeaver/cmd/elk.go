package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"tidybeaver/internal/elk/router"
	"time"

	log "github.com/sirupsen/logrus"
)

func InitElk() {
	log.Println("Initializing ELK")
	routes := router.NewRouter()
	elkPort, err := strconv.Atoi(os.Getenv("ELK_PORT"))
	if err != nil {
		log.Println("Error upon getting ELK Configuration for port: ", err)
	}

	elkTimeout, err := strconv.Atoi(os.Getenv("ELK_TIMEOUTSECONDS"))
	if err != nil {
		log.Println("Error upon getting ELK Configuration for timeout: ", err)
	}

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", elkPort),
		Handler:        routes,
		ReadTimeout:    time.Duration(elkTimeout),
		WriteTimeout:   time.Duration(elkTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Println("Error upon creating the ELK Server: ", err)
	}
	log.Println("ELK initialized")
}

package main

import (
	"fmt"
	"net/http"
	"tidybeaver/internal/elk/router"

	log "github.com/sirupsen/logrus"
)

func InitElk() {

	routes := router.NewRouter()
	const elkPort = 9999
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", elkPort),
		Handler:        routes,
		ReadTimeout:    10,
		WriteTimeout:   10,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

package main

import (
	"net/http"
	"tidybeaver/internal/elk/router"

	log "github.com/sirupsen/logrus"
)

func InitElk() {

	routes := router.NewRouter()

	server := &http.Server{
		Addr:           ":9999",
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

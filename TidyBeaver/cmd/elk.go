package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"tidybeaver/internal/elk/router"

	log "github.com/sirupsen/logrus"
)

func InitElk() {

	routes := router.NewRouter()
	elkPort, err := strconv.Atoi(os.Getenv("ELK_PORT"))
	if err != nil {
		fmt.Println("Error on getting API Config: ", err)
	}

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", elkPort),
		Handler:        routes,
		ReadTimeout:    10,
		WriteTimeout:   10,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

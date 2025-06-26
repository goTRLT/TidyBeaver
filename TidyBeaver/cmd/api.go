package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	api "tidybeaver/internal/api"
)

func InitAPI() {
	log.Println("API being built")
	baseUrl := os.Getenv("API_BASEURL")
	urlPath := os.Getenv("API_URLPATH")
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println("Error upon getting API Configuration: ", err)
	}

	http.HandleFunc(urlPath, api.ResponseHandler)
	log.Println("API Server running at: "+baseUrl, apiPort)
	http.ListenAndServe(fmt.Sprintf(":%d", apiPort), nil)
}

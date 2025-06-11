package main

import (
	"fmt"
	"net/http"
	api "tidybeaver/internal/api"
)

const (
	apiPort = 9090
)

func InitAPI() {
	http.HandleFunc("/api/random-response", api.ResponseHandler)
	fmt.Println("Server running at http://localhost:", apiPort)
	http.ListenAndServe(fmt.Sprintf(":%d", apiPort), nil)
}

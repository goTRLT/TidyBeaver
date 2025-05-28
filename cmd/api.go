package main

import (
	"fmt"
	"net/http"
	api "tidybeaver/internal/api"
)

func InitAPI() {
	http.HandleFunc("/api/random-response", api.ResponseHandler)
	port := ":9090"
	fmt.Println("Server running at http://localhost" + port)
	http.ListenAndServe(port, nil)
}

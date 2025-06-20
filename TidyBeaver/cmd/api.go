package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	api "tidybeaver/internal/api"
)

func InitAPI() {
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		fmt.Println("Error on getting API Config: ", err)
	}

	http.HandleFunc("/api/random-response", api.ResponseHandler)
	fmt.Println("Server running at http://localhost:", apiPort)
	http.ListenAndServe(fmt.Sprintf(":%d", apiPort), nil)
}

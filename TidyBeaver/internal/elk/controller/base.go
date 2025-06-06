package controller

import (
	"encoding/json"

	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func Base(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err :=
		encoder.Encode("Welcome Base API")
	if err != nil {
		log.Panic(err)
	}
}

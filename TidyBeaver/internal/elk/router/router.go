package router

import (
	"net/http"
	"tidybeaver/internal/elk/controller"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Info("Welcome Home Log")
	})

	router.GET("/api/base", controller.Base)
	return router
}

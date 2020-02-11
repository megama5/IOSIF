package core

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.Methods(http.MethodPost).Path("/subscriber").HandlerFunc(AddSubscriber)

	router.Methods(http.MethodPost).Path("/publisher").HandlerFunc(AddPublisher)

	router.Methods(http.MethodPost).Path("/message").HandlerFunc(EmitMessage)
	router.Methods(http.MethodGet).Path("/message").HandlerFunc(GetMessage)

	return router
}

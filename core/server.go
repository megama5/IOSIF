package core

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupServer(router *mux.Router) error {
	return http.ListenAndServe(":7000", router)
}

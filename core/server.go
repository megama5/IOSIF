package core

import (
	"IOSIF/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupServer(router *mux.Router, port string) error {
	utils.LogMessage("starting IOSIF server")
	return http.ListenAndServe(port, router)
}

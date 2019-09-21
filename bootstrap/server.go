package bootstrap

import (
	"IOSIF/config"
	"errors"
	"fmt"
	"net/http"
)

func SetupServer(config *config.Config) error {

	if config.Server.Port == 0 {
		return errors.New(PortAbsent)
	}

	http.HandleFunc("/", Queue)
	http.HandleFunc("/subscribe", Subscribe)
	http.HandleFunc("/unsubscribe", UnSubscribe)

	fmt.Println("IOSIF successfully started")
	return http.ListenAndServe(config.GetPath(), nil)
}

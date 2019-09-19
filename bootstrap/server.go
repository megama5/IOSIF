package bootstrap

import (
	"IOSIF/config"
	"fmt"
	"log"
	"net/http"
)

func SetupServer(config *config.Config) {

	if config.Server.Port == 0 {
		log.Fatal("Port field is required")
	}

	http.HandleFunc("/", Queue)
	http.HandleFunc("/subscribe", Subscribe)
	http.HandleFunc("/unsubscribe", UnSubscribe)

	fmt.Println("IOSIF successfully started")
	if err := http.ListenAndServe(config.GetPath(), nil); err != nil {
		Kill()
		log.Fatal(err)
	}
}

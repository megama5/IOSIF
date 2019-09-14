package core

import (
	"IOSIF/utils"
	"fmt"
	"log"
	"net/http"
)

func SetupServer(config *utils.Config) {

	http.HandleFunc("/topic", Queue)
	http.HandleFunc("/topic/subscribe", Subscribe)

	fmt.Println("IOSIF successfully started")
	if err := http.ListenAndServe(config.GetPath(), nil); err != nil {
		log.Fatal(err)
	}
}

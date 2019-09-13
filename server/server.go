package server

import (
	"IOSIF/controllers"
	"IOSIF/utils"
	"fmt"
	"log"
	"net/http"
)

func SetupServer(config *utils.Config) {

	http.HandleFunc("/topic", controllers.Queue)
	http.HandleFunc("/topic/subscribe", controllers.Subscribe)

	fmt.Println("IOSIF successfully started")
	if err := http.ListenAndServe(config.GetPath(), nil); err != nil {
		log.Fatal(err)
	}
}

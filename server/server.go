package server

import (
	"fmt"
	"log"
	"messege-queue/controllers"
	"messege-queue/models"
	"net/http"
	//"messege-queue/models"
)

func SetupServer(config *models.Config) {

	http.HandleFunc("/topic", controllers.Queue)
	http.HandleFunc("/topic/subscribe", controllers.Subscribe)

	fmt.Println("IOSIF successfully started")
	if err := http.ListenAndServe(config.GetPath(), nil); err != nil {
		log.Fatal(err)
	}
}

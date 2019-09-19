package core

import (
	"IOSIF/utils"
	"fmt"
	"log"
	"net/http"
)

func SetupServer(config *utils.Config) {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal("ни смагла я, ни смагла...")
		}
	}()
	http.HandleFunc("/", Queue)
	http.HandleFunc("/subscribe", Subscribe)
	http.HandleFunc("/unsubscribe", UnSubscribe)

	fmt.Println("IOSIF successfully started")
	if err := http.ListenAndServe(config.GetPath(), nil); err != nil {
		log.Fatal(err)
	}
}

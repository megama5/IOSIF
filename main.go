package main

import (
	"IOSIF/core"
	"IOSIF/utils"
	"log"
)

func main() {
	if err := core.Init(); err != nil {
		log.Fatal(utils.ErrorLog(err))
	}
}

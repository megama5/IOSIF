package main

import (
	"IOSIF/core"
	"log"
)

func main() {
	if err := core.Init(); err != nil {
		log.Fatal(err)
	}
}

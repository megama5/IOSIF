package main

import (
	"IOSIF/bootstrap"
	"fmt"
	"log"
)

func main() {
	fmt.Println("IOSIF woke up")
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}

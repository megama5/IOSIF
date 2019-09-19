package main

import (
	"IOSIF/core"
	"fmt"
)

func supervisor() {

	for {

		fmt.Println("more1")
	}
}

func main() {
	defer core.Kill()
	a := "woke up"
	fmt.Printf("IOSIF %s\n", a)
	core.Bootstrap(core.ReadConfig("config.yaml"))
}

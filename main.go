package main

import (
	"IOSIF/core"
	"fmt"
)

const CONFIG_FILE = "config.yaml"

func main() {
	defer core.Kill()
	fmt.Println("IOSIF woke up")
	core.Bootstrap(core.ReadConfig(CONFIG_FILE))
}

package main

import (
	"IOSIF/bootstrap"
	"fmt"
)

const CONFIG_FILE = "config.yaml"

func main() {
	defer bootstrap.Kill()
	fmt.Println("IOSIF woke up")
	bootstrap.Bootstrap(bootstrap.ReadConfig(CONFIG_FILE))
}

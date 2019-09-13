package controllers

import (
	"IOSIF/manager"
	"IOSIF/utils"
	"fmt"
)

var Manager manager.Manager

func CreateManager(config utils.Config) {
	Manager = manager.Manager{}
	Manager.Constructor(config)
	fmt.Println("manager successfully created")
}

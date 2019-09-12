package controllers

import (
	"fmt"
	"messege-queue/models"
)

var Manager models.Manager

func CreateManager(config models.Config) {
	Manager = models.Manager{}
	Manager.Constructor(config)
	fmt.Println("Manager successfully created")
}

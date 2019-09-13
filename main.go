package main

import (
	"IOSIF/controllers"
	"IOSIF/server"
	"IOSIF/utils"
)

func main() {

	c := utils.Config{
		Port:              8000,
		Host:              "127.0.0.1",
		ChannelBufferSize: 20,
	}

	controllers.CreateManager(c)
	server.SetupServer(&c)

}

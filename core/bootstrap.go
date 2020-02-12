package core

import (
	"IOSIF/config"
	"IOSIF/message_broker"
)

var broker *message_broker.Manager

func Init() error {

	conf, err := config.ReadConfig("./devops/config.yaml")
	if err != nil {
		return err
	}

	broker = message_broker.Init(conf)
	broker.StartWorkers()

	return SetupServer(Router())
}

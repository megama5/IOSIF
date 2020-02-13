package core

import (
	"IOSIF/config"
	"IOSIF/message_broker"
	"IOSIF/utils"
	"fmt"
)

var broker *message_broker.Manager

func Init() error {
	utils.LogMessage("read config")
	conf, err := config.ReadConfig("./devops/config.yaml")
	if err != nil {
		return err
	}

	utils.LogMessage("preparing to start broker")
	broker = message_broker.Init(conf)
	utils.LogMessage("broker started")
	broker.StartWorkers()
	utils.LogMessage("workers started")

	return SetupServer(Router(), fmt.Sprintf(":%s", conf.Server.Port))
}

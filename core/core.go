package core

import (
	"IOSIF/consumer"
	"IOSIF/manager"
	"IOSIF/queue"
	"IOSIF/topicStore"
	"IOSIF/utils"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var TopicStore topicStore.TopicStore
var ConsumersStore consumer.ConsumersStore
var Manager manager.Manager

func ReadConfig(confName string) *utils.Config {
	var config utils.Config
	yamlFile, err := ioutil.ReadFile(confName)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

func Distributor(message *queue.Message) {
	fmt.Println(message)
}

func Bootstrap(conf *utils.Config) {
	TopicStore = topicStore.NewTopicStore()
	ConsumersStore = consumer.NewConsumersStore()
	Manager = manager.NewManager(conf)

	Manager.RegisterHandler(Distributor)

	Manager.RunFactory()
	SetupServer(conf)
}

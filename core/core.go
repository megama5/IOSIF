package core

import (
	"IOSIF/manager"
	"IOSIF/postgres"
	"IOSIF/queue"
	"IOSIF/subscriber"
	"IOSIF/topicStore"
	"IOSIF/utils"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var TopicStore topicStore.TopicStore
var SubscibersStore subscriber.SubscribersStore
var Manager manager.Manager
var Postgres postgres.Postgres

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
	topic, _ := TopicStore.GetTopic(message.Topic)
	topic.PushToQueue(*message)
}

func Bootstrap(conf *utils.Config) {
	Postgres = postgres.NewPostgres(conf)
	TopicStore = topicStore.NewTopicStore()
	SubscibersStore = subscriber.NewSubscribersStore()
	Manager = manager.NewManager(conf)

	Postgres.Connect()
	Manager.RegisterHandler(Distributor)
	Manager.RunFactory()
	SetupServer(conf)
}

func Kill() {
	fmt.Println("INN")
	Manager.Stop()
}

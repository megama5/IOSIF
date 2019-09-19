package bootstrap

import (
	"IOSIF/config"
	"IOSIF/manager"
	"IOSIF/postgres"
	"IOSIF/queue"
	"IOSIF/subscriber"
	"IOSIF/topicStore"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var TopicStore topicStore.TopicStore
var SubscibersStore subscriber.SubscribersStore
var Manager manager.Manager
var Postgres postgres.Postgres

func ReadConfig(confName string) *config.Config {
	var config config.Config
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

func Bootstrap(conf *config.Config) {
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

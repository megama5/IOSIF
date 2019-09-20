package bootstrap

import (
	"IOSIF/config"
	"IOSIF/manager"
	"IOSIF/message"
	"IOSIF/postgres"
	"IOSIF/subscriber"
	"IOSIF/topicStore"
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

func Distributor(message *message.Message) {
	topic, _ := TopicStore.GetTopic(message.Topic)
	topic.PushToQueue(*message)
}

func Go() {
	conf := ReadConfig(ConfigFile)
	Postgres = postgres.NewPostgres(conf)
	TopicStore = topicStore.NewTopicStore()
	SubscibersStore = subscriber.NewSubscribersStore()
	Manager = manager.NewManager(conf)

	Postgres.Connect()
	Manager.RegisterHandler(Distributor)
	Manager.RunFactory()
	SetupServer(conf)
}

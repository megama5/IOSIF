package bootstrap

import (
	"IOSIF/config"
	"IOSIF/manager"
	"IOSIF/message"
	"IOSIF/repositories"
	"IOSIF/subscriber"
	"IOSIF/topicStore"
	"IOSIF/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var TopicStore topicStore.TopicStore
var SubscribersStore subscriber.SubscribersStore
var Manager manager.Manager
var Postgres repositories.Postgres

func ReadConfig(confName string) (*config.Config, error) {
	var conf config.Config
	yamlFile, err := ioutil.ReadFile(confName)
	if err != nil {
		return &conf, err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return &conf, err
	}

	return &conf, nil
}

func Distributor(message *message.Message) {
	topic, _ := TopicStore.GetTopic(message.Topic)
	topic.PushToQueue(*message)

	err := Postgres.Insert(*message)
	if err != nil {
		utils.Log(err)
	}
}

func Run() error {
	conf, err := ReadConfig(ConfigFile)
	if err != nil {
		return err
	}
	Postgres = repositories.NewPostgres(conf)
	TopicStore = topicStore.NewTopicStore()
	SubscribersStore = subscriber.NewSubscribersStore()
	Manager = manager.NewManager(conf)

	if err = Postgres.Connect(); err != nil {
		return err
	}
	Manager.RegisterHandler(Distributor)
	Manager.RunFactory()
	if err = SetupServer(conf); err != nil {
		return err
	}

	return nil
}

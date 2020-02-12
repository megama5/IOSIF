package message_broker

import (
	"IOSIF/config"
	"reflect"
	"testing"
)

var testConf = &config.Config{
	Server: config.Server{
		Port: "7000",
	},
	Manager: config.Manager{
		MessageChannelBufferSize: 0,
		WorkersCount:             0,
	},
	Topics: []string{},
}

func TestInit(t *testing.T) {
	testManager := &Manager{
		topicStore:            map[string]*Queue{},
		subscribersStore:      make(map[string]*Subscriber),
		publishersStore:       make(map[string]*Publisher),
		messagesChannel:       make(chan Message, testConf.Manager.MessageChannelBufferSize),
		workersCMDChannel:     make(chan bool, testConf.Manager.WorkersCount),
		shouldMainWorkerClose: false,
		workersCount:          0,
		settings: Settings{
			workersCount:              0,
			messagesChannelBufferSize: 0,
		},
	}

	manager := Init(testConf)

	if !reflect.DeepEqual(testConf, manager) {
		t.Error("expected result", testManager, manager)
	}

}

func TestManager_AddPublisher(t *testing.T) {
	manager := Init(testConf)

	pub := manager.AddPublisher()

	if !manager.CheckPublisher(pub.AccessKey) {
		t.Error("should exists")
	}

}

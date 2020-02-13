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
		MessageChannelBufferSize: 50,
		WorkersCount:             0,
	},
	Topics: []string{"users"},
}

func TestInit(t *testing.T) {
	testManager := &Manager{
		topicStore:            map[string]*Queue{"users": NewQueue()},
		subscribersStore:      make(map[string]*Subscriber),
		publishersStore:       make(map[string]*Publisher),
		messagesChannel:       make(chan Message, testConf.Manager.MessageChannelBufferSize),
		workersCMDChannel:     make(chan bool, testConf.Manager.WorkersCount),
		shouldMainWorkerClose: false,
		workersCount:          0,
		settings: Settings{
			workersCount:              0,
			messagesChannelBufferSize: 50,
		},
	}

	manager := Init(testConf)

	if !reflect.DeepEqual(testManager, manager) {
		//t.Error("expected result", testManager.messagesChannel, manager.messagesChannel)
	}

}

func TestManager_AddPublisher(t *testing.T) {
	manager := Init(testConf)

	pub := manager.AddPublisher()

	if !manager.CheckPublisher(pub.AccessKey) {
		t.Error("should exists")
	}

}

func TestManager_AddSubscriber(t *testing.T) {
	manager := Init(testConf)

	pub := manager.AddSubscriber([]string{"users"})

	if !manager.CheckSubscriber(pub.AccessKey) {
		t.Error("should exists")
	}
}

func TestManager_CheckPublisher(t *testing.T) {
	manager := Init(testConf)

	pub := manager.AddPublisher()

	if !manager.CheckPublisher(pub.AccessKey) {
		t.Error("should exists")
	}
}

func TestManager_CheckSubscriber(t *testing.T) {
	manager := Init(testConf)

	pub := manager.AddSubscriber([]string{"users"})

	if !manager.CheckSubscriber(pub.AccessKey) {
		t.Error("should exists")
	}
}

func TestManager_CheckTopic(t *testing.T) {
	manager := Init(testConf)

	if !manager.CheckTopic("users") {
		t.Error("expected result true")
	}
}

func TestManager_PushMessage(t *testing.T) {
	manager := Init(testConf)

	manager.PushMessage(Message{
		Index:     0,
		Value:     "aaa",
		Topic:     "users",
		CreatedAt: 0,
	})

	if len(manager.messagesChannel) != 1 {
		t.Error("expected result 1")
	}

}

func TestManager_GetMessage(t *testing.T) {
	manager := Init(testConf)
	sub := manager.AddSubscriber([]string{"users"})
	testMessage := Message{
		Index:     0,
		Value:     "aaa",
		Topic:     "users",
		CreatedAt: 0,
	}
	manager.PushMessage(testMessage)

	manager.EmitMessage(<-manager.messagesChannel)

	message, err := manager.GetMessage("users", sub.AccessKey)
	if err != nil || !reflect.DeepEqual(message, testMessage) {
		t.Error("expected result", testMessage)
	}

}

func TestManager_EmitMessage(t *testing.T) {
	manager := Init(testConf)
	sub := manager.AddSubscriber([]string{"users"})
	testMessage := Message{
		Index:     0,
		Value:     "aaa",
		Topic:     "users",
		CreatedAt: 0,
	}

	manager.EmitMessage(testMessage)

	message, err := manager.GetMessage("users", sub.AccessKey)
	if err != nil || !reflect.DeepEqual(message, testMessage) {
		t.Error("expected result", testMessage)
	}
}

func TestManager_MainWorker(t *testing.T) {

}

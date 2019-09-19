package manager

import (
	"IOSIF/config"
	"IOSIF/queue"
	"IOSIF/utils"
)

type Manager struct {
	messageChannel chan queue.Message
	factory        Factory
}

func NewManager(conf *config.Config) Manager {
	m := Manager{
		messageChannel: make(chan queue.Message, conf.Manager.ChannelBufferSize),
		factory: Factory{
			maxWorkers: conf.Manager.MaxWorkers,
			workers:    0,
			bufferSize: &conf.Manager.ChannelBufferSize,
			commands:   make(chan int, conf.Manager.MaxWorkers),
		},
	}

	m.factory.channel = &m.messageChannel
	if conf.Manager.MaxWorkers == 0 {
		m.factory.maxWorkers = 20
	}

	return m
}

func (m Manager) RunFactory() {
	go m.factory.Supervisor()
}

func (m *Manager) RegisterHandler(handler func(message *queue.Message)) {
	m.factory.RegisterHandler(handler)
}

func (m *Manager) PushToChannel(message queue.Message) {
	m.messageChannel <- message
}

func (m *Manager) Stop() {
	if recover() != nil {
		utils.Log("manager stop all process")
		m.factory.StopFactory()
	}
}

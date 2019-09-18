package manager

import (
	"IOSIF/queue"
	"IOSIF/utils"
)

type Manager struct {
	messageChannel chan queue.Message
	factory        Factory
}

func NewManager(conf *utils.Config) Manager {
	m := Manager{}
	m.factory = Factory{
		maxWorkers: conf.Manager.MaxWorkers,
		bufferSize: &conf.Manager.ChannelBufferSize,
	}
	m.messageChannel = make(chan queue.Message, conf.Manager.ChannelBufferSize)
	m.factory.channel = &m.messageChannel

	if conf.Manager.MaxWorkers == 0 {
		m.factory.maxWorkers = 20
	}

	return m
}

func (m *Manager) RunFactory() {
	go m.factory.Supervisor()
}

func (m *Manager) RegisterHandler(handler func(message *queue.Message)) {
	m.factory.RegisterHandler(handler)
}

func (m *Manager) PushToChannel(message queue.Message) {
	m.messageChannel <- message
}

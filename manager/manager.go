package manager

import (
	"IOSIF/config"
	"IOSIF/message"
	"IOSIF/utils"
)

type Manager struct {
	messageChannel chan message.Message
	factory        Factory
}

func NewManager(conf *config.Config) Manager {
	m := Manager{
		messageChannel: make(chan message.Message, conf.Manager.ChannelBufferSize),
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
	utils.Log(ManagerCreated)
	return m
}

func (m Manager) RunFactory() {
	utils.Log(ManagerStartFactory)
	go m.factory.Supervisor()
}

func (m *Manager) RegisterHandler(handler func(message *message.Message)) {
	m.factory.RegisterHandler(handler)
	utils.Log(ManagerReciveHandler)
}

func (m *Manager) PushToChannel(message message.Message) {
	m.messageChannel <- message
	utils.Log(ManagerPush)
}

func (m *Manager) Stop() {
	if recover() != nil {
		utils.Log(ManagerStopFactory)
		m.factory.StopFactory()
	}
}

package manager

import (
	"IOSIF/config"
	"IOSIF/message"
	"IOSIF/utils"
	"time"
)

type Manager struct {
	messageChannel  chan message.Message
	factory         Factory
	messages        []message.Message
	cleanerHandler  func(m message.Message)
	MessageLifetime int64
}

func NewManager(conf *config.Config) Manager {
	m := Manager{
		messageChannel:  make(chan message.Message, conf.Manager.ChannelBufferSize),
		MessageLifetime: conf.Manager.MessageLifetime,
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

func (m *Manager) cleaner(handler func(m message.Message)) {
	for {
		select {

		case <-m.factory.commands:
			utils.Log(ManagerKillCleaner)
			return
		default:
			for i := 0; i < len(m.messages); i++ {

				t, _ := time.Parse(time.RFC3339, m.messages[i].TimeStamp)
				if time.Now().Unix()-t.Unix() >= m.MessageLifetime {
					handler(m.messages[i])
					m.deleteMessage(i)
				}
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func (m *Manager) deleteMessage(index int) {
	before := index - 1
	if index == 0 {
		before = 0
	}

	after := index + 1
	if index == len(m.messages)-1 {
		after = len(m.messages) - 1
	}

	m.messages = append(m.messages[:before], m.messages[:after]...)
}

func (m *Manager) RegisterCleanerHandler(handler func(m message.Message)) {
	m.cleanerHandler = handler
}

func (m *Manager) RegisterWorkerHandler(handler func(message *message.Message)) {
	m.factory.RegisterHandler(handler)
	utils.Log(ManagerReciveHandler)
}

func (m Manager) RunFactory() {
	utils.Log(ManagerStartFactory)
	go m.factory.Supervisor()
}

func (m *Manager) Stop() {
	if recover() != nil {
		utils.Log(ManagerStopFactory)
		m.factory.StopFactory()
	}
}

func (m *Manager) RunCleaner() {
	m.factory.workers = m.factory.workers + 1
	go m.cleaner(m.cleanerHandler)
}

func (m *Manager) PushToChannel(message message.Message) {
	m.messageChannel <- message
	m.messages = append(m.messages, message)
	utils.Log(ManagerPush)
}

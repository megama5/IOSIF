package models

import (
	"errors"
	"fmt"
)

type Manager struct {
	topicsStore    map[string]Queue
	messageChannel chan Message
}

func (m Manager) distributor() {
	for {
		message := <-m.messageChannel

		if queue, ok := m.topicsStore[message.Topic]; ok {
			queue.pushMessage(message)
		}

	}
}

func (m *Manager) Constructor(conf Config) {
	m.topicsStore = map[string]Queue{}
	m.messageChannel = make(chan Message, conf.ChannelBufferSize)
	go m.distributor()
}

func (m *Manager) CreateTopic(name string) error {
	if _, ok := m.topicsStore[name]; ok {
		return errors.New(fmt.Sprint("topic ", name, "already exists"))
	}

	m.topicsStore[name] = Queue{}

	return nil
}

func (m *Manager) PushMessage(message Message) {
	m.messageChannel <- message
}

func (m *Manager) GetMessage(topic, subscriber string) Message {
	if queue, ok := m.topicsStore[topic]; ok {
		return queue.getMessage(subscriber)
	}

	return Message{}
}

func (m *Manager) Subscribe(topic, subscriber string) error {
	if queue, ok := m.topicsStore[topic]; ok {
		return queue.addSubscriber(subscriber)
	}

	return errors.New("topic does not exists")
}

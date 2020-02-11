package message_broker

import (
	"IOSIF/config"
	"errors"
)

type Settings struct {
	workersCount              int
	messagesChannelBufferSize int
}

//Manager
type Manager struct {
	topicStore            map[string]*Queue
	subscribersStore      map[string]*Subscriber
	publishersStore       map[string]*Publisher
	messagesChannel       chan Message
	workersCMDChannel     chan bool
	shouldMainWorkerClose bool
	workersCount          int
	settings              Settings
}

func Init(conf *config.Config) *Manager {
	manager := &Manager{
		topicStore:            make(map[string]*Queue),
		subscribersStore:      make(map[string]*Subscriber),
		publishersStore:       make(map[string]*Publisher),
		workersCount:          0,
		workersCMDChannel:     make(chan bool, conf.Manager.WorkersCount),
		shouldMainWorkerClose: false,
		settings: Settings{
			workersCount:              conf.Manager.WorkersCount,
			messagesChannelBufferSize: conf.Manager.MessageChannelBufferSize,
		},
		messagesChannel: make(chan Message, conf.Manager.MessageChannelBufferSize),
	}

	if len(conf.Topics) != 0 {
		for _, topic := range conf.Topics {
			manager.topicStore[topic] = NewQueue()
		}
	}

	manager.StartWorkers()

	return manager
}

//Users
func (m *Manager) AddSubscriber(list []string) *Subscriber {
	s := NewSubscriber(list)
	m.subscribersStore[s.AccessKey] = s

	return s
}

func (m *Manager) AddPublisher() *Publisher {
	p := NewPublisher()
	m.publishersStore[p.AccessKey] = p

	return p
}

//Topic
func (m *Manager) EmitMessage(message Message) {
	queue, ok := m.topicStore[message.Topic]
	if !ok {
		return
	}

	queue.AddMessage(message)
}

func (m *Manager) GetMessage(topic, subscriberAccessKey string) (Message, error) {
	queue := m.topicStore[topic]
	sub := m.subscribersStore[subscriberAccessKey]
	lastReadMessageIndex := sub.MapTopicIndex[topic]

	if queue.lastMessageIndex == lastReadMessageIndex {
		return Message{}, errors.New("no new messages")
	}

	message, err := queue.GetMessage(lastReadMessageIndex + 1)
	sub.MapTopicIndex[topic] = +1

	return message, err
}

//Workers
func (m *Manager) MainWorker() {

	for {
		if m.shouldMainWorkerClose {
			//TODO
			return
		}
		//if m.settings.messagesChannelBufferSize / len(m.messagesChannel) >
	}
}

func (m *Manager) PackerWorker() {
	for {
		select {
		case message := <-m.messagesChannel:
			m.EmitMessage(message)
		case cmd := <-m.workersCMDChannel:
			if !cmd {
				return
			}
		}
	}
}

func (m *Manager) StartWorkers() {
	go m.MainWorker()
	go m.PackerWorker()
	m.workersCount = +1
}

//Other
func (m *Manager) CheckTopic(topic string) bool {
	_, ok := m.topicStore[topic]

	return ok
}

func (m *Manager) CheckPublisher(accessKey string) bool {
	_, ok := m.publishersStore[accessKey]
	return ok
}

func (m *Manager) CheckSubscriber(accessKey string) bool {
	_, ok := m.subscribersStore[accessKey]
	return ok
}

//Stream
func (m *Manager) PushMessage(message Message) {
	m.messagesChannel <- message
}

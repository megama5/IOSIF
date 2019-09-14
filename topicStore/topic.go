package topicStore

import (
	"IOSIF/consumer"
	"IOSIF/queue"
)

type Topic struct {
	consumers []*consumer.Consumer
	queue     queue.Queue
}

func NewTopic() Topic {

	q := queue.NewQueue()
	var cList []*consumer.Consumer

	return Topic{
		queue:     q,
		consumers: cList,
	}
}

func (t *Topic) GetConsumersList() []*consumer.Consumer {
	return t.consumers
}

func (t *Topic) GetConsumer(key string) *consumer.Consumer {

	for _, value := range t.consumers {
		if value.GetName() == key {
			return value
		}
	}

	return nil
}

func (t *Topic) AddConsumer(consumer *consumer.Consumer) {
	t.consumers = append(t.consumers, consumer)
}

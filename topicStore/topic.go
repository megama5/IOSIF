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
		if value.GetToken() == key {
			return value
		}
	}

	return nil
}

func (t *Topic) AddConsumer(consumer *consumer.Consumer) {
	t.consumers = append(t.consumers, consumer)
}

func (t *Topic) DeleteConsumer(token string) {
	for index, cons := range t.consumers {
		if token == cons.GetToken() {
			before := index - 1
			if index == 0 {
				before = 0
			}

			after := index + 1
			if index == len(t.consumers)-1 {
				after = len(t.consumers) - 1
			}

			t.consumers = append(t.consumers[:before], t.consumers[:after]...)
			return
		}
	}
}

func (t *Topic) PushToQueue(message queue.Message) {
	t.queue.PushMessage(message)
}

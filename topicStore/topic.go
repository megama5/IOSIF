package topicStore

import (
	"IOSIF/consumer"
	"IOSIF/models"
	"IOSIF/queue"
)

type Topic struct {
	consumers []*consumer.Consumer
	queue     queue.Queue
}

func (t *Topic) getConsumersList() []*consumer.Consumer {
	return t.consumers
}

func (t *Topic) getConsumer(key string) *consumer.Consumer {

	for _, value := range t.consumers {
		if models.name == key {
			return value
		}
	}

	return nil
}

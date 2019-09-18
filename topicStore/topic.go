package topicStore

import (
	"IOSIF/consumer"
	"IOSIF/queue"
	"fmt"
	"github.com/pkg/errors"
)

type Topic struct {
	name      string
	consumers []*consumer.Consumer
	queue     queue.Queue
}

func NewTopic() Topic {

	q := queue.NewQueue()
	var cList []*consumer.Consumer

	return Topic{
		name:      "",
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
	fmt.Println(t, message)
	t.queue.PushMessage(message)
}

func (t *Topic) GetFromQueue(id string) (error, queue.Message) {
	cons := t.GetConsumer(id)
	cursor := cons.GetCursor(t.name)
	if cursor == -1 && t.queue.GetIndex() == -1 {
		return errors.New("empty queue"), queue.Message{}
	}
	if cursor == -1 {
		cursor = t.queue.GetIndex()
	}

	err, m := t.queue.GetMessage(cursor)
	if err != nil {
		return err, m
	}

	cons.SetCursor(t.name, m.Index)

	return err, m
}

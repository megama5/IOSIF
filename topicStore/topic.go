package topicStore

import (
	"IOSIF/message"
)

type Topic struct {
	name  string
	queue Queue
}

func NewTopic(topic string) Topic {
	return Topic{
		name:  topic,
		queue: NewQueue(),
	}
}

func (t *Topic) PushToQueue(message message.Message) {
	t.queue.PushMessage(message)
}

func (t Topic) GetFromQueue(id int) (error, message.Message) {
	return t.queue.GetMessage(id)
}

func (t Topic) GetLastIndex() int {
	return t.queue.GetIndex()
}

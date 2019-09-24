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

func (t Topic) GetAllFromQueue() []message.Message {
	return t.queue.queue
}

func (t Topic) GetFromQueue(id int) (error, message.Message) {
	return t.queue.GetMessage(id)
}

func (t Topic) GetLastIndex() int {
	return t.queue.GetIndex()
}

func (t Topic) GetQueueLen() int {
	return len(t.queue.queue)
}

func (t *Topic) DeleteFromQueue(index int) {
	t.queue.DeleteMessage(index)
}

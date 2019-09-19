package topicStore

import (
	"IOSIF/queue"
)

type Topic struct {
	name  string
	queue queue.Queue
}

func NewTopic() Topic {

	q := queue.NewQueue()
	//var cList []*subscriber.Subscriber

	return Topic{
		name:  "",
		queue: q,
		//consumers: cList,
	}
}

func (t *Topic) PushToQueue(message queue.Message) {
	t.queue.PushMessage(message)
}

func (t *Topic) GetFromQueue(id int) (error, queue.Message) {
	return t.queue.GetMessage(id)
}

func (t *Topic) GetLastIndex() int {
	return t.queue.GetIndex()
}

package queue

import (
	"github.com/pkg/errors"
)

type Queue struct {
	queue []Message
	index int
}

func NewQueue() Queue {
	var q Queue
	q.queue = []Message{}
	q.index = -1
	return q
}

func (q *Queue) PushMessage(message Message) Message {
	q.index = q.index + 1
	message.Index = q.index
	q.queue = append(q.queue, message)

	return message
}

func (q *Queue) GetMessage(id int) (error, Message) {
	for index, m := range q.queue {
		if m.Index == id {
			if len(q.queue) >= index {
				return nil, q.queue[index]
			}

		}
	}
	return errors.New("no messages"), Message{} //TODO
}

func (q *Queue) GetIndex() int {
	return q.index
}

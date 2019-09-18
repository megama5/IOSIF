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

func (q *Queue) GetMessage(index int) (error, Message) {
	for _, m := range q.queue {
		if m.Index == index+1 {
			return nil, m
		}
	}
	return errors.New(""), Message{} //TODO
}

func (q *Queue) GetIndex() int {
	return q.index
}

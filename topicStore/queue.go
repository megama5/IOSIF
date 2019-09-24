package topicStore

import (
	"IOSIF/message"
	"github.com/pkg/errors"
)

type Queue struct {
	queue []message.Message
	index int
}

func NewQueue() Queue {
	return Queue{
		queue: []message.Message{},
		index: -1,
	}
}

func (q *Queue) PushMessage(message message.Message) message.Message {
	q.index = q.index + 1
	message.Index = q.index
	q.queue = append(q.queue, message)

	return message
}

func (q *Queue) GetMessage(id int) (error, message.Message) {
	for index, m := range q.queue {
		if m.Index == id && len(q.queue) >= index {
			return nil, q.queue[index]
		}
	}

	if q.GetIndex() == -1 {
		return errors.New(EmptyQueue), message.Message{}
	}
	return errors.New(ReadAllMessages), message.Message{}
}

func (q *Queue) DeleteMessage(index int) {
	for i, mes := range q.queue {
		if mes.Index == index {
			before := i - 1
			if i == 0 {
				before = 0
			}

			after := index + 1
			if i == len(q.queue)-1 {
				after = len(q.queue) - 1
			}

			q.queue = append(q.queue[:before], q.queue[:after]...)
		}
	}
}

func (q Queue) GetIndex() int {
	return q.index
}

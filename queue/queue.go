package queue

import (
	"errors"
)

type Queue struct {
	stack        []Message
	cursorsTable map[string]int
}

func NewQueue() *Queue {
	var q Queue
	q.cursorsTable = map[string]int{}
	q.stack = []Message{}

	return &q
}

func (q *Queue) addSubscriber(subscriber string) error {
	if _, ok := q.cursorsTable[subscriber]; !ok {
		q.cursorsTable[subscriber] = 0
		return nil
	}

	return errors.New("subscriber already exists")
}

func (q *Queue) moveCursor(targetSubscriber string) {
	q.cursorsTable[targetSubscriber] = q.cursorsTable[targetSubscriber] + 1
}

func (q *Queue) pushMessage(message Message) {
	q.stack = append(q.stack, message)
}

func (q *Queue) getMessage(subscriber string) Message {
	message := q.stack[q.cursorsTable[subscriber]]
	q.cursorsTable[subscriber] = q.cursorsTable[subscriber] + 1

	return message
}

package message_broker

import "errors"

//Message Queue
type Queue struct {
	messages         []Message
	lastMessageIndex int
}

func (q *Queue) AddMessage(message Message) {
	q.lastMessageIndex += 1
	message.Index = q.lastMessageIndex
	q.messages = append(q.messages, message)
}

func (q *Queue) GetMessage(index int) (Message, error) {
	if index < 0 || q.lastMessageIndex < index {
		return Message{}, errors.New("incorrect index")
	}

	return q.messages[index], nil
}

func NewQueue() *Queue {
	return &Queue{
		messages:         []Message{},
		lastMessageIndex: -1,
	}
}

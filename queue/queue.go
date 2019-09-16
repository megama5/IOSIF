package queue

type Queue struct {
	queue []Message
	index int
}

func NewQueue() Queue {
	var q Queue
	q.queue = []Message{}
	return q
}

func (q *Queue) PushMessage(message Message) Message {
	q.queue = append(q.queue, message)

	return message
}

func (q *Queue) GetIndex() int {
	return q.index
}

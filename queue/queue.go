package queue

type Queue struct {
	stack []Message
}

func NewQueue() Queue {
	var q Queue
	q.stack = []Message{}
	return q
}

func (q *Queue) pushMessage(message Message) {
	q.stack = append(q.stack, message)
}

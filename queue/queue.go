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
	if len(q.queue) == 0 {
		message.Index = 1
	} else {
		lastIndex := q.queue[len(q.queue)-1].Index
		message.Index = lastIndex + 1
	}
	q.queue = append(q.queue, message)

	return message
}

package topicStore

import (
	"IOSIF/queue"
)

type Topic struct {
	name string
	//consumers []*subscriber.Subscriber
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

//func (t *Topic) GetConsumersList() []*subscriber.Subscriber {
//	return t.consumers
//}
//
//func (t *Topic) GetSubscriber(key string) *subscriber.Subscriber {
//
//	for _, value := range t.consumers {
//		if value.GetToken() == key {
//			return value
//		}
//	}
//
//	return nil
//}
//
//func (t *Topic) AddSubscriber(subscriber *subscriber.Subscriber) {
//	t.consumers = append(t.consumers, subscriber)
//}
//
//func (t *Topic) DeleteSubscriber(token string) {
//	for index, cons := range t.consumers {
//		if token == cons.GetToken() {
//			before := index - 1
//			if index == 0 {
//				before = 0
//			}
//
//			after := index + 1
//			if index == len(t.consumers)-1 {
//				after = len(t.consumers) - 1
//			}
//
//			t.consumers = append(t.consumers[:before], t.consumers[:after]...)
//			return
//		}
//	}
//}

func (t *Topic) PushToQueue(message queue.Message) {
	t.queue.PushMessage(message)
}

func (t *Topic) GetFromQueue(id int) (error, queue.Message) {
	return t.queue.GetMessage(id)
}

func (t *Topic) GetLastIndex() int {
	return t.queue.GetIndex()
}

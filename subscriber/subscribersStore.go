package subscriber

import "errors"

type SubscribersStore struct {
	subscribers []Subscriber
}

func NewSubscribersStore() SubscribersStore {
	return SubscribersStore{
		subscribers: []Subscriber{},
	}
}

func (cs *SubscribersStore) AddSubscriber(isAuto bool, topic string, index int) *Subscriber {
	c := NewSubscriber(isAuto)
	c.cursors[topic] = index
	cs.subscribers = append(cs.subscribers, c)

	return &c
}

func (cs *SubscribersStore) DeleteSubscriber(id string) error {
	for index, c := range cs.subscribers {
		if c.GetToken() == id {
			before := index - 1
			if index == 0 {
				before = 0
			}

			after := index + 1
			if index == len(cs.subscribers)-1 {
				after = len(cs.subscribers) - 1
			}

			cs.subscribers = append(cs.subscribers[:before], cs.subscribers[:after]...)
			return nil
		}
	}

	return errors.New("unknown subscriber id")
}

func (cs *SubscribersStore) GetSubscriber(id string) (error, *Subscriber) {
	for _, c := range cs.subscribers {
		if c.GetToken() == id {
			return nil, &c
		}
	}

	return errors.New("unknown subscriber"), nil
}

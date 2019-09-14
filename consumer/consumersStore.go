package consumer

import "errors"

type ConsumersStore struct {
	consumers []Consumer
}

func NewConsumersStore() ConsumersStore {
	cs := ConsumersStore{}

	cs.consumers = []Consumer{}

	return cs
}

func (cs *ConsumersStore) AddConsumer(isAuto bool) *Consumer {
	c := NewConsumer(isAuto)
	cs.consumers = append(cs.consumers, c)

	return &c
}

func (cs *ConsumersStore) DeleteConsumer(id string) error {
	for index, c := range cs.consumers {
		if c.GetName() == id {
			before := index - 1
			if index == 0 {
				before = 0
			}

			after := index + 1
			if index == len(cs.consumers)-1 {
				after = len(cs.consumers) - 1
			}

			cs.consumers = append(cs.consumers[:before], cs.consumers[:after]...)
			return nil
		}
	}

	return errors.New("unknown consumer id")
}

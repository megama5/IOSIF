package message_broker

import (
	"errors"
	"fmt"
)

//Subscriber
type Subscriber struct {
	AccessKey     string         `json:"accessKey"`
	MapTopicIndex map[string]int `json:"-"`
}

func (s *Subscriber) GetTopicIndex(topic string) (int, error) {
	index, ok := s.MapTopicIndex[topic]
	if !ok {
		return 0, errors.New("topic does not exists")
	}

	return index, nil
}

func (s *Subscriber) AddTopicIndex(topic string) error {
	_, ok := s.MapTopicIndex[topic]
	if !ok {
		return errors.New("topic already exists")
	}

	return nil
}

func NewSubscriber(list []string) *Subscriber {

	s := &Subscriber{
		AccessKey:     fmt.Sprintf("IO-%s-SUB", CreateRandString()),
		MapTopicIndex: map[string]int{},
	}

	for _, value := range list {
		if _, ok := s.MapTopicIndex[value]; ok {
			continue
		}
		s.MapTopicIndex[value] = -1
	}

	return s
}

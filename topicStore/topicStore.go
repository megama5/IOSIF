package topicStore

import (
	"IOSIF/utils"
	"errors"
)

type TopicStore struct {
	topicsStore map[string]*Topic
}

func NewTopicStore() TopicStore {
	return TopicStore{
		topicsStore: map[string]*Topic{},
	}
}

func (ts *TopicStore) GetTopic(topicName string) (*Topic, error) {
	topic, ok := ts.topicsStore[topicName]
	if ok {
		return topic, nil
	}

	return nil, errors.New(TopicDoesNotExists)
}

func (ts *TopicStore) CreateTopic(topicName string) error {
	_, err := ts.GetTopic(topicName)
	if err == nil {
		return errors.New(TopicAlreadyExists)
	}

	topic := NewTopic(topicName)
	ts.topicsStore[topicName] = &topic
	utils.LogAction(TopicCreated, topicName)
	return nil
}

func (ts *TopicStore) DeleteTopic(topicName string) error {

	if _, ok := ts.topicsStore[topicName]; !ok {
		return errors.New(TopicDoesNotExists)
	}

	delete(ts.topicsStore, topicName)
	if _, ok := ts.topicsStore[topicName]; !ok {
		utils.LogAction(TopicDeleted, topicName)
		return nil
	}

	return nil
}

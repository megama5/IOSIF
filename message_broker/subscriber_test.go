package message_broker

import (
	"reflect"
	"testing"
)

func TestNewSubscriber(t *testing.T) {
	topicsList := []string{"users"}
	subscriber := NewSubscriber(topicsList)

	testSubscriber := &Subscriber{
		AccessKey:     subscriber.AccessKey,
		MapTopicIndex: map[string]int{"users": -1},
	}

	if !reflect.DeepEqual(subscriber, testSubscriber) {
		t.Error("expected result", testSubscriber)
	}

}

func TestSubscriber_AddTopicIndex(t *testing.T) {
	topicsList := []string{"users"}
	subscriber := NewSubscriber(topicsList)
	testSubscriber := &Subscriber{
		AccessKey:     subscriber.AccessKey,
		MapTopicIndex: map[string]int{"users": -1},
	}

	err := subscriber.AddTopicIndex("users")
	if err != nil || !reflect.DeepEqual(subscriber, testSubscriber) {
		t.Error("expected result", testSubscriber)
	}
}

func TestSubscriber_GetTopicIndex(t *testing.T) {
	topicsList := []string{"users"}
	subscriber := NewSubscriber(topicsList)
	testSubscriber := &Subscriber{
		AccessKey:     subscriber.AccessKey,
		MapTopicIndex: map[string]int{"users": -1},
	}

	err := subscriber.AddTopicIndex("users")
	if err != nil || !reflect.DeepEqual(subscriber, testSubscriber) {
		t.Error("expected result", testSubscriber)
	}

	index, err := subscriber.GetTopicIndex("users")
	if err != nil || index != -1 {
		t.Error("expected result", -1)
	}
}

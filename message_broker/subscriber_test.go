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

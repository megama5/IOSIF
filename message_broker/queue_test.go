package message_broker

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()

	testQueue := &Queue{
		messages:         []Message{},
		lastMessageIndex: -1,
	}

	if !reflect.DeepEqual(queue, testQueue) {
		t.Error("expected result", queue)
	}
}

func TestQueue_AddMessage(t *testing.T) {
	queue := NewQueue()
	testMessage := Message{
		Index:     0,
		Value:     "aaa",
		Topic:     "users",
		CreatedAt: 0,
	}

	queue.AddMessage(testMessage)
	message, err := queue.GetMessage(0)
	if err != nil || !reflect.DeepEqual(message, testMessage) {
		t.Error("expected result", testMessage)
	}
}

func TestQueue_GetMessage(t *testing.T) {
	queue := NewQueue()
	testMessage := Message{
		Index:     0,
		Value:     "aaa",
		Topic:     "users",
		CreatedAt: 0,
	}

	queue.AddMessage(testMessage)
	message, err := queue.GetMessage(0)
	if err != nil || !reflect.DeepEqual(message, testMessage) {
		t.Error("expected result", testMessage)
	}
}

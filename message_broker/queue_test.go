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

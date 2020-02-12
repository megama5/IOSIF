package message_broker

import (
	"reflect"
	"testing"
)

func TestNewPublisher(t *testing.T) {
	queue := NewPublisher()

	testQueue := &Publisher{
		AccessKey: queue.AccessKey,
	}

	if !reflect.DeepEqual(queue, testQueue) {
		t.Error("expected result", testQueue)
	}

}

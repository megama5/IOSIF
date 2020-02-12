package message_broker

import (
	"reflect"
	"testing"
	"time"
)

func TestNewMessage(t *testing.T) {
	testMessage := &Message{
		Index:     0,
		Value:     "",
		Topic:     "",
		CreatedAt: time.Now().Unix(),
	}

	if !reflect.DeepEqual(NewMessage(0, ""), testMessage) {
		t.Error("message should be like", testMessage)
	}
}

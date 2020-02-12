package message_broker

import (
	"reflect"
	"testing"
)

func TestCreateRandString(t *testing.T) {
	if reflect.TypeOf(CreateRandString()).String() != "string" {
		t.Error("CreateRandString ")
	}
}

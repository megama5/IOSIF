package message_broker

import (
	"reflect"
	"testing"
)

func TestCreateRandString(t *testing.T) {
	tmp := make(map[string]interface{},10)
	for i:=0; i < 10; i++ {
		randomStr := CreateRandString()
		if _, v := tmp[randomStr]; ok {
			t.Error("CreateRandString - failed to generate random string")
		}
		
		tmp[randomStr] = nil
	}
}

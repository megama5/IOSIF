package config

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	conf, err := ReadConfig("../devops/config.yaml")
	if err != nil {
		t.Error(err)
	}

	testConfig := &Config{
		Server: Server{
			Port: "7000",
		},
		Manager: Manager{
			MessageChannelBufferSize: 50,
			WorkersCount:             5,
		},
		Topics: []string{"users"},
	}

	if !reflect.DeepEqual(conf, testConfig) {
		t.Error("expected result", testConfig)
	}
}

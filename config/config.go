package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Server struct {
	Port string `yaml:"port"`
}

type Manager struct {
	MessageChannelBufferSize int `yaml:"messageChannelBufferSize"`
	WorkersCount             int `yaml:"workersCount"`
}

type Config struct {
	Server  Server   `yaml:"server"`
	Manager Manager  `yaml:"manager"`
	Topics  []string `yaml:"topics"`
}

func ReadConfig(configPath string) (*Config, error) {

	var conf Config
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

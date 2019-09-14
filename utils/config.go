package utils

import (
	"fmt"
	"log"
)

type Config struct {
	Port              int    `yaml:"port"`
	Host              string `yaml:"host"`
	Path              string `yaml:"path"`
	MaxWorkers        int    `yaml:"max_workers"`
	ChannelBufferSize int    `yaml:"channel_buffer_size"`
}

func (c *Config) GetPath() string {
	path := ""

	if c.Port == 0 {
		log.Fatal("Port field is required")
	}

	path = path + c.Host
	if c.Host == "" {
		path = path + "localhost"
	}

	path = path + ":" + fmt.Sprint(c.Port)

	if c.Path != "" {
		path = path + c.Path
	}

	return path
}

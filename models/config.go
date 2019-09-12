package models

import (
	"fmt"
	"log"
)

type Config struct {
	Port              int
	Host              string
	Path              string
	ChannelBufferSize int
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

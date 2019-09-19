package utils

import (
	"fmt"
)

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

type ManagerConfig struct {
	MaxWorkers        int `yaml:"max_workers"`
	ChannelBufferSize int `yaml:"channel_buffer_size"`
}

type DataBaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SSLMode  bool   `yaml:"ssl_mode"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Manager  ManagerConfig  `yaml:"manager"`
	DataBase DataBaseConfig `yaml:"data_base"`
}

func (c *Config) GetPath() string {

	path := c.Server.Host
	if c.Server.Host == "" {
		path = "localhost"
	}

	return fmt.Sprint(path, ":", c.Server.Port, c.Server.Path)
}

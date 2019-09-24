package config

import (
	"fmt"
)

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

type ManagerConfig struct {
	MaxWorkers        int   `yaml:"maxWorkers"`
	ChannelBufferSize int   `yaml:"channelBufferSize"`
	MessageLifetime   int64 `yaml:"messageLifetime"`
}

type DataBaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
	SSLMode  bool   `yaml:"sslMode"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Manager  ManagerConfig  `yaml:"manager"`
	DataBase DataBaseConfig `yaml:"database"`
}

func (c *Config) GetPath() string {

	path := c.Server.Host
	if c.Server.Host == "" {
		path = "localhost"
	}

	return fmt.Sprint(path, ":", c.Server.Port, c.Server.Path)
}

package manager

import (
	"IOSIF/queue"
	"IOSIF/utils"
)

type Manager struct {
	messageChannel chan queue.Message
}

func (m *Manager) Constructor(conf utils.Config) {
	m.messageChannel = make(chan queue.Message, conf.ChannelBufferSize)
}

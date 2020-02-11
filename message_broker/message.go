package message_broker

import "time"

//Message
type Message struct {
	Index     int    `json:"index"`
	Value     string `json:"value"`
	Topic     string `json:"-"`
	CreatedAt int64  `json:"createdAt"`
}

func NewMessage(index int, value string) *Message {
	return &Message{
		Index:     index,
		Value:     value,
		CreatedAt: time.Now().Unix(),
	}
}

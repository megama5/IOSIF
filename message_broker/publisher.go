package message_broker

import "fmt"

//Publisher
type Publisher struct {
	AccessKey string `json:"accessKey"`
}

func NewPublisher() *Publisher {
	return &Publisher{
		AccessKey: fmt.Sprintf("IO-%s-PUB", CreateRandString()),
	}
}

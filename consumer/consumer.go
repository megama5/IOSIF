package consumer

import (
	uuid "github.com/satori/go.uuid"
)

type Consumer struct {
	Token      string `json:"token"`
	autoCursor bool
	cursors    map[string]int
}

func NewConsumer(autoCursor bool) Consumer {
	name, _ := uuid.NewV4()

	return Consumer{
		autoCursor: autoCursor,
		Token:      name.String(),
		cursors:    map[string]int{},
	}
}

func (c *Consumer) IsAutoCursorMove() bool {
	return c.autoCursor
}

func (c *Consumer) GetToken() string {
	return c.Token
}

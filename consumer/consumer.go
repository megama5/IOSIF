package consumer

import (
	"fmt"
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

func (c *Consumer) GetCursors() map[string]int {
	return c.cursors
}

func (c *Consumer) GetCursor(topic string) int {
	fmt.Println(c.cursors[topic])
	return c.cursors[topic]
}

func (c *Consumer) SetCursor(topic string, messageIndex int) {
	c.cursors[topic] = messageIndex
}

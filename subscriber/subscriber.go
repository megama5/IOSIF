package subscriber

import (
	uuid "github.com/satori/go.uuid"
)

type Subscriber struct {
	Token      string         `json:"token"`
	autoCursor bool           `json:"-"`
	cursors    map[string]int `json:"-"`
}

func NewSubscriber(autoCursor bool) Subscriber {
	name, _ := uuid.NewV4()

	return Subscriber{
		autoCursor: autoCursor,
		Token:      name.String(),
		cursors:    map[string]int{},
	}
}

func (c Subscriber) IsAutoCursorMove() bool {
	return c.autoCursor
}

func (c Subscriber) GetToken() string {
	return c.Token
}

func (c Subscriber) GetCursor(topic string) int {
	return c.cursors[topic]
}

func (c *Subscriber) SetCursor(topic string, messageIndex int) {
	c.cursors[topic] = messageIndex
}

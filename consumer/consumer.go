package consumer

import (
	uuid "github.com/satori/go.uuid"
)

type Consumer struct {
	Name       string `json:"name"`
	autoCursor bool
	cursors    map[string]int
}

func NewConsumer(autoCursor bool) Consumer {
	name, _ := uuid.NewV4()

	return Consumer{
		autoCursor: autoCursor,
		Name:       name.String(),
		cursors:    map[string]int{},
	}
}

func (c *Consumer) IsAutoCursorMove() bool {
	return c.autoCursor
}

func (c *Consumer) GetName() string {
	return c.Name
}

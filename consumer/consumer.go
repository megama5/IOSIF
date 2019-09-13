package consumer

type Consumer struct {
	autoCursorMove bool
	name           string
	index          int
}

func (c *Consumer) isAutoCursorMove() bool {
	return c.autoCursorMove
}

func (c *Consumer) getName() string {
	return c.name
}

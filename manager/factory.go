package manager

import (
	"IOSIF/queue"
	"time"
)

type Factory struct {
	maxWorkers int
	workers    int
	bufferSize *int
	commands   chan int
	channel    *chan queue.Message
	handler    func(message *queue.Message)
}

func (f *Factory) Supervisor() {
	if f.workers == 0 {
		go f.Worker()
		f.workers = +1
	}

	for {
		if *f.bufferSize/2 < len(*f.channel) {

		} else {
			if f.workers > 1 {
				f.commands <- 0
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (f *Factory) Worker() {
	for {
		select {
		case <-f.commands:
			break
		case message := <-*f.channel:
			f.handler(&message)
		}
	}
}

func (f *Factory) RegisterHandler(handler func(message *queue.Message)) {
	f.handler = handler
}

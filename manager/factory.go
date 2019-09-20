package manager

import (
	"IOSIF/message"
	"IOSIF/utils"
	"time"
)

type Factory struct {
	maxWorkers int
	workers    int
	bufferSize *int
	commands   chan int
	channel    *chan message.Message
	handler    func(message *message.Message)
}

func (f *Factory) Supervisor() {
	if f.workers == 0 {
		go f.Worker()
		f.workers = +1
	}

	for {
		if *f.bufferSize/2 < len(*f.channel) {

		} else if f.workers == 0 {
			utils.Log("supervise killed all workers")
			return
		} else {
			if f.workers > 1 {
				utils.Log(FactoryKillWorker)
				f.commands <- 0
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (f *Factory) Worker() {
	utils.Log(FactoryNewWorker)
	for {
		select {
		case <-f.commands:
			return
		case message := <-*f.channel:
			f.handler(&message)
		}
	}
}

func (f *Factory) RegisterHandler(handler func(message *message.Message)) {
	f.handler = handler
}

func (f *Factory) StopFactory() {
	for i := f.workers; i > 0; i-- {
		f.commands <- 0
		utils.Log(FactoryKillWorker)
	}
	f.workers = 0
}

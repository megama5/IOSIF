package queue

import "time"

type Message struct {
	TraceId   string `json:"trace_id"`
	Index     int
	Topic     string `json:"topic"`
	TimeStamp string `json:"time_stamp"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func (m *Message) signTimeStamp() {
	m.TimeStamp = time.Now().String()
}

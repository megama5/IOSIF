package models

import "time"

type Message struct {
	TraceId   string `json:"trace_id"`
	Topic     string `json:"topic"`
	TimeStamp string `json:"time_stamp"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func (m *Message) signTimeStamp() {
	m.TimeStamp = time.Now().String()
}

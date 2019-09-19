package queue

import (
	"encoding/json"
	"time"
)

type Message struct {
	TraceId   string `json:"trace_id"`
	Index     int
	Topic     string `json:"topic"`
	TimeStamp string `json:"time_stamp"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func (m *Message) SignTimeStamp() {
	m.TimeStamp = time.Now().Format(time.RFC3339)
}

func (m *Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

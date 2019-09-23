package message

import (
	"encoding/json"
	"time"
)

type Message struct {
	TraceId   string `json:"traceId" sql:"traceId"`
	Index     int    `json:"index" sql:"index"`
	Topic     string `json:"topic" sql:"topic"`
	TimeStamp string `json:"time_stamp" sql:"time_stamp"`
	Key       string `json:"key" sql:"key"`
	Value     string `json:"value" sql:"value"`
}

func (m *Message) SignTimeStamp() {
	m.TimeStamp = time.Now().Format(time.RFC3339)
}

func (m Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

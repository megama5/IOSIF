package core

import (
	"IOSIF/consumer"
	"IOSIF/queue"
	"encoding/json"
	"net/http"
)

//<---------------Topic Controllers--------------->
func Queue(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMessage(w, r)
	case http.MethodPut:
		postMessage(w, r)
	case http.MethodPost:
		initTopic(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func getMessage(w http.ResponseWriter, r *http.Request) {

}

func postMessage(w http.ResponseWriter, r *http.Request) {
	var m queue.Message

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}

	m.Topic = GetQuery("topic", r)
	Manager.PushToChannel(m)

	w.WriteHeader(http.StatusCreated)
}

func initTopic(w http.ResponseWriter, r *http.Request) {

}

//<---------------Subscriber Controllers--------------->
func Subscribe(w http.ResponseWriter, r *http.Request) {
	var cons *consumer.Consumer

	topic := GetQuery("topic", r)
	isAuto := GetQuery("isAutoCounter", r)
	if isAuto == "true" {
		cons = ConsumersStore.AddConsumer(true, topic)
	} else {
		cons = ConsumersStore.AddConsumer(false, topic)
	}

	response, _ := json.Marshal(*cons)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(response)
}

package core

import (
	"IOSIF/message_broker"
	"encoding/json"
	"net/http"
)

//Subscriber
func AddSubscriber(w http.ResponseWriter, r *http.Request) {
	var topicsList []string

	err := json.NewDecoder(r.Body).Decode(&topicsList)

	subscriber := broker.AddSubscriber(topicsList)

	response, err := json.Marshal(subscriber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(response)
}

//Publisher
func AddPublisher(w http.ResponseWriter, r *http.Request) {
	publisher := broker.AddPublisher()
	response, err := json.Marshal(publisher)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(response)
}

//Messages
func EmitMessage(w http.ResponseWriter, r *http.Request) {
	accessKey := r.Header.Get(ACCESS_KEY)
	topic := r.URL.Query().Get(TOPIC)
	if accessKey == "" || topic == "" || !broker.CheckPublisher(accessKey) || !broker.CheckTopic(topic) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var message message_broker.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message.Topic = topic

	broker.PushMessage(message)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	accessKey := r.Header.Get(ACCESS_KEY)
	topic := r.URL.Query().Get(TOPIC)
	if accessKey == "" || topic == "" || !broker.CheckSubscriber(accessKey) || !broker.CheckTopic(topic) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	message, err := broker.GetMessage(topic, accessKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(response)
}

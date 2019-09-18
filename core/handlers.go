package core

import (
	"IOSIF/consumer"
	"IOSIF/queue"
	"IOSIF/utils"
	"encoding/json"
	"fmt"
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
		createTopic(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	topicName := GetQuery("topic", r)
	consumerToken := GetHeader(TOKEN_HEADER, r)

	topic, ok := TopicStore.GetTopic(topicName)
	if !ok || consumerToken == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cons := topic.GetConsumer(consumerToken)
	err, m := topic.GetFromQueue(cons.Token)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cons.SetCursor(topicName, m.Index)
	response, err := json.Marshal(&m)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(response)
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

func createTopic(w http.ResponseWriter, r *http.Request) {
	topicName := GetQuery("topic", r)

	if _, ok := TopicStore.GetTopic(topicName); ok {
		w.WriteHeader(http.StatusConflict)
		return
	}

	TopicStore.CreateTopic(topicName)
	utils.Log("new topic creating " + topicName)
	w.WriteHeader(http.StatusCreated)
}

//<---------------Subscriber Controllers--------------->
func Subscribe(w http.ResponseWriter, r *http.Request) {
	var cons *consumer.Consumer

	topicName := GetQuery("topic", r)
	topic, ok := TopicStore.GetTopic(topicName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isAuto := GetQuery("autoCounter", r)
	if isAuto == "true" {
		cons = ConsumersStore.AddConsumer(true, topicName)
	} else {
		cons = ConsumersStore.AddConsumer(false, topicName)
	}

	topic.AddConsumer(cons)
	utils.Log("user subscribe on topicName" + topicName)
	response, _ := json.Marshal(*cons)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(response)
}

func UnSubscribe(w http.ResponseWriter, r *http.Request) {
	token := GetHeader("x-sub-token", r)
	cons := ConsumersStore.GetConsumer(token)
	if cons == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for topic, _ := range cons.GetCursors() {
		t, ok := TopicStore.GetTopic(topic)
		if ok {
			t.DeleteConsumer(token)
		}
	}

	err := ConsumersStore.DeleteConsumer(token)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

package core

import (
	"IOSIF/queue"
	"IOSIF/subscriber"
	"IOSIF/utils"
	"encoding/json"
	"net/http"
)

//<---------------Topic Controllers--------------->
func Queue(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
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
	defer Manager.Stop()
	subscriberId := GetHeader(TOKEN_HEADER, r)
	err, sub := SubscibersStore.GetSubscriber(subscriberId)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	topicId := GetQuery("topic", r)
	cursor := sub.GetCursor(topicId)
	topic, ok := TopicStore.GetTopic(topicId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err, message := topic.GetFromQueue(cursor + 1)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := message.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sub.SetCursor(message.Topic, message.Index)
	_, _ = w.Write(response)

}

func postMessage(w http.ResponseWriter, r *http.Request) {
	defer Manager.Stop()
	var m queue.Message

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}
	m.SignTimeStamp()
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
	var cons *subscriber.Subscriber

	topicName := GetQuery("topic", r)
	topic, ok := TopicStore.GetTopic(topicName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isAuto := GetQuery("autoCounter", r) == "true"
	cons = SubscibersStore.AddSubscriber(isAuto, topicName, topic.GetLastIndex())

	utils.Log("user subscribe on topicName " + topicName)
	response, _ := json.Marshal(*cons)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(response)
}

func UnSubscribe(w http.ResponseWriter, r *http.Request) {
	token := GetHeader(TOKEN_HEADER, r)

	err := SubscibersStore.DeleteSubscriber(token)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

package bootstrap

import (
	"IOSIF/message"
	"IOSIF/subscriber"
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

	subscriberId := GetHeader(TokenHeader, r)
	err, sub := SubscribersStore.GetSubscriber(subscriberId)
	if err != nil {
		SendError(err, http.StatusNotFound, w)
		return
	}

	topicId := GetQuery("topic", r)
	cursor := sub.GetCursor(topicId)
	topic, err := TopicStore.GetTopic(topicId)
	if err != nil {
		SendError(err, http.StatusNotFound, w)
		return
	}

	err, mes := topic.GetFromQueue(cursor + 1)
	if err != nil {
		SendError(err, http.StatusNotFound, w)
		return
	}

	response, err := mes.ToJSON()
	if err != nil {
		SendError(err, http.StatusNotFound, w)
		return
	}

	sub.SetCursor(mes.Topic, mes.Index)
	_, _ = w.Write(response)

}

func postMessage(w http.ResponseWriter, r *http.Request) {
	defer Manager.Stop()
	var m message.Message

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		SendError(err, http.StatusNotFound, w)
		return
	}

	m.SignTimeStamp()
	m.Topic = GetQuery("topic", r)

	_, err := TopicStore.GetTopic(m.Topic)
	if err != nil {
		SendError(err, http.StatusNotFound, w)
		return
	}

	Manager.PushToChannel(m)

	w.WriteHeader(http.StatusCreated)
}

func createTopic(w http.ResponseWriter, r *http.Request) {
	topicName := GetQuery("topic", r)

	if err := TopicStore.CreateTopic(topicName); err != nil {
		SendError(err, http.StatusConflict, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//<---------------Subscriber Controllers--------------->
func Subscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var cons *subscriber.Subscriber
	topicName := GetQuery("topic", r)

	topic, err := TopicStore.GetTopic(topicName)
	if err != nil {
		SendError(err, http.StatusNotFound, w)
		return
	}

	isAuto := GetQuery("autoCounter", r) == "true"
	cons = SubscribersStore.AddSubscriber(isAuto, topicName, topic.GetLastIndex())

	response, _ := json.Marshal(*cons)

	_, _ = w.Write(response)
}

func UnSubscribe(w http.ResponseWriter, r *http.Request) {
	token := GetHeader(TokenHeader, r)

	err := SubscribersStore.DeleteSubscriber(token)
	if err != nil {
		SendError(err, http.StatusConflict, w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

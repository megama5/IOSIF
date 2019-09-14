package core

import (
	"IOSIF/queue"
	"encoding/json"
	"net/http"
)

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

	json.NewDecoder(r.Body).Decode(&m)

	Manager.PushToChannel(m)

	w.WriteHeader(http.StatusCreated)
}

func initTopic(w http.ResponseWriter, r *http.Request) {

}

func Subscribe(w http.ResponseWriter, r *http.Request) {

}

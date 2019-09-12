package controllers

import (
	"encoding/json"
	"fmt"
	"messege-queue/models"
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

	message := Manager.GetMessage(r.URL.Query().Get("name"), r.URL.Query().Get("subscriber"))

	JSON, err := json.Marshal(message)
	if err != nil {

	}

	w.Write(JSON)

}

func postMessage(w http.ResponseWriter, r *http.Request) {

	var message models.Message

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
		return
	}

	Manager.PushMessage(message)
}

func initTopic(w http.ResponseWriter, r *http.Request) {
	if err := Manager.CreateTopic(r.URL.Query().Get("name")); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	err := Manager.Subscribe(r.URL.Query().Get("name"), r.URL.Query().Get("subscriber"))
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprint(err)))
	}

	w.WriteHeader(http.StatusAccepted)
}

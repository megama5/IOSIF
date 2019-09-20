package bootstrap

import (
	"IOSIF/utils"
	"net/http"
)

func GetQuery(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}

func GetHeader(key string, r *http.Request) string {
	return r.Header.Get(key)
}

func SendError(err error, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	response, _ := utils.NewError(err.Error(), status).ToJSON()
	_, _ = w.Write(response)
}

package core

import "net/http"

func GetQuery(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}

package utils

import "net/http"

func WithErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	errMessage := ErrorLog(err)
	LogMessage(errMessage)

	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(errMessage))
}

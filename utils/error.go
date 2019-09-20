package utils

import (
	"encoding/json"
)

type Error struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewError(message string, statusCode int) Error {
	return Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e Error) ToJSON() ([]byte, error) {
	JSON, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	Log(e.Message)

	return JSON, nil
}

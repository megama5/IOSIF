package utils

import "encoding/json"

type Error struct {
	Message string `json:"message"`
}

func NewError(message string) Error {
	return Error{
		Message: message,
	}
}

func (e Error) toJSON() ([]byte, error) {
	JSON, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return JSON, nil
}

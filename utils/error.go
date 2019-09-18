package utils

import "encoding/json"

type Error struct {
	message string `json:"message"`
}

func NewError(message string) Error {
	return Error{
		message: message,
	}
}

func (e *Error) toJSON() ([]byte, error) {
	JSON, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return JSON, nil
}

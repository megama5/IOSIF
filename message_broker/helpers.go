package message_broker

import "math/rand"

//Helpers
func CreateRandString() string {
	const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)

	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(b)
}

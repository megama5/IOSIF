package repositories

import "IOSIF/message"

type Driver interface {
	Connect() error
	Insert(message message.Message) error
	Delete(message message.Message) error
}

type Repository struct {
	currentDb string
	drivers   map[string]*Driver
}

func NewRepository() Repository {
	repo := Repository{}

	return
}

func (r *Repository) AddOne(message message.Message) {

}

func (r *Repository) DeleteOne(message message.Message) {

}

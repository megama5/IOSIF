package repositories

import (
	"IOSIF/config"
	"IOSIF/message"
	"errors"
	"fmt"
)

const (
	PostgresDB = "postgres"
)

var availableDrivers = []string{PostgresDB}

type Driver interface {
	Connect() error
	Insert(message message.Message) error
	Delete(message message.Message) error
}

type Repository struct {
	currentDriver string
	drivers       map[string]Driver
}

func NewRepository(config *config.Config) (Repository, error) {

	for _, v := range availableDrivers {
		if v == config.DataBase.Driver {
			return Repository{
				currentDriver: config.DataBase.Driver,
				drivers: map[string]Driver{
					PostgresDB: NewPostgres(config),
				},
			}, nil
		}
	}

	return Repository{}, errors.New(fmt.Sprintf("unknown driver -> %s", config.DataBase.Driver))
}

func (r *Repository) Connect() error {
	return r.drivers[r.currentDriver].Connect()
}

func (r *Repository) Insert(message message.Message) error {
	return r.drivers[r.currentDriver].Insert(message)
}

func (r *Repository) Delete(message message.Message) error {
	return r.drivers[r.currentDriver].Delete(message)
}

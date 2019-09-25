package repositories

import (
	"IOSIF/config"
	"IOSIF/message"
	"IOSIF/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type postgres struct {
	DB           *sql.DB
	user         string
	password     string
	dbName       string
	messageTable string
	sslmode      bool
}

func NewPostgres(conf *config.Config) *postgres {
	p := postgres{
		user:         conf.DataBase.User,
		password:     conf.DataBase.Password,
		dbName:       conf.DataBase.DBName,
		sslmode:      conf.DataBase.SSLMode,
		messageTable: "messages",
	}

	return &p
}

func (p *postgres) Connect() error {

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s ", p.user, p.password, p.dbName)
	connStr = connStr + "sslmode=disable"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		defer conn.Close()
		return err
	}
	err = conn.Ping()
	if err != nil {
		defer conn.Close()
		return err
	}
	p.DB = conn

	err = p.createTable()
	if err != nil {
		defer conn.Close()
		return err
	}

	return nil
}

func (p postgres) createTable() error {
	table := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (TraceId text, Index int, Topic text, TimeStamp text, Key text, Value text)", p.messageTable)
	_, err := p.DB.Exec(table)
	if err != nil {
		return err
	}

	return nil
}

func (p postgres) Insert(m message.Message) error {
	value := fmt.Sprintf("VALUES ('%s', %d, '%s', '%s', '%s', '%s')", m.TraceId, m.Index, m.Topic, m.TimeStamp, m.Key, m.Value)
	query := fmt.Sprintf("INSERT INTO %s %s", p.messageTable, value)

	_, err := p.DB.Exec(query)
	if err != nil {
		utils.Log(err)
		return err
	}

	return nil
}

func (p postgres) Delete(m message.Message) error {
	value := fmt.Sprintf("DELETE FROM %s WHERE Topic = '%s' AND Index = %d", p.messageTable, m.Topic, m.Index)

	_, err := p.DB.Exec(value)
	if err != nil {
		return err
	}

	return nil
}

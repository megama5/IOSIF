package postgres

import (
	"IOSIF/config"
	"IOSIF/message"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Postgres struct {
	DB       *sql.DB
	user     string
	password string
	dbName   string
	sslmode  bool
}

func NewPostgres(conf *config.Config) Postgres {
	p := Postgres{
		user:     conf.DataBase.User,
		password: conf.DataBase.Password,
		dbName:   conf.DataBase.DBName,
		sslmode:  conf.DataBase.SSLMode,
	}

	return p
}

func (p *Postgres) Connect() {
	connStr := fmt.Sprintf("user=%s ", p.user)
	connStr = connStr + fmt.Sprintf("password=%s ", p.password)
	connStr = connStr + fmt.Sprintf("dbname=%s", p.dbName)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		defer conn.Close()
		log.Fatal(err)
	}

	p.DB = conn
}

func (p *Postgres) AddLog(message message.Message) {

}

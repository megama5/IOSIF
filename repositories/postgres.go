package repositories

import (
	"IOSIF/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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

func (p *Postgres) Connect() error {
	connStr := fmt.Sprintf("user=%s ", p.user)
	connStr = connStr + fmt.Sprintf("password=%s ", p.password)
	connStr = connStr + fmt.Sprintf("dbname=%s", p.dbName)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		defer conn.Close()
		return err
	}

	p.DB = conn
	return nil
}

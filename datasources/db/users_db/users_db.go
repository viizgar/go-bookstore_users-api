package users_db

import (
	"log"

	"github.com/jackc/pgx"
)

const (
	pql_users_username = "go"
	pql_users_password = "go"
	pql_users_host     = "127.0.0.1"
	pql_users_schema   = "go"
	pql_users_db       = "go"
)

var (
	Client *pgx.Conn

	username = pql_users_username
	password = pql_users_password
	host     = pql_users_host
	schema   = pql_users_schema
)

func init() {
	var err error
	Client, err = pgx.Connect(pgx.ConnConfig{Host: host, Port: 5432, Database: schema, User: username, Password: password})
	if err != nil {
		panic(err)
	}

	log.Println("database successfully connected")
}

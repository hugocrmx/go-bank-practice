package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB *sqlx.DB
)

func init() {
	var err error
	DB, err = sqlx.Connect("postgres", "host=localhost user=postgres password=admin port=5432 dbname=go_practice sslmode=disable")
	if err != nil {
		log.Fatal("Couldn't connect to DB", err)
	}
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
}

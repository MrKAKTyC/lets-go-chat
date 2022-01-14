package repository

import (
	"database/sql"
	"github.com/MrKAKTyC/lets-go-chat/pkg/config"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func GetDBConnection(config config.DBConfig) *sql.DB {
	if db == nil {
		db, err = sql.Open("postgres", config.URL)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("connection to %s established", config.URL)
	}
	return db
}

package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func GetDBConnection(dbUrl string) (*sql.DB, error) {
	if db == nil {
		db, err = sql.Open("postgres", dbUrl)
		if err != nil {
			log.Fatal(err)
			db = nil
			return nil, err
		}
		log.Printf("connection to %s established", dbUrl)
	}
	return db, nil
}

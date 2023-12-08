package configs

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var connected bool = false
var db *sql.DB

func connect() {
	var err error
	db, err = sql.Open("sqlite3", "db/db")
	if err != nil {
		log.Fatal(err)
	}
	connected = true

	if err != nil {
		log.Fatal(err)
	}

}

func GetDB() *sql.DB {

	if !connected {
		connect()
		return db
	} else {
		return db
	}
}

func Close() {
	if connected {
		db.Close()

	}
}

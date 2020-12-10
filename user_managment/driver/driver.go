package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateConnection() *sql.DB {
	if db != nil {
		return db
	}
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err := sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)
	return db
}

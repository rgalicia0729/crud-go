package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// PostgresConnect Connect to the database
func PostgresConnect(pgUri string) {
	once.Do(func() {
		var err error

		db, err = sql.Open("postgres", pgUri)
		if err != nil {
			log.Fatalf("Can't open db: %v\n", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v\n", err)
		}
	})
}

// PostgresPool Returns a database connection
func PostgresPool() *sql.DB {
	return db
}

package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName = "database.sqlite"
)

func NewDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", dbName)
}

func RunMigration(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender TEXT,
		recipient TEXT,
		message TEXT
	)
`)
	if err != nil {
		log.Fatal(err)
	}
}

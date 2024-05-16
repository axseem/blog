package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func Migrate(db *sql.DB) error {
	const query = `CREATE TABLE IF NOT EXISTS article (
	id TEXT PRIMARY KEY,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS author (
	username TEXT PRIMARY KEY NOT NULL,
	password BLOB NOT NULL
);`

	_, err := db.Exec(query)
	return err
}

func deferErr(fn func() error, err *error) {
	deferErr := fn()
	if *err != nil {
		if deferErr != nil {
			log.Printf("ERROR: %v\n", *err)
		}
		return
	}
	*err = deferErr
}

func MustDefer(fn func() error) {
	if err := fn(); err != nil {
		log.Fatal(err)
	}
}

func MustOpen(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

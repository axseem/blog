package storage

import (
	"database/sql"
	"log"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func Migrate(db *sql.DB) error {
	const query = `CREATE TABLE article (
	id TEXT PRIMARY KEY,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
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

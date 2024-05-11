package main

import (
	"blomple/handler"
	"blomple/storage"
	"blomple/view"
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "dev.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := storage.Migrate(db); err != nil {
		log.Fatal(err)
	}

	h := handler.New(storage.New(db), view.NewTemplates())

	root := http.NewServeMux()
	root.HandleFunc("GET /", h.HomePage)
	root.HandleFunc("POST /", h.ArticleCreate)

	log.Fatal(http.ListenAndServe(":8080", root))
}

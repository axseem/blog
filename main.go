package main

import (
	"blomple/database"
	"blomple/handler"
	"blomple/view"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	db := database.MustOpen("dev.db")
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	h := handler.New(database.NewStorage(db), view.NewTemplates())

	root := http.NewServeMux()
	root.HandleFunc("GET /", h.HomePage)
	root.HandleFunc("GET /{id}", h.ArticlePage)
	root.HandleFunc("POST /", h.ArticleCreate)

	log.Fatal(http.ListenAndServe(":8080", root))
}

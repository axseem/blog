package main

import (
	"blomple/database"
	"blomple/handler"
	"blomple/view"
	"log"
	"net/http"
)

func main() {
	db := database.MustOpen("dev.db")
	defer database.MustDefer(db.Close)

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

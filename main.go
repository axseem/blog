package main

import (
	"blog/database"
	"blog/handler"
	"blog/middleware"
	"blog/view"
	"log"
	"net/http"
)

func main() {
	db := database.MustOpen("dev.db")
	defer database.MustDefer(db.Close)

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	s := database.NewStorage(db)
	h := handler.New(s, view.New())
	m := middleware.New(s)

	root := http.NewServeMux()
	root.HandleFunc("GET /", h.HomePage)
	root.HandleFunc("GET /{id}", h.ArticlePage)
	root.HandleFunc("POST /", m.Authorized(h.ArticleCreate))

	log.Fatal(http.ListenAndServe(":8080", root))
}

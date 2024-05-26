package main

import (
	"log"
	"net/http"

	"github.com/axseem/website/article"
	"github.com/axseem/website/handler"
	"github.com/axseem/website/static"
	"github.com/axseem/website/view"
)

func main() {
	staticFS := static.Static()

	articles, err := article.ExtractFromFS(&staticFS)
	if err != nil {
		panic(err)
	}

	indexPage, err := view.GenerateIndexPage(&articles)
	if err != nil {
		panic(err)
	}

	articlePages, err := view.GenerateArticles(&articles)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /", http.FileServerFS(&staticFS))
	mux.HandleFunc("GET /{$}", handler.Static(indexPage))
	mux.HandleFunc("GET /blog/{id}/{$}", handler.Articles(&articlePages))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

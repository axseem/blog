package main

import (
	"blog/article"
	"blog/handler"
	"blog/static"
	"blog/view"
	"log"
	"net/http"
)

func main() {
	articlesFS := static.Articles()
	v := view.New()

	articles, err := article.ExtractFromFS(&articlesFS)
	if err != nil {
		panic(err)
	}

	indexPage, err := v.GenerateIndexPage(&articles)
	if err != nil {
		panic(err)
	}

	articlePages, err := v.GenerateArticles(&articles)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /", http.FileServerFS(static.Static()))
	mux.HandleFunc("GET /{$}", handler.Static(indexPage))
	mux.HandleFunc("GET /{id}", handler.ArticlePage(&articlePages))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

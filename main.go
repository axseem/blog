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
	// assetsFS := static.Assets()
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

	root := http.NewServeMux()
	root.HandleFunc("GET /", handler.Static(indexPage))
	root.HandleFunc("GET /{id}", handler.ArticlePage(&articlePages))

	log.Fatal(http.ListenAndServe(":8080", root))
}

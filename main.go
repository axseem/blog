package main

import (
	"log"
	"net/http"
	"os"

	"github.com/axseem/website/article"
	"github.com/axseem/website/view"
)

func main() {
	staticFS := os.DirFS("./static")

	articles, err := article.ExtractFromFS(staticFS)
	if err != nil {
		panic(err)
	}

	err = view.GenerateIndexPage(&articles)
	if err != nil {
		panic(err)
	}

	err = view.GenerateArticles(&articles)
	if err != nil {
		panic(err)
	}

	http.Handle("GET /", http.FileServerFS(staticFS))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package handler

import (
	"log"
	"net/http"
)

func (s *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	articles, err := s.storage.ArticleList()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to fetch articles", 500)
		return
	}

	if err = s.view.Index.Execute(w, articles); err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to parse articles", 500)
		return
	}
}

func (s *Handler) ArticlePage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	article, err := s.storage.ArticleFind(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}

	if err = s.view.Article.Execute(w, article); err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to parse article", 500)
		return
	}
}

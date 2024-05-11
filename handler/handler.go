package handler

import (
	"blomple/model"
	"blomple/storage"
	"blomple/view"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	storage storage.Storage
	view    view.Templates
}

func New(storage *storage.Storage, view *view.Templates) *Handler {
	return &Handler{
		storage: *storage,
		view:    *view,
	}
}

func (s *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	articles, err := s.storage.ArticleList()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to fetch articles", 500)
		return
	}

	err = s.view.Index.Execute(w, articles)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to parse articles", 500)
		return
	}
}

func (s *Handler) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	if len(strings.TrimSpace(article.Title)) == 0 {
		http.Error(w, "Title can not be empty", 400)
	}

	if len(strings.TrimSpace(article.Content)) == 0 {
		http.Error(w, "Article can not be empty", 400)
	}

	if err := s.storage.ArticleCreate(article); err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to create article", 500)
		return
	}
}

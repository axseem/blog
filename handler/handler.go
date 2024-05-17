package handler

import (
	"blog/database"
	"blog/model"
	"blog/view"
	"log"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	storage database.Storage
	view    view.View
}

func New(storage *database.Storage, view *view.View) *Handler {
	return &Handler{
		storage: *storage,
		view:    *view,
	}
}

func (s *Handler) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	if len(strings.TrimSpace(article.Title)) == 0 {
		http.Error(w, "Title can not be empty", http.StatusBadRequest)
		return
	}

	if len(strings.TrimSpace(article.Content)) == 0 {
		http.Error(w, "Article can not be empty", http.StatusBadRequest)
		return
	}

	if err := s.storage.ArticleCreate(article); err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to create article", http.StatusInternalServerError)
		return
	}

	time.Now().Format(time.RFC3339)
}

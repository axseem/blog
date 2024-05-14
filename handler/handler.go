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
	view    view.Templates
}

func New(storage *database.Storage, view *view.Templates) *Handler {
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

func (s *Handler) ArticleCreate(w http.ResponseWriter, r *http.Request) {
	article := &model.Article{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	if len(strings.TrimSpace(article.Title)) == 0 {
		http.Error(w, "Title can not be empty", 400)
		return
	}

	if len(strings.TrimSpace(article.Content)) == 0 {
		http.Error(w, "Article can not be empty", 400)
		return
	}

	if err := s.storage.ArticleCreate(article); err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to create article", 500)
		return
	}

	time.Now().Format(time.RFC3339)
}

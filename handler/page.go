package handler

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gomarkdown/markdown"
)

func (s *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	articles, err := s.storage.ArticleList()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to fetch articles", http.StatusInternalServerError)
		return
	}

	if err = s.view.Template.ExecuteTemplate(w, "index.html", articles); err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to parse articles", http.StatusInternalServerError)
		return
	}
}

type PageData struct {
	ID        string
	Title     string
	Content   template.HTML
	CreatedAt time.Time
}

func (s *Handler) ArticlePage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	article, err := s.storage.ArticleFind(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	pageData := PageData{
		ID:        article.ID,
		Title:     article.Title,
		Content:   template.HTML(markdown.Render(s.view.Parser.Parse([]byte(article.Content)), s.view.Renderer)),
		CreatedAt: article.CreatedAt,
	}

	if err = s.view.Template.ExecuteTemplate(w, "article.html", pageData); err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to parse article", 500)
		return
	}
}

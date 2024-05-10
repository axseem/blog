package handler

import (
	"blomple/storage"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	storage storage.Storage
}

func New(storage storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (s *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	articles, err := s.storage.ArticleList()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		http.Error(w, "Failed to fetch articles", 500)
		return
	}
	io.WriteString(w, strconv.Itoa(len(articles)))
}

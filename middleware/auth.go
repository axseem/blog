package middleware

import (
	"blog/database"
	"net/http"
)

type Middleware struct {
	storage *database.Storage
}

func New(storage *database.Storage) *Middleware {
	return &Middleware{
		storage: storage,
	}
}

func (m *Middleware) Authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			if err := m.storage.ValidAuthor(username, password); err == nil {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

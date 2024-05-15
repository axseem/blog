package middleware

import (
	"blog/database"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Middleware struct {
	storage database.Storage
}

func New() {

}

func (m *Middleware) Authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 24)
			if err == nil {
				if m.storage.ValidAuthor(username, passwordHash) {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

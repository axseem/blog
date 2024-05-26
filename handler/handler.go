package handler

import (
	"net/http"

	_ "embed"
)

func Static(page []byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(page)
		if err != nil {
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
			panic(err)
		}
	}
}

func Articles(articles *map[string][]byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		article, ok := (*articles)[id]
		if !ok {
			code := http.StatusNotFound
			http.Error(w, http.StatusText(code), code)
			return
		}

		_, err := w.Write(article)
		if err != nil {
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
			panic(err)
		}
	}
}

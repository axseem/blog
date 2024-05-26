package view

import (
	"bytes"
	_ "embed"
	"html/template"

	"github.com/axseem/website/article"
)

//go:embed index.html
var indexTemplate string

func GenerateIndexPage(articles *[]article.Article) ([]byte, error) {
	var buf bytes.Buffer

	t := template.Must(template.New("").Parse(indexTemplate))
	if err := t.ExecuteTemplate(&buf, "", articles); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//go:embed article.html
var articleTemplate string

func GenerateArticles(articles *[]article.Article) (map[string][]byte, error) {
	pages := make(map[string][]byte)

	t := template.Must(template.New("").Parse(articleTemplate))
	for _, article := range *articles {
		var buf bytes.Buffer
		if err := t.ExecuteTemplate(&buf, "", article); err != nil {
			return nil, err
		}
		pages[article.ID()] = buf.Bytes()
	}

	return pages, nil
}

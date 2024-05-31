package view

import (
	"bytes"
	_ "embed"
	"html/template"
	"os"
	"path/filepath"

	"github.com/axseem/website/article"
)

//go:embed index.html
var indexTemplate string

func GenerateIndexPage(articles *[]article.Article) error {
	var buf bytes.Buffer

	t := template.Must(template.New("").Parse(indexTemplate))
	if err := t.ExecuteTemplate(&buf, "", articles); err != nil {
		return err
	}

	path := filepath.Join("static", "index.html")
	if err := os.WriteFile(path, buf.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

//go:embed article.html
var articleTemplate string

func GenerateArticles(articles *[]article.Article) error {
	t := template.Must(template.New("").Parse(articleTemplate))
	for _, article := range *articles {
		var buf bytes.Buffer
		if err := t.ExecuteTemplate(&buf, "", article); err != nil {
			return err
		}

		path := filepath.Join("static", "blog", article.ID(), "index.html")
		if err := os.WriteFile(path, buf.Bytes(), 0666); err != nil {
			return err
		}
	}
	return nil
}
